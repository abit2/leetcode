package main

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	} else if len(nums) == 2 {
		if nums[0] == nums[1] {
			return [][]int{nums}
		}
		return [][]int{nums, []int{nums[1], nums[0]}}
	}

	var resp [][]int
	m := map[int]bool{}
	for i, v := range nums {
		// fmt.Println(m[v])
		if !m[v] {
			newNum := make([]int, len(nums)-1)
			copy(newNum[:i], nums[:i])
			copy(newNum[i:], nums[i+1:])

			slices := permuteUnique(newNum)
			// fmt.Println(slices)
			for _, sl := range slices {
				resp = append(resp, []int{v})
				resp[len(resp)-1] = append(resp[len(resp)-1], sl...)
			}
			m[v] = true
		}
	}

	return resp
}
