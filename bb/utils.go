package bb

import (
	"bytes"
	"fmt"
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
