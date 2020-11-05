package main

func minCostToMoveChips(position []int) int {
	even := 0
	odd := 0
	for _, v := range position {
		remainder := v % 2
		if remainder == 0 {
			even++
		} else {
			odd++
		}
	}

	ans := odd
	if odd > even {
		ans = even
	}
	return ans
}
