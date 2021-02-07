package main

func smallestRepunitDivByK(k int) int {
	m := make(map[int]bool)
	if k == 1 {
		return 1
	}
	i := 1
	count := 1
	ok := false
	for i < k {
		count++
		i = i*10 + 1
	}
	// fmt.Println(i)
	var tmp int
	for !ok {
		tmp = i % k
		if tmp == 0 {
			return count
		}

		count++
		i = tmp*10 + 1
		// fmt.Println(i)
		if _, ok = m[tmp]; ok {
			return -1
		}
		m[tmp] = true
	}
	return -1
}
