package main

func findMinHeightTrees(n int, edges [][]int) []int {
	var ans, resp []int

	ed := process(edges)

	visited := make([]int, n)
	ans = bfs(0, ed, visited)
	// fmt.Println(ans)
	visited = make([]int, n)
	ans = bfs(ans[len(ans)-1], ed, visited)
	if len(ans)%2 == 0 {
		resp = append(resp, ans[len(ans)/2])
		resp = append(resp, ans[len(ans)/2-1])
	} else {
		resp = append(resp, ans[len(ans)/2])
	}

	return resp
}

func process(edges [][]int) map[int]map[int]bool {
	m := make(map[int]map[int]bool)
	for _, ed := range edges {
		n1, n2 := ed[0], ed[1]
		if m[n1] == nil {
			m[n1] = make(map[int]bool)
		}
		if m[n2] == nil {
			m[n2] = make(map[int]bool)
		}
		m[n1][n2] = true
		m[n2][n1] = true
	}

	return m
}

func bfs(node int, ed map[int]map[int]bool, visited []int) []int {
	var q []int
	ans := []int{node}

	visited[node] = 1
	for n, v := range ed[node] {
		if v != true {
			continue
		}
		if visited[n] != 1 {
			q = append(q, n)
		}
	}

	var maxl []int
	// fmt.Println("node ", node, " visited ", visited, " height ", height, " queue ", q)
	for len(q) != 0 {
		n := q[0] //first
		q = q[1:] //pop
		tmp := bfs(n, ed, visited)
		maxl = maxi(maxl, tmp)
	}
	ans = append(ans, maxl...)

	return ans
}

func maxi(i, j []int) []int {
	if len(i) > len(j) {
		return i
	}
	return j
}

func mini(i, j int) int {
	if i < j {
		return i
	}
	return j
}
