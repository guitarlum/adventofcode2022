package main

func findStartOfPacket(line string, messageLength int) int {
	for i, j, m := 0, messageLength-1, make(map[byte]int); i < len(line); i, j, m = i+1, j+1, clear(m) {
		for c := i; c <= j; c++ {
			m[line[c]]++
		}
		if len(m) == messageLength {
			return j + 1
		}
	}
	return 0
}

func clear(m map[byte]int) map[byte]int {
	for k := range m {
		delete(m, k)
	}
	return m
}
