package main

const maxV = 1<<63 - 1

func maxDistToClosest(seats []int) int {
	left := make([]int, len(seats))
	right := make([]int, len(seats))

	for i, _ := range seats {
		left[i], right[i] = maxV, maxV
	}

	for i, v := range seats {
		if v == 1 {
			left[i] = 0
			continue
		}
		if i == 0 {
			continue
		}

		if left[i-1] == maxV {
			continue
		}

		left[i] = left[i-1] + 1
	}


	for i:=len(seats)-1; i>=0; i-- {
		if seats[i] == 1 {
			right[i] = 0
			continue
		}

		if i == len(seats)-1 {
			continue
		}

		if right[i+1] == maxV {
			continue
		}

		right[i] = right[i+1]+1
	}
	// fmt.Println(left)
	// fmt.Println(right)

	resp := -1
	for i:=0; i<len(seats); i++ {
		rr := mini(left[i], right[i])
		resp = maxo(rr, resp)
	}

	return resp
}

func maxo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}
