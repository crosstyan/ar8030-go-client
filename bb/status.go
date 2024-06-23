package bb

import (
	"fmt"
	"strings"
	"unsafe"
)

// UnsafeGetStatusOut is a helper function to convert a byte slice to a [GetStatusOut] struct
// since it's casting a byte slice without any proper validation, it's unsafe
// It's expecting a C like alignment of the binary data, in native endian
func UnsafeGetStatusOut(data []byte) GetStatusOut {
	return *(*GetStatusOut)(unsafe.Pointer(&data[0]))
}

func (s *GetStatusOut) GetMaskedStatus() ([]UserStatus, []LinkStatus) {
	// can also use CfgBmp
	bitMap := s.RtSbmp
	var it uint32 = 0
	sts := make([]UserStatus, 0, BB_DATA_USER_MAX)
	lSts := make([]LinkStatus, 0, BB_SLOT_MAX)
	for it < 8 {
		if bitMap&0x1 == 1 {
			sts = append(sts, s.UserStatus[it])
			lSts = append(lSts, s.LinkStatus[it])
		}
		bitMap >>= 1
		it++
	}
	return sts, lSts
}

func MacToString(mac *MacAddr) string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x", mac[0], mac[1], mac[2], mac[3])
}

func (s *PhyStatus) String() string {
	ss := new(strings.Builder)
	_, _ = fmt.Fprintf(ss, "{Mcs:%d, ", s.Mcs)
	_, _ = fmt.Fprintf(ss, "RfMode:%d, ", s.RfMode)
	_, _ = fmt.Fprintf(ss, "TintlvEnable:%d, ", s.TintlvEnable)
	_, _ = fmt.Fprintf(ss, "TintlvNum:%d, ", s.TintlvNum)
	_, _ = fmt.Fprintf(ss, "TintlvLen:%d, ", s.TintlvLen)
	_, _ = fmt.Fprintf(ss, "Bandwidth:%d, ", s.Bandwidth)
	_, _ = fmt.Fprintf(ss, "FreqKhz:%d}", s.FreqKhz)
	return ss.String()
}

func (s *UserStatus) String() string {
	ss := new(strings.Builder)
	_, _ = fmt.Fprintf(ss, "{TxStatus:%s, ", s.TxStatus.String())
	_, _ = fmt.Fprintf(ss, "RxStatus:%s}", s.RxStatus.String())
	return ss.String()
}

func (s *LinkStatus) String() string {
	ss := new(strings.Builder)
	_, _ = fmt.Fprintf(ss, "{State:%d, ", s.State)
	_, _ = fmt.Fprintf(ss, "RxMcs:%d, ", s.RxMcs)
	_, _ = fmt.Fprintf(ss, "PeerMac:%s}", MacToString(&s.PeerMac))
	return ss.String()
}

func (s *GetStatusOut) String() string {
	ss := new(strings.Builder)
	_, _ = fmt.Fprintf(ss, "{Role:%d, ", s.Role)
	_, _ = fmt.Fprintf(ss, "Mode:%d, ", s.Mode)
	_, _ = fmt.Fprintf(ss, "SyncMode:%d, ", s.SyncMode)
	_, _ = fmt.Fprintf(ss, "SyncMaster:%d, ", s.SyncMaster)
	_, _ = fmt.Fprintf(ss, "CfgSbmp:%d, ", s.CfgSbmp)
	_, _ = fmt.Fprintf(ss, "RtSbmp:%d, ", s.RtSbmp)
	_, _ = fmt.Fprintf(ss, "Mac:%s, ", MacToString(&s.Mac))
	uSts, lSts := s.GetMaskedStatus()
	for i, sts := range uSts {
		_, _ = fmt.Fprintf(ss, "UserStatus[%d]:%s, ", i, sts.String())
	}
	for i, sts := range lSts {
		_, _ = fmt.Fprintf(ss, "LinkStatus[%d]:%s, ", i, sts.String())
	}
	_, _ = fmt.Fprintf(ss, "}")
	return ss.String()
}
