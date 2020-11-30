package main

import "container/heap"

type tower struct {
	left   int
	right  int
	height int
	index  int
}

type PriorityQ []*tower

func (pq PriorityQ) Less(i, j int) bool { return pq[i].height > pq[j].height }
func (pq PriorityQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq PriorityQ) Len() int { return len(pq) }

func (pq *PriorityQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*tower)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQ) Pop() interface{} {
	n := len(*pq)
	old := *pq
	item := old[n-1]
	item.index = -1
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func getSkyline(buildings [][]int) [][]int {
	m := map[int][]*tower{}
	for _, b := range buildings {
		t := &tower{
			left:   b[0],
			right:  b[1],
			height: b[2],
		}
		m[t.left] = append(m[t.left], t)

		m[t.right] = append(m[t.right], t)
	}

	var positions []int
	for ind, _ := range m {
		positions = append(positions, ind)
	}

	sort.Ints(positions)

	pq := PriorityQ{&tower{
		left:   -1,
		right:  -1,
		height: 0,
	}}

	var resp [][]int
	for _, p := range positions {
		curr := pq[0].height

		for _, b := range m[p] {
			// fmt.Println(" pos ", p, " building ", b)
			if b.right > p {
				heap.Push(&pq, b)
			} else {
				heap.Remove(&pq, b.index)
			}
		}

		// fmt.Println("position ", p, " pQ " , pq, " 0th index ", pq[0])
		if curr != pq[0].height {
			resp = append(resp, []int{p, pq[0].height})
		}
	}

	return resp
}
