package main

func lastStoneWeightII(stones []int) int {
	n := len(stones)
	dp := make([][]bool, n+1)
	sum := 0

	for _, v := range stones {
		sum += v
	}
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}

	// fmt.Println(dp)
	for i := 0; i <= n; i++ {
		// fmt.Println(i)
		dp[i][0] = true
	}

	for j := 1; j <= sum; j++ {
		dp[0][j] = false
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if stones[i-1] <= j {
				dp[i][j] = dp[i][j] || dp[i-1][j-stones[i-1]]
			}
		}
	}

	// fmt.Println("sum ", sum, " dp[n] ", dp)
	var val int

	for i := sum / 2; i >= 0; i-- {
		// fmt.Println(dp[len(dp)-1][i])
		if dp[n][i] {
			val = i
			break
		}
	}

	return sum - 2*val
}
