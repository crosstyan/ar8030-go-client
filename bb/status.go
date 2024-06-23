package bb

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (s *GetStatusOut) Read(data []byte) error {
	// https://go.dev/blog/errors-are-values
	var err error
	buf := bytes.NewBuffer(data)
	err = binary.Read(buf, binary.NativeEndian, s)
	if err != nil {
		return err
	}

	return nil
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
			// TODO: figure out why it is reversed
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

//func (s *GetStatusOut) String() string {
//	ss := new(strings.Builder)
//	fmt.Fprintf(ss, "{Role:%d, ", s.Role)
//	fmt.Fprintf(ss, "Mode:%d, ", s.Mode)
//	fmt.Fprintf(ss, "SyncMode:%d, ", s.SyncMode)
//	fmt.Fprintf(ss, "SyncMaster:%d, ", s.SyncMaster)
//	fmt.Fprintf(ss, "CfgSbmp:%d, ", s.CfgSbmp)
//	fmt.Fprintf(ss, "RtSbmp:%d, ", s.RtSbmp)
//	fmt.Fprintf(ss, "Mac:%s, ", MacToString(&s.Mac))
//	uSts, lSts := s.GetMaskedStatus()
//	fmt.Fprintf(ss, "%#v, ", uSts)
//	fmt.Fprintf(ss, "%#v}", lSts)
//	return ss.String()
//}
