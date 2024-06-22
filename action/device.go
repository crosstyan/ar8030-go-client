package action

import "net"

/*
Device is the counterpart of the following C structure:

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

Personally I don't think every entry in the C struct is necessary.
*/
type Device struct {
	conn  *net.TCPConn
	selId uint32
}
