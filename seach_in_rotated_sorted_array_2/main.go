package main

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return true
		}
		return false
	}

	var i int
	for i = 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			break
		}
	}

	original := make([]int, len(nums))
	copy(original, nums[i:])
	copy(original[len(nums)-i:], nums[:i])

	return binarySearch(original, target)
}

func binarySearch(a []int, target int) bool {
	left, right := 0, len(a)-1
	for left <= right {
		m := (left + right) >> 1
		if a[m] == target {
			return true
		}
		if a[m] < target {
			left = m + 1
			continue
		}
		if a[m] > target {
			right = m - 1
			continue
		}
	}
	return false
}
