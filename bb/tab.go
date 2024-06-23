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
	FreqKhz      uint32          // Frequency point in KHz.
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
