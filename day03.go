package main

func calcRucksack(lines []string) int {
	var score int = 0

	for _, v := range lines {
		for i := 0; i < (len(v)/2)+1; i++ {
			point := byte(v[i])
			if occ := count(v[len(v)/2:], point); occ > 0 {
				var multiplier int = 0
				if point < byte('a') {
					multiplier = int((point - byte('A'))) + 27
				} else {
					multiplier = int((point - byte('a'))) + 1
				}

				score += multiplier
				break
			}
		}
	}

	return score
}

func calcRucksackOf3(lines []string) int {
	var score int = 0

	for i := 0; i < len(lines); i += 3 {
		for line, j := lines[i], 0; j < len(line); j++ {
			point := line[j]
			if count(lines[i+1], point) > 0 && count(lines[i+2], point) > 0 {
				var multiplier int = 0

				if point < byte('a') {
					multiplier = int((point - byte('A'))) + 27
				} else {
					multiplier = int((point - byte('a'))) + 1
				}

				score += multiplier
				break
			}
		}
	}

	return score
}
func count(subject string, find byte) int {
	var count int = 0
	for i := 0; i < len(subject); i++ {
		if subject[i] == find {
			count++
		}
	}
	return count
}
