package main

func checkValidString(s string) bool {
	var st1, st2 []int
	r := []rune(s)
	if len(r) == 0 {
		return true
	}

	id := -1
	for id < len(r)-1 {
		id++
		// fmt.Println("st1 ", st1, " st2 ", st2, " ", string(r[id]))
		if r[id] == '(' {
			st1 = append(st1, id)
			continue
		} else if r[id] == ')' {
			if len(st1) == 0 && len(st2) == 0 {
				return false
			}
			if len(st1) != 0 {
				st1 = pop(st1)
				continue
			} else {
				st2 = pop(st2)
			}
		} else {
			st2 = append(st2, id)
		}
	}
	// fmt.Println("st1 ", st1, " st2 ", st2, " ", string(r[id]))
	if len(st1) > len(st2) {
		return false
	}
	for len(st1) != 0 && len(st2) != 0 {
		t1, t2 := st1[len(st1)-1], st2[len(st2)-1]
		st1, st2 = pop(st1), pop(st2)
		if t1 > t2 {
			return false
		}
	}

	return true
}

func pop(st []int) []int {
	if len(st) == 0 {
		return []int{}
	}

	st = st[:len(st)-1]
	return st
}
