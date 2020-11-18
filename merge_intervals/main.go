package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)
	var st [][]int
	st = append(st, intervals[0])
	i := 0
	for i < len(intervals)-1 {
		i++
		top := st[len(st)-1]
		if top[1] < intervals[i][0] {
			st = append(st, intervals[i])
		} else if top[1] >= intervals[i][0] {
			st = st[:len(st)-1]
			if top[1] >= intervals[i][1] {
				st = append(st, top)
			} else if top[1] < intervals[i][1] {
				st = append(st, []int{top[0], intervals[i][1]})
			}
		}
	}
	return st
}
