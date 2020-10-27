package main

func bagOfTokensScore(tokens []int, P int) int {
	if len(tokens) == 0 {
		return 0
	}

	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i] < tokens[j]
	})

	score := 0
	i, j := 0, len(tokens)-1

	if tokens[0] > P {
		return score
	}

	for i <= j {
		// fmt.Println(tokens)
		// fmt.Println("i", i, " j ", j, " score ", score, " pow ", P)
		if tokens[i] > P && score > 0 && j-i >= 2 {
			score -= 1
			P += tokens[j]
			j--
			continue
		}
		if P < tokens[i] && j-i < 2{
			break
		}

		P -= tokens[i]
		score +=1
		i++

		mak


	}

	return score
}
