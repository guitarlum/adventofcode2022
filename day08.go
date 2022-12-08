package main

func findTrees(treeMap [][]int) (int, int) {
	var trees int = 0
	var scenicScore int = 0

	for y, row := range treeMap {
		for x, cell := range row {
			var upScore int = 0
			var downScore int = 0
			var leftScore int = 0
			var rightScore int = 0

			inSight := false
			for walkX := x; walkX >= 0; {
				walkX--
				if walkX == -1 {
					inSight = true
					break
				}
				if treeMap[y][walkX] >= cell {
					leftScore++
					break
				}
				leftScore++
			}

			for walkX := x; walkX <= len(row)-1; {
				walkX++
				if walkX == len(row) {
					inSight = true
					break
				}
				if treeMap[y][walkX] >= cell {
					rightScore++
					break
				}
				rightScore++
			}

			for walkY := y; walkY >= 0; {
				walkY--
				if walkY == -1 {
					inSight = true
					break
				}
				if treeMap[walkY][x] >= cell {
					upScore++
					break
				}
				upScore++
			}

			for walkY := y; walkY <= len(treeMap)-1; {
				walkY++
				if walkY == len(treeMap) {
					inSight = true
					break
				}
				if treeMap[walkY][x] >= cell {
					downScore++
					break
				}
				downScore++
			}

			currentScore := upScore * downScore * leftScore * rightScore

			if currentScore > scenicScore {
				scenicScore = currentScore
			}

			if inSight {
				trees++
			}
		}
	}

	return trees, scenicScore
}
