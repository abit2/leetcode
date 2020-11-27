package main

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum % 2 == 1 {return false}

	half := sum/2

	dp := make([][]int, len(nums)+1)

	for i:=0; i<len(nums)+1; i++ {
		dp[i] = make([]int, half+1)
		for j:=0; j<=half; j++ {
			// fmt.Println("i ", i, " j ", j)
			if j == 0 {
				dp[i][j] = 1
				continue
			}

			if i == 0 {continue}

			dp[i][j] = dp[i-1][j]

			if j >= nums[i-1] {
				dp[i][j] = (dp[i-1][j-nums[i-1]] | dp[i][j])
			}
		}
	}



	for i:=0; i<len(nums)+1; i++ {
		// fmt.Println(dp[i])
		if dp[i][half] == 1 {
			return true
		}
	}
	return false
}
