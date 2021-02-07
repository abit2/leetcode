package split_array_largest_sum


func splitArray(nums []int, m int) int {
	var sum int
	M := math.MinInt64
	for _, val := range nums {
		sum += val
		M = max(M, val)
	}

	var l, h, mid int
	l, h = M, sum
	for l < h {
		mid = (h+l)>>1
		cc := count(nums, mid)
		// fmt.Println("l ", l, " h ", h, " mid ", mid, " count ", cc)
		if cc > m {
			l = mid+1
		} else if cc <= m {
			h = mid
		}
	}
	return l
}

func count(nums []int, d int) int {
	ans := 1
	temp := 0
	for _, val := range nums {
		temp += val
		if temp > d {
			ans++
			temp = val
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {return i}
	return j
}