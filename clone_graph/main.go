package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make([]int, 100)
	nodes := make([]*Node, 100)

	for i := 1; i <= 100; i++ {
		nodes[i-1] = &Node{
			Val: i,
		}
	}

	cloneUtil(nodes, visited, node)

	return nodes[0]
}

func cloneUtil(nodes []*Node, visited []int, node *Node) {
	if visited[node.Val-1] == 1 {
		return
	}

	n := nodes[node.Val-1]

	for _, neigh := range node.Neighbors {
		n.Neighbors = append(n.Neighbors, nodes[neigh.Val-1])
		// nodes[neigh.Val-1].Neighbors = append(nodes[neigh.Val-1].Neighbors, n)
	}
	visited[node.Val-1] = 1

	for _, neigh := range node.Neighbors {
		cloneUtil(nodes, visited, neigh)
	}
	return
}
