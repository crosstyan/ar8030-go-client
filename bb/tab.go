package bb

type IoCtlEntry struct {
	Req   RequestId
	ISize uint
	OSize uint
}

type GetStatusIn struct {
	// UserBmp represents the bitmap of the physical users to be fetched.
	// If you are not concerned about the physical layer information, you can fill it with 0 to ignore.
	UserBmp uint16
}

type MacAddr [BB_MAC_LEN]byte

type PhyStatus struct {
	Mcs          PhyMcs          // MCS级别 注意：如果是RX端此字段无意义，应该从link status中获取 类型 bb_phy_mcs_e
	RfMode       uint8           // TX和RX模式，用户根据所在是RX参数或TX参数来决定值的意义 类型 bb_tx_mode_e 或 bb_rx_mode_e
	TintlvEnable TimeIntlvEnable // 是否进行时域交织 see `bb_timeintlv_enable_e`
	TintlvNum    TimeIntlvNum    // 交织块数量 `bb_timeintlv_num_e`
	TintlvLen    TimeIntlvLen    // 交织块OFDM数量 `bb_timeintlv_len_e`
	Bandwidth    Bandwidth       // Bandwidth. `bb_bandwidth_e`
	FreqKhz      uint32          // Frequency in KHz.
}

type UserStatus struct {
	TxStatus PhyStatus // Transmission side status
	RxStatus PhyStatus // Reception side status
}

type LinkStatus struct {
	State   LinkState // Link layer state. Type: bb_link_state_e
	RxMcs   PhyMcs    // Link layer receive MCS. Type: bb_phy_mcs_e
	PeerMac MacAddr   // Peer MAC
}

// GetStatusOut implements the output parameter of BB_GET_STATUS
type GetStatusOut struct {
	Role       Role                         // Device role. Type: bb_role_e
	Mode       Mode                         // Baseband mode. Type: bb_mode_e
	SyncMode   uint8                        // Chip sync mode. 1: enable, 0: disable
	SyncMaster uint8                        // Identity in sync mode. 1: master, 0: slave
	CfgSbmp    uint8                        // Configured SLOT bitmap
	RtSbmp     uint8                        // Runtime SLOT bitmap
	Mac        MacAddr                      // Local MAC address
	UserStatus [BB_DATA_USER_MAX]UserStatus // Physical user status
	LinkStatus [BB_SLOT_MAX]LinkStatus      // Link status
}

// GetSysInfoOut implements the output parameter of BB_GET_SYS_INFO
type GetSysInfoOut struct {
	Uptime      uint64   // 获取系统运行时间
	CompileTime [32]byte // 编译时间
	SoftVer     [32]byte // 软件版本
	HardwareVer [32]byte // 硬件版本
	FirmwareVer [32]byte // 固件版本
}

const GetCfgInMaxLength = BB_CFG_PAGE_SIZE - 12 - 1

// GetCfgIn 定义读取命令 BB_GET_CFG 的输入参数结构
type GetCfgIn struct {
	Seq    uint16 // 命令序列号，单调递增
	Mode   uint8  // 加载模式, 0:auto, 1:memory, 2:flash
	Offset uint16 // 读取基带配置文件偏移量
	Length uint16 // 读取基带配置文件的字节长度，应 <= ioctrl 缓冲区最大长度, use 1011 (0x03f3) for now
}

// GetCfgOut 定义读取命令 BB_GET_CFG 的输出参数结构
type GetCfgOut struct {
	Seq         uint16 // 命令序列号，等于请求的序列号
	Rsv         uint16 // 保留字段
	TotalLength uint16 // 基带配置文件总长度
	TotalCrc16  uint16 // 基带配置文件crc16校验码
	Offset      uint16 // 设置基带配置文件偏移量
	Length      uint16 // 设置基带配置文件的字节长度
	Data        [BB_CFG_PAGE_SIZE - 12]uint8
}

// SetPairModeIn implements the input parameter of BB_SET_PAIR_MODE
type SetPairModeIn struct {
	Start    uint8                       // 配对动作，1: 进入配对模式, 0: 退出配对模式
	SlotBmp  uint8                       // AP侧允许进行配对的SLOT位置bitmap，DEV忽略本字段
	BlackMac [BB_BLACK_LIST_SIZE]MacAddr // 进入配对时，设置禁止配对的黑名单MAC, 如不设黑名单请清0
}

// SetApMacIn implements the input parameter of BB_SET_AP_MAC
// Note that only the DEV role can set the AP MAC address.
type SetApMacIn struct {
	Mac MacAddr // AP MAC address
}

// ConfCandidates 定义配置命令 BB_CFG_CANDIDATES 的输入参数结构
type ConfCandidates struct {
	Slot   Slot                                  // 接入候选人配置，当slot>=BB_SLOT_MAX表示不固定slot位置
	MacNum uint8                                 // 指定 mac_tab 中MAC的数量
	MacTab [BB_CONFIG_MAX_SLOT_CANDIDATE]MacAddr // MAC表
}

// SetCandidatesIn is an alias of ConfCandidates, for BB_SET_CANDIDATES command
type SetCandidatesIn = ConfCandidates

// Quality means ar8030 物理用户信号质量
type Quality struct {
	Snr     uint16 // 物理用户接收信噪比，转化db公式 10log(snr/36)
	LdpcErr uint16 // 交织块中解码错误的 LDPC 块个数
	LdpcNum uint16 // 交织块中总的 LDPC 块个数
	GainA   uint8  // A路天线接收信号强度
	GainB   uint8  // B路天线接收信号强度
}

// GetPairResultOut 定义读取命令 BB_GET_PAIR_RESULT 的输出参数结构
type GetPairResultOut struct {
	SlotBmp uint8                // 配对成功的 SLOT 位置 bitmap
	PeerMac [BB_SLOT_MAX]MacAddr // 对应 SLOT 上配对成功的 MAC 地址
	Quality [BB_SLOT_MAX]Quality // 配对完成时的无线质量统计
}
