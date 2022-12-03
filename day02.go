package main

func calcRPC(lines []string, two bool) int {
	var score int = 0

	player := make(map[byte]int)
	player['X'] = 0
	player['Y'] = 1
	player['Z'] = 2

	oppo := make(map[byte]int)
	oppo['A'] = 0
	oppo['B'] = 1
	oppo['C'] = 2

	for _, v := range lines {
		var playerChoice int = player[v[2]]
		var oppoChoice int = oppo[v[0]]

		if two {
			switch playerChoice {
			case 0:
				playerChoice = (oppoChoice + 2) % 3
			case 1:
				playerChoice = oppoChoice
			case 2:
				playerChoice = (oppoChoice + 1) % 3
			}
		}

		var result int = 0

		if playerChoice == (oppoChoice+1)%3 {
			result = 6
		} else if oppoChoice == playerChoice {
			result = 3
		} else {
			result = 0
		}

		score += result + playerChoice + 1
	}

	return score
}
