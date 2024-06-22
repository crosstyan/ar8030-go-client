package bb

const (
	BB_MAC_LEN         = 4    /**<MAC地址字节长度*/
	BB_REG_PAGE_NUM    = 16   /**<基带寄存器页表的数量*/
	BB_REG_PAGE_SIZE   = 256  /**<基带寄存器页表的字节数量*/
	BB_CFG_PAGE_SIZE   = 1024 /**<基带配置文件分页的字节数量*/
	BB_PLOT_POINT_MAX  = 10   /**<基带plot事件的最大数据点数量*/
	BB_BLACK_LIST_SIZE = 3    /**<基带配对黑名单大小*/
	BB_RC_FREQ_NUM     = 4

	// socket option flags
	BB_SOCK_FLAG_RX       = 1 << 0 /**<@attention 指示socket传输方向为接收的bit位标志*/
	BB_SOCK_FLAG_TX       = 1 << 1 /**<@attention 指示socket传输方向为发送的bit位标志*/
	BB_SOCK_FLAG_TROC     = 1 << 2 /**<@attention 指示socket当基带连接时清空TX buffer中的历史数据（TX buffer reset on connect），仅芯片侧支持*/
	BB_SOCK_FLAG_DATAGRAM = 1 << 3 /**<@attention 指示socket传输为数据包模式，仅host driver侧支持*/

	// chan cfg flags
	BB_CHAN_HOP_AUTO       = 1 << 0 /**<@note 指示使能信道自适应的bit位标志*/
	BB_CHAN_BAND_HOP_AUTO  = 1 << 1 /**<@note 指示使能band自适应的bit位标志*/
	BB_CHAN_COMPLIANCE     = 1 << 2 /**<@note 指示使能信道合规模式*/
	BB_CHAN_MULTI_MODE     = 1 << 3 /**<@note 指示使能多套模式*/
	BB_CHAN_SUBCHAN_ENABLE = 1 << 4 /**<@note 指示使能子信道机制的bit位标志*/

	// mcs cfg flags
	BB_MCS_SWITCH_ENABLE = 1 << 0 /**<@note 指示使能MCS切换的bit位标志*/
	BB_MCS_SWITCH_AUTO   = 1 << 1 /**<@note 指示使能MCS自适应的bit位标志*/
)

const (
	BB_REQ_CFG = iota
	BB_REQ_GET
	BB_REQ_SET
	BB_REQ_RPC_IOCTL
	BB_REQ_RPC
	BB_REQ_PLAT_CTL
)
