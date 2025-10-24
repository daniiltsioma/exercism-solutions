package yacht

func Score(dice []int, category string) int {
	diceMap := map[int]int{}
	for _, d := range dice {
		diceMap[d]++
	}

	score := 0

	switch category {
	case "ones":
		score = diceMap[1]
	case "twos":
		score = diceMap[2] * 2
	case "threes":
		score = diceMap[3] * 3
	case "fours":
		score = diceMap[4] * 4
	case "fives":
		score = diceMap[5] * 5
	case "sixes":
		score = diceMap[6] * 6
	case "full house":
		twoFound, threeFound := false, false
		total := 0
		for k, v := range diceMap {
			if v == 2 {
				twoFound = true
			} else if v == 3 {
				threeFound = true
			}
			total += k * v
		}
		if twoFound && threeFound {
			score = total
		}
	case "four of a kind":
		for k, v := range diceMap {
			if v >= 4 {
				score = k * 4
			}
		}
	case "little straight":
		littleStraight := true
		for _, v := range []int{1,2,3,4,5} {
			if diceMap[v] < 1 {
				littleStraight = false
			}
		}
		if littleStraight {
			score = 30
		}
	case "big straight":
		bigStraight := true
		for _, v := range []int{2,3,4,5,6} {
			if diceMap[v] < 1 {
				bigStraight = false
			}
		}
		if bigStraight {
			score = 30
		}
	case "choice":
		for k, v := range diceMap {
			score += k * v
		}
	case "yacht":
		for _, v := range diceMap {
			if v == 5 {
				score = 50
			}
		}
	}

	return score
}
