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
	Mcs          uint8  // MCS level. Note: If it is the RX end, this field is meaningless and should be obtained from the link status.
	RfMode       uint8  // TX and RX mode. The user decides the meaning of the value based on whether it is an RX parameter or a TX parameter.
	TintlvEnable uint8  // Whether to perform time domain interleaving. see `bb_timeintlv_enable_e`
	TintlvNum    uint8  // Number of interleaving blocks. `bb_timeintlv_num_e`
	TintlvLen    uint8  // Number of OFDM in the interleaving block. `bb_timeintlv_len_e`
	Bandwidth    uint8  // Bandwidth. `bb_bandwidth_e`
	FreqKhz      uint32 // Frequency point in KHz.
}

type UserStatus struct {
	TxStatus PhyStatus // Transmission side status
	RxStatus PhyStatus // Reception side status
}

type LinkStatus struct {
	State   uint8   // Link layer state. Type: bb_link_state_e
	RxMcs   uint8   // Link layer receive MCS. Type: bb_phy_mcs_e
	PeerMac MacAddr // Peer MAC
}

type GetStatusOut struct {
	Role       uint8                        // Device role. Type: bb_role_e
	Mode       uint8                        // Baseband mode. Type: bb_mode_e
	SyncMode   uint8                        // Chip sync mode. 1: enable, 0: disable
	SyncMaster uint8                        // Identity in sync mode. 1: master, 0: slave
	CfgSbmp    uint8                        // Configured SLOT bitmap
	RtSbmp     uint8                        // Runtime SLOT bitmap
	Mac        MacAddr                      // Local MAC address
	UserStatus [BB_DATA_USER_MAX]UserStatus // Physical user status
	LinkStatus [BB_SLOT_MAX]LinkStatus      // Link status
}
