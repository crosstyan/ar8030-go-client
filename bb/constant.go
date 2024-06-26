package bb

type RequestId uint32 // RequestId has 24 bit req id, 8 bit domain id
type SubCmd uint32    // SubCmd only use 24 bit

const (
	BB_MAC_LEN         = 4    /*MAC地址字节长度*/
	BB_REG_PAGE_NUM    = 16   /*基带寄存器页表的数量*/
	BB_REG_PAGE_SIZE   = 256  /*基带寄存器页表的字节数量*/
	BB_CFG_PAGE_SIZE   = 1024 /*基带配置文件分页的字节数量*/
	BB_PLOT_POINT_MAX  = 10   /*基带plot事件的最大数据点数量*/
	BB_BLACK_LIST_SIZE = 3    /*基带配对黑名单大小*/
	BB_RC_FREQ_NUM     = 4

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

type SockFlag byte

const (
	// socket option flags
	BB_SOCK_FLAG_RX               SockFlag = 1 << 0 /*@attention 指示socket传输方向为接收的bit位标志*/
	BB_SOCK_FLAG_TX               SockFlag = 1 << 1 /*@attention 指示socket传输方向为发送的bit位标志*/
	BB_SOCK_FLAG_TROC             SockFlag = 1 << 2 /*@attention 指示socket当基带连接时清空TX buffer中的历史数据（TX buffer reset on connect），仅芯片侧支持*/
	BB_SOCK_FLAG_DATAGRAMSockFlag SockFlag = 1 << 3 /*@attention 指示socket传输为数据包模式，仅host driver侧支持*/
)

type DomainId byte

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

type Event byte

// Note that daemon will subscribe 0...9 except 5
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

type Role byte

const (
	BB_ROLE_AP  Role = iota // Network root device
	BB_ROLE_DEV             // Network leaf device
	BB_ROLE_MAX
)

type Mode byte

const (
	BB_MODE_SINGLE_USER Mode = iota // Single user mode
	BB_MODE_MULTI_USER              // Multi user mode
	BB_MODE_RELAY                   // Relay mode (not supported, reserved)
	BB_MODE_DIRECTOR                // 导演模式, 一对多可靠广播模式, 不支持MCS负数
	BB_MODE_MAX
)

type PhyMcs byte

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

type User byte

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
	BB_DATA_USER_MAX                // 最大数据用户标识，用于软件编程辅助
	BB_USER_SWEEP_SHORT             // 物理短扫频用户
)

const BB_USER_DFS User = BB_USER_7
const BB_USER_SWEEP User = BB_DATA_USER_MAX // 物理长扫频用户
const BB_USER_MAX User = BB_USER_SWEEP_SHORT + 1

type Slot byte

const (
	BB_SLOT_0   Slot = iota // Logical position SLOT0
	BB_SLOT_1               // Logical position SLOT1
	BB_SLOT_2               // Logical position SLOT2
	BB_SLOT_3               // Logical position SLOT3
	BB_SLOT_4               // Logical position SLOT4
	BB_SLOT_5               // Logical position SLOT5
	BB_SLOT_6               // Logical position SLOT6
	BB_SLOT_7               // Logical position SLOT7
	BB_SLOT_MAX             // Maximum logical position
)

const BB_SLOT_AP Slot = BB_SLOT_0 // DEV侧用于标识AP的逻辑位置

type SlotMode byte

const (
	BB_SLOT_MODE_FIXED   SlotMode = iota // SLOT数量固定不变
	BB_SLOT_MODE_DYNAMIC                 // 根据DEV的接入与退出，动态的调整帧结构
)

type LinkState byte

const (
	BB_LINK_STATE_IDLE    LinkState = iota // Idle state
	BB_LINK_STATE_LOCK                     // 链路信道与对方Lock
	BB_LINK_STATE_CONNECT                  // 链路信道、数据信道都与对方Lock
	BB_LINK_STATE_MAX
)

type Band byte

const (
	BB_BAND_1G Band = iota // 1G: 150MHZ  ~ 1000MHZ
	BB_BAND_2G             // 2G: 1000MHZ ~ 4000MHZ
	BB_BAND_5G             // 5G: 4000MHZ ~ 7000MHZ
	BB_BAND_MAX
)

type RFPath byte

const (
	BB_RF_PATH_A RFPath = iota // 射频A路
	BB_RF_PATH_B               // 射频B路
	BB_RF_PATH_MAX
)

type BandMode byte

const (
	BB_BAND_MODE_SINGLE BandMode = iota // 单频模式
	BB_BAND_MODE_2G_5G                  // 2G和5G组合
	BB_BAND_MODE_1G_2G                  // 1G和2G组合
	BB_BAND_MODE_1G_5G                  // 1G和5G组合
	BB_BAND_MODE_MAX
)

type Bandwidth byte

const (
	BB_BW_1_25M Bandwidth = iota // 1.25MHZ
	BB_BW_2_5M                   // 2.5MHZ
	BB_BW_5M                     // 5MHZ
	BB_BW_10M                    // 10MHZ
	BB_BW_20M                    // 20MHZ
	BB_BW_40M                    // 40MHZ
	BB_BW_MAX
)

type TimeIntlvLen byte

const (
	BB_TIMEINTLV_LEN_3  TimeIntlvLen = iota // 3  OFDM
	BB_TIMEINTLV_LEN_6                      // 6  OFDM
	BB_TIMEINTLV_LEN_12                     // 12 OFDM
	BB_TIMEINTLV_LEN_24                     // 24 OFDM
	BB_TIMEINTLV_LEN_48                     // 48 OFDM
	BB_TIMEINTLV_LEN_MAX
)

type TimeIntlvEnable byte

const (
	BB_TIMEINTLV_OFF TimeIntlvEnable = iota // 交织块时域不交织
	BB_TIMEINTLV_ON                         // 交织块时域交织
	BB_TIMEINTLV_ENABLE_MAX
)

type TimeIntlvNum byte

const (
	BB_TIMEINTLV_1_BLOCK TimeIntlvNum = iota // 一个交织块
	BB_TIMEINTLV_2_BLOCK                     // 两个交织块
	BB_TIMEINTLV_NUM_MAX
)

// SoCmdOpt defines socket related command options
type SoCmdOpt byte

const (
	SoOpen SoCmdOpt = iota
	SoWrite
	SoRead
	SoClose
)

const (
	SoQueryLen   SoCmdOpt = 0x90
	SoSetTxLimit SoCmdOpt = 0x91
	SoGetTxLimit SoCmdOpt = 0x92

	SoUserBaseStart SoCmdOpt = 0xc0
	SoUserBaseEnd   SoCmdOpt = 0xff
)

// SockCmd defines socket related command
// see `bb_socket_com_opt` and `bb_sock_cmd_e`
// only use in C to forward command
/*
type SockCmd byte

const (
	BB_SOCK_QUERY_TX_BUFF_LEN SockCmd = iota
	BB_SOCK_QUERY_RX_BUFF_LEN
	BB_SOCK_READ_INV_DATA
	BB_SOCK_SET_TX_LIMIT
	BB_SOCK_GET_TX_LIMIT
)

const SO_USER_BASE_START = 0xc0
const (
	BB_SOCK_IOCTL_ECHO   SockCmd = SO_USER_BASE_START + 0
	BB_SOCK_TX_LEN_GET   SockCmd = SO_USER_BASE_START + 1
	BB_SOCK_TX_LEN_RESET SockCmd = SO_USER_BASE_START + 2

	BB_SOCK_IOCTL_MAX SockCmd = SO_USER_BASE_START + 3
)
*/

const (
	BB_CONFIG_PSRAM_ENABLE            = 1     // 使能基带PSRAM机制
	BB_CONFIG_MRC_ENABLE              = 1     // 使能基带MRC机制
	BB_CONFIG_MAX_TRANSPORT_PER_SLOT  = 4     // 每个SLOT上最大的transport数量
	BB_CONFIG_MAX_INTERNAL_MSG_SIZE   = 128   // SDK内部消息通道最大消息长度，不含消息头
	BB_CONFIG_MAX_TX_NODE_NUM         = 10    // MAC层最大发送节点数量
	BB_CONFIG_MAC_RX_BUF_SIZE         = 60000 // 默认socket的接收buffer大小
	BB_CONFIG_MAC_TX_BUF_SIZE         = 40000 // 默认socket的发送buffer大小
	BB_CONFIG_MAX_USER_MCS_NUM        = 16    // 最大用户可设置的MCS等级数量
	BB_CONFIG_MAX_CHAN_NUM            = 32    // 最大用户可设置的信道数量
	BB_CONFIG_MAX_CHAN_HOP_ITEM_NUM   = 5     // 最大跳频触发项条目数量
	BB_CONFIG_MAX_SLOT_CANDIDATE      = 5     // 每个SLOT可设置的最大候选人数量
	BB_CONFIG_BR_FREQ_OFFSET          = 0     // BR与信道的频偏值 单位：KHz
	BB_CONFIG_LINK_UNLOCK_TIMEOUT     = 1000  // Link通道超时门限 单位：毫秒
	BB_CONFIG_SLOT_UNLOCK_TIMEOUT     = 1000  // FCH超时门限 单位：毫秒
	BB_CONFIG_IDLE_SLOT_THRED         = 10    // SLOT空闲门限，用于动态slot模式，单位：秒
	BB_CONFIG_EOP_SAMPLE_NUM          = 8     // EOP处理最近样本大小
	BB_CONFIG_ENABLE_BR_MCS           = 1     // 使能BR的MCS控制，仅对1V1模式有效
	BB_CONFIG_ENABLE_BLOCK_SWITCH     = 1     // 使用阻塞式模式切换机制（实验室阶段）
	BB_CONFIG_1V1_DEV_CTRL_BR_CHAN    = 1     // 1V1模式下，使能DEV控制BR的TX信道
	BB_CONFIG_1V1_COMPT_BR            = 1     // 1V1模式下，BR压缩模式（实验室阶段）
	BB_CONFIG_ENABLE_LTP              = 1     // 使能网络隔离机制
	BB_CONFIG_ENABLE_TIME_DISPATCH    = 1     // 使能链路授时机制
	BB_CONFIG_ENABLE_FRAME_CHANGE     = 1     // 使能1V1模式下，改变帧结构的功能
	BB_CONFIG_ENABLE_RC_HOP_POLICY    = 1     // 使能选择性跳频策略
	BB_CONFIG_ENABLE_AUTO_BAND_POLICY = 0     // 使能频段自适应功能
	BB_CONFIG_ENABLE_1V1_POWER_SAVE   = 1     // 使能1V1模式的节能机制
	BB_CONFIG_DEMO_STREAM             = 0     // TBD
	BB_CONFIG_OLD_PLOT_MODE           = 0     // TBD
	BB_CONFIG_FRAME_CROPPING          = 1     // 1VN模式下，动态删除或增加csma帧结构
	BB_CONFIG_LINK_BY_GROUPID         = 0     // 分组配对开关(for hyy)
	BB_CONFIG_ENABLE_RF_FILTER_PATCH  = 0     // 使能RF滤波patch
)

type Payload byte

const (
	BB_PAYLOAD_ON  Payload = iota // slot内有payload
	BB_PAYLOAD_OFF                // slot内没有payload
	BB_PAYLOAD_MAX
)

type FCHInfoLen byte

const (
	BB_FCH_INFO_96BITS  FCHInfoLen = iota // 96bits FCH, 默认推荐
	BB_FCH_INFO_48BITS                    // 48bits FCH, 谨慎使用
	BB_FCH_INFO_192BITS                   // 192bits FCH
	BB_FCH_INFO_MAX
)

// 定义 8030 支持的 RF 发送模式
type TXMode byte

const (
	BB_TX_1TX      TXMode = iota // 单天线发送信号
	BB_TX_2TX_STBC               // 双天线发送相关信号
	BB_TX_2TX_MIMO               // 双天线发送独立信号
	BB_TX_MODE_MAX
)

// 定义8030支持的RF接收模式
type RXMode byte

const (
	BB_RX_1T1R      RXMode = iota // 单天线接收信号
	BB_RX_1T2R                    // 双天线接收对端的单发信号
	BB_RX_2T2R_STBC               // 双天线接收对方的双发STBC信号
	BB_RX_2T2R_MIMO               // 双天线接收对方的双发MIMO信号
	BB_RX_MODE_MAX
)

// 定义 8030 功率模式
type PhyPwrMode byte

const (
	BB_PHY_PWR_OPENLOOP  PhyPwrMode = iota // 开环模式
	BB_PHY_PWR_CLOSELOOP                   // 闭环模式
)

// 定义 8030 BR 跳频模式
type BRHopMode byte

const (
	BB_BR_HOP_MODE_FIXED          BRHopMode = iota // 固定模式
	BB_BR_HOP_MODE_FOLLOW_UP_CHAN                  // BR信道与AP的上行信道保持同步
	BB_BR_HOP_MODE_HOP_ON_IDLE                     // BR信道在无DEV连接时，周期性改变
	BB_BR_HOP_MODE_MAX
)

// 定义 8030 auto band 的切换类型
type BandHopItem byte

const (
	BB_BAND_HOP_2G_2_5G BandHopItem = iota
	BB_BAND_HOP_5G_2_2G
	BB_BAND_HOP_ITEM_MAX
)

// 定义 dfs 认证类型
type DFSType byte

const (
	BB_DFS_TYPE_FCC DFSType = iota
	BB_DFS_TYPE_CE
)

// 定义 dfs 操作类型
type DFSSubCmd byte

const (
	BB_DFS_CONF_GET DFSSubCmd = iota
	BB_DFS_CONF_SET
	BB_DFS_EVENT
)
