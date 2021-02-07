package main

type level struct {
	glass []float64
}

func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([][]float64, 101)

	for i := 0; i < 101; i++ {
		dp[i] = make([]float64, i+1)
	}

	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			// fmt.Println(i, j, dp[i])
			if i == 0 && j == 0 {
				dp[i][j] = float64(poured)
			}

			overflow := (dp[i][j] - 1) / 2
			if overflow > 0 {
				dp[i+1][j] += overflow
				dp[i+1][j+1] += overflow
			}

		}
	}
	return math.Min(1, dp[query_row][query_glass])
}
