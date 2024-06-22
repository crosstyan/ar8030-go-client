package bb

func xorCheck(buff []byte) uint8 {
	var xor uint8 = 0xff
	for _, b := range buff {
		xor ^= b
	}
	return xor
}
