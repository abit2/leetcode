package main

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	for _, v := range asteroids {
		// fmt.Println(res)
		if len(res) == 0 {
			res = append(res, v)
			continue
		}
		endElement := last(res)
		if endElement == 0 {
			fmt.Println("doing this wrong")
		}

		flag := 0
		for endElement > 0 && v < 0 {
			if len(res) == 0 {
				break
			}

			if magnitude(endElement) > magnitude(v) {
				flag = 1
				break
			} else if magnitude(endElement) < magnitude(v) {
				res = pop(res)
				endElement = last(res)
			} else if magnitude(endElement) == magnitude(v) {
				res = pop(res)
				endElement = last(res)
				flag = 1
				break
			}
			// edge case where it becomes empty
			if endElement == 0 {
				break
			}
		}
		if flag == 1 {
			continue
		}
		res = append(res, v)

	}
	return res
}

func last(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[len(arr)-1]
}

func pop(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	return arr[:len(arr)-1]
}

func magnitude(v int) int {
	if v < 0 {
		v = v * -1
	}
	return v
}
