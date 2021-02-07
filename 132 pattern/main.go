package main

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	smaller := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			smaller[i] = nums[i]
			continue
		}

		if v < smaller[i-1] {
			smaller[i] = v
			continue
		}
		smaller[i] = smaller[i-1]
	}

	var b bool
	for i, v := range nums {
		if v > smaller[i] {
			aa := make([]int, len(nums))
			copy(aa, nums)
			b = binarysearch(aa[i+1:], smaller, i, v)
		}

		if b == true {
			return true
		}
	}
	return false
}

func binarysearch(nums []int, smaller []int, index int, v2 int) bool {
	sort.Slice(nums, func(ii, jj int) bool {
		return nums[ii] < nums[jj]
	})

	v1 := smaller[index]

	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2

		if v1 < nums[middle] && v2 > nums[middle] {
			return true
		}

		if v1 >= nums[middle] {
			left = middle + 1
			continue
		}

		if v2 <= nums[middle] {
			right = middle - 1
			continue
		}
	}
	return false
}
[10,5,13,4,8,4,5,11,14,9,16,10,20,8]
8
[7,2,5,10,8]
2
[1,4,4]
3
[1,2,3,4,5]
2
[2,16,14,15]
2