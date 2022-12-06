package main

func FindStartOfPacket(line string, messageLength int) int {
	var chars int = 0
	m := make(map[byte]int)
	for i, j := 0, messageLength-1; i < len(line); i, j = i+1, j+1 {
		for k := range m {
			delete(m, k)
		}
		for c := i; c <= j; c++ {
			m[line[c]]++
		}
		if len(m) == messageLength {
			chars = j + 1
			break
		}
	}
	return chars
}
