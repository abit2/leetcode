package main

func insertionSortList(head *ListNode) *ListNode {
	newHead := &ListNode{}
	curr := head
	for curr != nil {
		newNode := curr
		curr = curr.Next
		tmp := newHead
		for tmp != nil {
			if tmp.Next == nil {
				tmp.Next = newNode
				newNode.Next = nil
				break
			}

			if tmp.Next.Val >= newNode.Val {
				p := tmp.Next
				tmp.Next = newNode
				newNode.Next = p
				break
			}
			tmp = tmp.Next
		}
	}
	return newHead.Next
}

