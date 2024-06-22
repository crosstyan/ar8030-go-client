package bb

type RequestId uint32 // RequestId has 24 bit req id, 8 bit domain id
type DomainId uint8
type SubCmd uint32 // SubCmd only use 24 bit actually
type EventType int

const (
	BB_MAC_LEN         = 4    /*MAC地址字节长度*/
	BB_REG_PAGE_NUM    = 16   /*基带寄存器页表的数量*/
	BB_REG_PAGE_SIZE   = 256  /*基带寄存器页表的字节数量*/
	BB_CFG_PAGE_SIZE   = 1024 /*基带配置文件分页的字节数量*/
	BB_PLOT_POINT_MAX  = 10   /*基带plot事件的最大数据点数量*/
	BB_BLACK_LIST_SIZE = 3    /*基带配对黑名单大小*/
	BB_RC_FREQ_NUM     = 4

	// socket option flags
	BB_SOCK_FLAG_RX       = 1 << 0 /*@attention 指示socket传输方向为接收的bit位标志*/
	BB_SOCK_FLAG_TX       = 1 << 1 /*@attention 指示socket传输方向为发送的bit位标志*/
	BB_SOCK_FLAG_TROC     = 1 << 2 /*@attention 指示socket当基带连接时清空TX buffer中的历史数据（TX buffer reset on connect），仅芯片侧支持*/
	BB_SOCK_FLAG_DATAGRAM = 1 << 3 /*@attention 指示socket传输为数据包模式，仅host driver侧支持*/

	// chan cfg flags
	BB_CHAN_HOP_AUTO       = 1 << 0 /*@note 指示使能信道自适应的bit位标志*/
	BB_CHAN_BAND_HOP_AUTO  = 1 << 1 /*@note 指示使能band自适应的bit位标志*/
	BB_CHAN_COMPLIANCE     = 1 << 2 /*@note 指示使能信道合规模式*/
	BB_CHAN_MULTI_MODE     = 1 << 3 /*@note 指示使能多套模式*/
	BB_CHAN_SUBCHAN_ENABLE = 1 << 4 /*@note 指示使能子信道机制的bit位标志*/

	// mcs cfg flags
	BB_MCS_SWITCH_ENABLE = 1 << 0 /*@note 指示使能MCS切换的bit位标志*/
	BB_MCS_SWITCH_AUTO   = 1 << 1 /*@note 指示使能MCS自适应的bit位标志*/
)

const (
	BB_REQ_CFG       DomainId = 0
	BB_REQ_GET       DomainId = 1
	BB_REQ_SET       DomainId = 2
	BB_REQ_CB        DomainId = 3
	BB_REQ_SOCKET    DomainId = 4
	BB_REQ_DBG       DomainId = 5
	BB_REQ_RPC       DomainId = 10
	BB_REQ_RPC_IOCTL DomainId = 11
	BB_REQ_PLAT_CTL  DomainId = 12
)

const (
	BB_EVENT_LINK_STATE     EventType = iota // 链路状态发生变化事件
	BB_EVENT_MCS_CHANGE                      // MCS等级发生变化事件
	BB_EVENT_CHAN_CHANGE                     // 工作信道发生变化事件
	BB_EVENT_PLOT_DATA                       // 用于debug的异步信号质量Plot数据
	BB_EVENT_FRAME_START                     // 每一个基带帧开始的事件
	BB_EVENT_OFFLINE                         // 当设备离线时获得通知, 仅host侧有效
	BB_EVENT_PRJ_DISPATCH                    // 项目自定义事件分发
	BB_EVENT_PAIR_RESULT                     // 配对结果事件分发
	BB_EVENT_PRJ_DISPATCH2                   // 项目自定义事件分发2
	BB_EVENT_MCS_CHANGE_END                  // MCS等级发生变化结束事件
)
