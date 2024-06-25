package bb

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

func xorCheck(buff []byte) uint8 {
	var xor uint8 = 0xff
	for _, b := range buff {
		xor ^= b
	}
	return xor
}

func FromNullTermString(data []byte) string {
	n := bytes.IndexByte(data, 0x00)
	if n == -1 {
		return string(data)
	}
	return string(data[:n])
}

func PrintAsClangArray(data []byte) {
	fmt.Println("unsigned char data[] = {")
	for i, b := range data {
		if i%16 == 0 {
			fmt.Print("\t")
		}
		fmt.Printf("0x%02x, ", b)
		if i%16 == 15 || i == len(data)-1 {
			fmt.Println()
		}
	}
	fmt.Println("};")
}

func MacLike(data []byte, sep string) string {
	var ss string
	for i, b := range data {
		ss += fmt.Sprintf("%02x", b)
		if i < len(data)-1 {
			ss += sep
		}
	}
	return ss
}

// SubscribeRequestId concatenates event type with other information to form a request id
func SubscribeRequestId(event Event) RequestId {
	return RequestId(uint32(BB_REQ_CB)<<24 | SUBSCRIBE_REQ<<16 | uint32(event))
}

func NewBuffer(size uint) *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, 0, size))
}

// NewTCPFromConn creates a new TCP connection from an existing connection
func NewTCPFromConn(conn *net.TCPConn) (*net.TCPConn, error) {
	addr := conn.RemoteAddr()
	conn, err := net.DialTCP("tcp", nil, addr.(*net.TCPAddr))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// MergeChannels merges multiple channels into one channel
//
// TODO: find a way to add channel to the merge group after the merge group is created
//   - https://go.dev/blog/pipelines
//   - https://medium.com/justforfunc/two-ways-of-merging-n-channels-in-go-43c0b57cd1de
//   - https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6
func MergeChannels[T any](cs ...<-chan T) (<-chan T, func(...<-chan T)) {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan T) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	addChannels := func(cs ...<-chan T) {
		wg.Add(len(cs))
		for _, c := range cs {
			go func(c <-chan T) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out, addChannels
}

func mergeTwo[T any](a, b <-chan T) <-chan T {
	c := make(chan T)

	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func RecursiveMergeChannels[T any](chs ...<-chan T) <-chan T {
	switch len(chs) {
	case 0:
		c := make(chan T)
		close(c)
		return c
	case 1:
		return chs[0]
	default:
		m := len(chs) / 2
		return mergeTwo(
			RecursiveMergeChannels(chs[:m]...),
			RecursiveMergeChannels(chs[m:]...))
	}
}
