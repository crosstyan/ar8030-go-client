package bb

import (
	"fmt"
	"github.com/barweiss/go-tuple"
	"strings"
	"time"
	"unsafe"
)

// UnsafeFromByteSlice is a helper function to convert a byte slice to a struct
// since it's casting a byte slice without any proper validation, it's unsafe
// It's expecting a C like alignment of the binary data, in native endian
func UnsafeFromByteSlice[T any](data []byte) T {
	return *(*T)(unsafe.Pointer(&data[0]))
}

type UserStatusPair = tuple.T2[int, UserStatus]
type LinkStatusPair = tuple.T2[int, LinkStatus]

func (s *GetStatusOut) getMaskedStatus(bitMap byte) ([]UserStatusPair, []LinkStatusPair) {
	// can also use CfgBmp
	// bitMap := s.RtSbmp
	var it uint32 = 0
	sts := make([]UserStatusPair, 0, BB_DATA_USER_MAX)
	lSts := make([]LinkStatusPair, 0, BB_SLOT_MAX)
	for it < 8 {
		if bitMap&0x1 == 1 {
			sts = append(sts, tuple.New2(int(it), s.UserStatus[it]))
			lSts = append(lSts, tuple.New2(int(it), s.LinkStatus[it]))
		}
		bitMap >>= 1
		it++
	}
	return sts, lSts
}

// GetCfgMaskedStatus returns the masked status of the configured slots
func (s *GetStatusOut) GetCfgMaskedStatus() ([]UserStatusPair, []LinkStatusPair) {
	return s.getMaskedStatus(s.CfgSbmp)
}

// GetRuntimeMaskedStatus returns the masked status of the runtime slots
func (s *GetStatusOut) GetRuntimeMaskedStatus() ([]UserStatusPair, []LinkStatusPair) {
	return s.getMaskedStatus(s.RtSbmp)
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
	_, _ = fmt.Fprintf(ss, "Mac:'%s'", MacToString(&s.Mac))
	_, _ = fmt.Fprintf(ss, ", ")
	uSts, lSts := s.GetRuntimeMaskedStatus()
	for i, sts := range uSts {
		idx, uSt := sts.Values()
		_, _ = fmt.Fprintf(ss, "UserStatus[%d]:%s", idx, uSt.String())
		if i < len(uSts)-1 {
			_, _ = fmt.Fprintf(ss, ", ")
		}
	}
	_, _ = fmt.Fprintf(ss, ", ")
	for i, sts := range lSts {
		idx, lSt := sts.Values()
		_, _ = fmt.Fprintf(ss, "LinkStatus[%d]:%s", idx, lSt.String())
		if i < len(lSts)-1 {
			_, _ = fmt.Fprintf(ss, ", ")
		}
	}
	_, _ = fmt.Fprintf(ss, "}")
	return ss.String()
}

func (s *GetSysInfoOut) String() string {
	ss := new(strings.Builder)
	// note that they're null terminated strings
	upTimeDur := time.Millisecond * time.Duration(s.Uptime)
	_, _ = fmt.Fprintf(ss, "{Uptime:%s, ", upTimeDur.String())
	_, _ = fmt.Fprintf(ss, "CompileTime:%s, ", FromNullTermString(s.CompileTime[:]))
	_, _ = fmt.Fprintf(ss, "SoftVer:%s, ", FromNullTermString(s.SoftVer[:]))
	_, _ = fmt.Fprintf(ss, "HardwareVer:%s, ", FromNullTermString(s.HardwareVer[:]))
	_, _ = fmt.Fprintf(ss, "FirmwareVer:%s}", FromNullTermString(s.FirmwareVer[:]))
	return ss.String()
}
