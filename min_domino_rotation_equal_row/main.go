package main

func minDominoRotations(A []int, B []int) int {
	storeB := make([]map[int]bool, 6)
	storeA := make([]map[int]bool, 6)

	storeA = help(A, storeA)
	storeB = help(B, storeB)

	store := make([]map[int]bool, 6)
	store = help(A, store)
	store = help(B, store)

	for i, _ := range storeA {
		if len(store[i]) == len(A) {
			if len(storeA[i]) < len(storeB[i]) {
				return len(storeA[i]) - (len(storeA[i]) + len(storeB[i]) - len(A))
			} else {
				return len(storeB[i]) - (len(storeA[i]) + len(storeB[i]) - len(A))
			}
		}
	}

	return -1
}

func help(A []int, store []map[int]bool) []map[int]bool {
	for i, v := range A {
		if store[v-1] == nil {
			store[v-1] = make(map[int]bool)
		}
		store[v-1][i] = true
	}
	return store
}
