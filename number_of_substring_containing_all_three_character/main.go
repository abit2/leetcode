package main

func numberOfSubstrings(s string) int {
	store := make([][]int, 3)

	for i, c := range s {
		store[c-'a'] = append(store[c-'a'], i)
	}
	// fmt.Println(store)

	indexA, indexB, indexC := 0, 0, 0
	count := 0
	for indexA < len(store[0]) && indexB < len(store[1]) && indexC < len(store[2]) {
		index, mm := maximini(store, indexA, indexB, indexC)
		count += len(s)-mm
		// fmt.Println("index ", mm, " count ", count, " char ", s[mm]=='c')

		switch s[index] {
		case 'a':
			indexA++
		case 'b':
			indexB++
		case 'c':
			indexC++
		}
	}
	return count
}

func maximini(st [][]int, a, b, c int) (int, int) {
	minimum := 1<<63-1
	maximum := -1<<63
	minimum = mini(st[0][a], minimum)
	minimum = mini(st[1][b], minimum)
	minimum = mini(st[2][c], minimum)

	maximum = maxi(st[0][a], maximum)
	maximum = maxi(st[1][b], maximum)
	maximum = maxi(st[2][c], maximum)
	return minimum, maximum
}

func mini(i, j int) int {
	if i<j {return i}
	return j
}

func maxi(i, j int) int {
	if i<j {return j}
	return i
}
