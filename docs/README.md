# Code analysis of AR8030 client

I don't plan to rewrite the daemon but
I do need to reimplement the client.
I can't bare C code, it's hard to read.

```c
// bb_dev_handle_t overloaded
// a linked list of callback sessions (by `dev_insert_session`)
// and a linked list of devices (handles)
//
// Actually I'm confused how could that be possible
typedef struct bb_dev_handle_t {
    bb_host_t*       phost;
    struct list_head bb_dev_handle_list;

    void*            ioctl_sess;
    struct list_head cblshead;
    pthread_mutex_t  cbmtx;
    pthread_cond_t   cbcv;

    pthread_mutex_t ioctl_lk;

    uint32_t sel_id;
} bb_dev_handle_t;

typedef struct BB_HANDLE {
    SOCKETFD rpcfd;
    int bg_running_flg;
    pthread_t pbg;
    BASE_SESSION* session;
    uint32_t workid;
    int initflg;
    char txbuff[MAX_TXBUF];
} BB_HANDLE;

typedef struct BASE_SESSION {
    const BASE_FUN*   fun;
    struct BB_HANDLE* phd;
    pthread_mutex_t   mtx;
    pthread_cond_t    cv;
    uint8_t           wakeup_need : 1;
    uint8_t           wakeup_set  : 1;
    uint8_t           exit_flg    : 1;
    pthread_cond_t    exitcv;
} BASE_SESSION;

typedef struct CB_SESSION {
    BASE_SESSION            base;
    bb_set_event_callback_t pcb;
    struct bb_dev_handle_t* pdev;
    int                     staflg;
    struct list_head        node;
} CB_SESSION;

typedef struct {
    uint32_t id;
    uint32_t status;
    bb_dev_info_t bb_mac;
} bb_event_hotplug_t;
```

```rust
// I chose to close Rust syntax for the sake of readability
// it's a pseudo code after all

// (work_id i.e. sel_id, i.e. bb_dev_t)
bb_dev_open(work_id) => bb_dev_handle_t

// bb_dev_t is only a work id
// bb_dev_handle_t has session stuff
// like a linked list of callback sessions
// and mutex/condition variable

// NOTE: Unless explicit mentioning, all of the function types are
// treated as closure by ignoring the user data (void*) argument
type bb_event_callback = () => Unit 
type bb_ev_cb = bb_event_callback
type EventType = bb_event_e
struct bb_set_event_callback_t {
    cb: bb_ev_cb
    event: EventType
}
ar8030_driver_bb_set_cb // note it just a wrapper from `uav demo`
    -> bb_ioctl(&bb_dev_handle_t, BB_SET_EVENT_SUBSCRIBE | BB_SET_EVENT_UNSUBSCRIBE)
        -> cb_bb_ioctl(&bb_dev_handle_t, bb_set_event_callback_t)
            type cb_chk_t = (&CB_SESSION) => int
            // naming convention
            // CB_SESSION cb (cb means callback)
            // BASE_SESSION bs
            let pcb: Optional[&CB_SESSION] = 
                -> dev_find_node(&bb_dev_handle_t, cb_chk)
                    -> // find node by linked list 
                       // by callback sessions `cblshead`
            match pcb:
                Some(cb) => // update the existing callback session
                            // won't do anything else since the callback session
                            // (i.e. the pthread) is already there
                None => -> create_new_cb(...) // @ref
            @ref create_new_cb(&bb_dev_handle_t, bb_ev_cb)
                let cbsession = 
                    -> get_new_cb(&bb_dev_handle_t, bb_ev_cb) -> &CB_SESSION
                        -> bb_gethandle(&bb_dev_handle_t) => &BB_HANDLE
                            -> bb_gethandle_from_host(&bb_host_t)
                                // using the `bb_host_t` to create a new TCP connection
                                // NOTE: finally something interesting
                                //
                                // a separate thread for reading the data 
                                // for only this callback session
                                @annotate("create a new thread")
                                pthread_create(bb_read_thread)
                // send the SUBSCRIBE command to the daemon (i.e. host)
                // wait for condition variable
                @block("until condition variable is set by `bb_read_thread`")
                -> bs_send_usbpack_and_wait(&BASE_SESSION, ...) 
                // check `cbsession` i.e. the return value of setting up the callback session
                -> dev_insert_session(&bb_dev_handle_t, &CB_SESSION)
            // ignore the UNSUBSCRIBE path
        // ignore the other path of `bb_ioctl`
bb_read_thread(handle: &BB_HANDLE)
    -> for pack in loop
        if pack.reqid == BB_RPC_SEL_ID && pack.sta < 0:
            // close socket since there's no valid work_id found in host/daemon
            // @see `rpc_recv_chk_init` in `rpc_dev_bind.c` in daemon
            // I'm not sure if one thread in daemon is designed 
            // to handle one device per TCP connection or even one callback session
        handle.session.fun.rdcb(&BB_HANDLE, pack, handle.session) // call the registered callback function
        // update buffer size
        // remove stuff when the current size is over half of the buffer size
        // no idea why
    // session close, clean up the session
    handle.bg_running_flg = false
```

See also `BB_RPC_GET_HOTPLUG_EVENT` for detecting hotplug (add/remove device
when daemon is running)

See also `PF_AR8030_Init` in `pf_ar8030.c` for the initialization of the client.

## Daemon

`rpc_recv_chk_init` runs in a separate thread for each TCP connection.
