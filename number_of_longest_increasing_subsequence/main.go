package main

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))

	for i, _ := range nums {
		dp[i] = 1
		cnt[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] <= dp[j] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[i]-dp[j] == 1 {
					cnt[i] += cnt[j]
				}
			}
		}
		// fmt.Println(i, " len ", dp)
		// fmt.Println(i, " count ", cnt)
	}

	maLen := 0
	for i := 0; i < len(dp); i++ {
		maLen = max(maLen, dp[i])
	}

	count := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] == maLen {
			count += cnt[i]
		}
	}
	// fmt.Println(dp)
	return count
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
