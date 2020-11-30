package main

func canReach(arr []int, start int) bool {
	visited := make([]int, len(arr))
	var q []int

	q = append(q, start)

	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			size--
			index := q[0]
			visited[index] = 1
			if len(q) == 1 {
				q = []int{}
			} else {
				q = q[1:]
			}

			if arr[index] == 0 {
				return true
			}

			if index+arr[index] < len(arr) && visited[index+arr[index]] == 0 {
				q = append(q, index+arr[index])
			}

			if index-arr[index] >=0 && visited[index-arr[index]] == 0 {
				q = append(q, index-arr[index])
			}
		}
	}

	return false
}
