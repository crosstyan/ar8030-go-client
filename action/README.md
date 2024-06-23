Original Client Calling Chain

```
// (work_id i.e. sel_id, i.e. bb_dev_t)
bb_dev_open(work_id) => bb_dev_handle_t

// bb_dev_t is only a work id
// bb_dev_handle_t has session stuff
// like a linked list of callback sessions
// and mutex/condition variable
/**
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
*/

// NOTE: Unless explicit mention, all of the function types are
// treated as closure by ignoring the user data (void*) argument
using bb_event_callback = (ev: EventType, arg: Any) => Unit 
using bb_ev_cb = bb_event_callback
ar8030_driver_bb_set_cb // note it just a wrapper from `uav demo`
    -> bb_ioctl(&bb_dev_handle_t, BB_SET_EVENT_SUBSCRIBE | BB_SET_EVENT_UNSUBSCRIBE)
        -> cb_bb_ioctl(&bb_dev_handle_t, bb_ev_cb)
            using cb_chk_t = (CB_SESSION) => int 
            // naming convention
            // CB_SESSION cb
            // BASE_SESSION bs
            let pcb: Optional[&CB_SESSION] = 
                -> dev_find_node(&bb_dev_handle_t, cb_chk)
                    -> // find node by linked list 
                       // (callback sessions are stored in linked list)
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
                                pthread_create(bb_read_thread)
                // send the SUBSCRIBE command to the daemon (i.e. host)
                // wait for condition variable
                @block("until condition variable is set")
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
            // I'm not sure if the daemon is designed to handle one device per TCP connection
            // or even one callback session
        handle.session.fun.rdcb(&BB_HANDLE, pack, handle.session) // call the registered callback function
        // update buffer, no idea why
    // session close, clean up the session
    handle.bg_running_flg = false
```

See also `BB_RPC_GET_HOTPLUG_EVENT`

```c
typedef struct {
    uint32_t id;
    uint32_t status;
    bb_dev_info_t bb_mac;
} bb_event_hotplug_t;
```
