package bb

type RequestId uint32 // RequestId has 24 bit req id, 8 bit domain id
type DomainId uint8
type SubCmd uint32 // SubCmd only use 24 bit actually
type Event int

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
	BB_EVENT_LINK_STATE     Event = iota // 链路状态发生变化事件
	BB_EVENT_MCS_CHANGE                  // MCS等级发生变化事件
	BB_EVENT_CHAN_CHANGE                 // 工作信道发生变化事件
	BB_EVENT_PLOT_DATA                   // 用于debug的异步信号质量Plot数据
	BB_EVENT_FRAME_START                 // 每一个基带帧开始的事件
	BB_EVENT_OFFLINE                     // 当设备离线时获得通知, 仅host侧有效
	BB_EVENT_PRJ_DISPATCH                // 项目自定义事件分发
	BB_EVENT_PAIR_RESULT                 // 配对结果事件分发
	BB_EVENT_PRJ_DISPATCH2               // 项目自定义事件分发2
	BB_EVENT_MCS_CHANGE_END              // MCS等级发生变化结束事件
)

const (
	SOCK_LEN_DAEMON_TO_APP = 200 * 1024
	SOCK_LEN_APP_TO_DAEMON = 256 * 1024
)

const (
	SUBSCRIBE_REQ     = 1
	SUBSCRIBE_REQ_RET = 2
	SUBSCRIBE_DAT_RET = 3
	SUBSCRIBE_REQ_FAL = 4
)

type ClkSel int

const (
	BB_CLK_100M_SEL ClkSel = iota // 100M working frequency
	BB_CLK_200M_SEL               // 200M working frequency
)

type Role int

const (
	BB_ROLE_AP  Role = iota // Network root device
	BB_ROLE_DEV             // Network leaf device
	BB_ROLE_MAX
)

type Mode int

const (
	BB_MODE_SINGLE_USER Mode = iota // Single user mode
	BB_MODE_MULTI_USER              // Multi user mode
	BB_MODE_RELAY                   // Relay mode (not supported, reserved)
	BB_MODE_DIRECTOR                // 导演模式, 一对多可靠广播模式, 不支持MCS负数
	BB_MODE_MAX
)

type PhyMcs int

const (
	BB_PHY_MCS_NEG_2 PhyMcs = iota // BPSK   CR_1/2 REP_4 single stream
	BB_PHY_MCS_NEG_1               // BPSK   CR_1/2 REP_2 single stream
	BB_PHY_MCS_0                   // BPSK   CR_1/2 REP_1 single stream
	BB_PHY_MCS_1                   // BPSK   CR_2/3 REP_1 single stream
	BB_PHY_MCS_2                   // BPSK   CR_3/4 REP_1 single stream
	BB_PHY_MCS_3                   // QPSK   CR_1/2 REP_1 single stream
	BB_PHY_MCS_4                   // QPSK   CR_2/3 REP_1 single stream
	BB_PHY_MCS_5                   // QPSK   CR_3/4 REP_1 single stream
	BB_PHY_MCS_6                   // 16QAM  CR_1/2 REP_1 single stream
	BB_PHY_MCS_7                   // 16QAM  CR_2/3 REP_1 single stream
	BB_PHY_MCS_8                   // 16QAM  CR_3/4 REP_1 single stream
	BB_PHY_MCS_9                   // 64QAM  CR_1/2 REP_1 single stream
	BB_PHY_MCS_10                  // 64QAM  CR_2/3 REP_1 single stream
	BB_PHY_MCS_11                  // 64QAM  CR_3/4 REP_1 single stream
	BB_PHY_MCS_12                  // 256QAM CR_1/2 REP_1 single stream
	BB_PHY_MCS_13                  // 256QAM CR_2/3 REP_1 single stream
	BB_PHY_MCS_14                  // QPSK   CR_1/2 REP_1 dual stream
	BB_PHY_MCS_15                  // QPSK   CR_2/3 REP_1 dual stream
	BB_PHY_MCS_16                  // QPSK   CR_3/4 REP_1 dual stream
	BB_PHY_MCS_17                  // 16QAM  CR_1/2 REP_1 dual stream
	BB_PHY_MCS_18                  // 16QAM  CR_2/3 REP_1 dual stream
	BB_PHY_MCS_19                  // 16QAM  CR_3/4 REP_1 dual stream
	BB_PHY_MCS_20                  // 64QAM  CR_1/2 REP_1 dual stream
	BB_PHY_MCS_21                  // 64QAM  CR_2/3 REP_1 dual stream
	BB_PHY_MCS_22                  // 64QAM  CR_3/4 REP_1 dual stream
	BB_PHY_MCS_MAX
)

type User int

const (
	BB_USER_0           User = iota // Physical user 0
	BB_USER_1                       // Physical user 1
	BB_USER_2                       // Physical user 2
	BB_USER_3                       // Physical user 3
	BB_USER_4                       // Physical user 4
	BB_USER_5                       // Physical user 5
	BB_USER_6                       // Physical user 6
	BB_USER_7                       // Physical user 7
	BB_USER_BR_CS                   // Physical user BR/CS
	BB_USER_BR2_CS2                 // Physical user BR2/CS2
	BB_DATA_USER_MAX                // Maximum data user identifier, used for software programming assistance
	BB_USER_SWEEP                   // Physical long sweep frequency user
	BB_USER_SWEEP_SHORT             // Physical short sweep frequency user
)

const BB_USER_MAX User = BB_USER_SWEEP_SHORT + 1

type Slot int

const (
	BB_SLOT_0   Slot = iota // Logical position SLOT0
	BB_SLOT_AP              // Logical position for AP on DEV side
	BB_SLOT_1               // Logical position SLOT1
	BB_SLOT_2               // Logical position SLOT2
	BB_SLOT_3               // Logical position SLOT3
	BB_SLOT_4               // Logical position SLOT4
	BB_SLOT_5               // Logical position SLOT5
	BB_SLOT_6               // Logical position SLOT6
	BB_SLOT_7               // Logical position SLOT7
	BB_SLOT_MAX             // Maximum logical position
)

type SlotMode int

const (
	BB_SLOT_MODE_FIXED   SlotMode = iota // SLOT数量固定不变
	BB_SLOT_MODE_DYNAMIC                 // 根据DEV的接入与退出，动态的调整帧结构
)

type LinkState int

const (
	BB_LINK_STATE_IDLE    LinkState = iota // Idle state
	BB_LINK_STATE_LOCK                     // 链路信道与对方Lock
	BB_LINK_STATE_CONNECT                  // 链路信道、数据信道都与对方Lock
	BB_LINK_STATE_MAX
)

type Band int

const (
	BB_BAND_1G Band = iota // 1G: 150MHZ  ~ 1000MHZ
	BB_BAND_2G             // 2G: 1000MHZ ~ 4000MHZ
	BB_BAND_5G             // 5G: 4000MHZ ~ 7000MHZ
	BB_BAND_MAX
)

type RFPath int

const (
	BB_RF_PATH_A RFPath = iota // 射频A路
	BB_RF_PATH_B               // 射频B路
	BB_RF_PATH_MAX
)

type BandMode int

const (
	BB_BAND_MODE_SINGLE BandMode = iota // 单频模式
	BB_BAND_MODE_2G_5G                  // 2G和5G组合
	BB_BAND_MODE_1G_2G                  // 1G和2G组合
	BB_BAND_MODE_1G_5G                  // 1G和5G组合
	BB_BAND_MODE_MAX
)

type Bandwidth int

const (
	BB_BW_1_25M Bandwidth = iota // 1.25MHZ
	BB_BW_2_5M                   // 2.5MHZ
	BB_BW_5M                     // 5MHZ
	BB_BW_10M                    // 10MHZ
	BB_BW_20M                    // 20MHZ
	BB_BW_40M                    // 40MHZ
	BB_BW_MAX
)

type TimeIntlvLen int

const (
	BB_TIMEINTLV_LEN_3  TimeIntlvLen = iota // 3  OFDM
	BB_TIMEINTLV_LEN_6                      // 6  OFDM
	BB_TIMEINTLV_LEN_12                     // 12 OFDM
	BB_TIMEINTLV_LEN_24                     // 24 OFDM
	BB_TIMEINTLV_LEN_48                     // 48 OFDM
	BB_TIMEINTLV_LEN_MAX
)

type TimeIntlvEnable int

const (
	BB_TIMEINTLV_OFF TimeIntlvEnable = iota // 交织块时域不交织
	BB_TIMEINTLV_ON                         // 交织块时域交织
	BB_TIMEINTLV_ENABLE_MAX
)

type TimeIntlvNum int

const (
	BB_TIMEINTLV_1_BLOCK TimeIntlvNum = iota // 一个交织块
	BB_TIMEINTLV_2_BLOCK                     // 两个交织块
	BB_TIMEINTLV_NUM_MAX
)

type Payload int

const (
	BB_PAYLOAD_ON  Payload = iota // slot内有payload
	BB_PAYLOAD_OFF                // slot内没有payload
	BB_PAYLOAD_MAX
)

type FCHInfoLen int

const (
	BB_FCH_INFO_96BITS  FCHInfoLen = iota // 96bits FCH, 默认推荐
	BB_FCH_INFO_48BITS                    // 48bits FCH, 谨慎使用
	BB_FCH_INFO_192BITS                   // 192bits FCH
	BB_FCH_INFO_MAX
)

type TXMode int

const (
	BB_TX_1TX      TXMode = iota // 单天线发送信号
	BB_TX_2TX_STBC               // 双天线发送相关信号
	BB_TX_2TX_MIMO               // 双天线发送独立信号
	BB_TX_MODE_MAX
)

// 定义 8030 支持的 RF 发送模式
type RXMode int

const (
	BB_RX_1T1R      RXMode = iota // 单天线接收信号
	BB_RX_1T2R                    // 双天线接收对端的单发信号
	BB_RX_2T2R_STBC               // 双天线接收对方的双发STBC信号
	BB_RX_2T2R_MIMO               // 双天线接收对方的双发MIMO信号
	BB_RX_MODE_MAX
)

// 定义 8030 功率模式
type PhyPwrMode int

const (
	BB_PHY_PWR_OPENLOOP  PhyPwrMode = iota // 开环模式
	BB_PHY_PWR_CLOSELOOP                   // 闭环模式
)

// 定义 8030 BR 跳频模式
type BRHopMode int

const (
	BB_BR_HOP_MODE_FIXED          BRHopMode = iota // 固定模式
	BB_BR_HOP_MODE_FOLLOW_UP_CHAN                  // BR信道与AP的上行信道保持同步
	BB_BR_HOP_MODE_HOP_ON_IDLE                     // BR信道在无DEV连接时，周期性改变
	BB_BR_HOP_MODE_MAX
)

// 定义 8030 auto band 的切换类型
type BandHopItem int

const (
	BB_BAND_HOP_2G_2_5G BandHopItem = iota
	BB_BAND_HOP_5G_2_2G
	BB_BAND_HOP_ITEM_MAX
)

// 定义 dfs 认证类型
type DFSType int

const (
	BB_DFS_TYPE_FCC DFSType = iota
	BB_DFS_TYPE_CE
)

// 定义 dfs 操作类型
type DFSSubCmd int

const (
	BB_DFS_CONF_GET DFSSubCmd = iota
	BB_DFS_CONF_SET
	BB_DFS_EVENT
)

// SubscribeRequestId concatenates event type with other information to form a request id
func SubscribeRequestId(event Event) RequestId {
	return RequestId(uint32(BB_REQ_CB)<<24 | SUBSCRIBE_REQ<<16 | uint32(event))
}
