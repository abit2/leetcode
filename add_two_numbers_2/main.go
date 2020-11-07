package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var st1,st2 stack
	cl1, cl2 := l1, l2
	for cl1 != nil {
		st1.push(cl1.Val)
		cl1 = cl1.Next
	}
	for cl2 != nil {
		st2.push(cl2.Val)
		cl2 = cl2.Next
	}
	// fmt.Println(st1.vals, st2.vals)

	newHead := &ListNode{}
	curr := newHead
	carry := 0
	for !st1.size() && !st2.size() {
		v1, v2 := st1.pop(), st2.pop()
		sum := (v1+v2+carry)%10
		carry = (v1+v2+carry)/10

		tmp := &ListNode{
			Val: sum,
		}

		tmp.Next = curr.Next
		curr.Next = tmp
	}

	for !st1.size() {
		v1 := st1.pop()
		sum := (v1+carry)%10
		carry = (v1+carry)/10

		tmp := &ListNode{
			Val: sum,
		}

		tmp.Next = curr.Next
		curr.Next = tmp
	}

	for !st2.size() {
		v1 := st2.pop()
		sum := (v1+carry)%10
		carry = (v1+carry)/10

		tmp := &ListNode{
			Val: sum,
		}

		tmp.Next = curr.Next
		curr.Next = tmp
	}

	if carry != 0 {
		sum := (carry)%10

		tmp := &ListNode{
			Val: sum,
		}

		tmp.Next = curr.Next
		curr.Next = tmp
	}

	return newHead.Next
}

type stack struct {
	vals []int
}

func (s *stack) pop() int {
	val := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return val
}

func (s *stack) push(val int) {
	s.vals = append(s.vals, val)
}

func (s *stack) size() bool {
	if len(s.vals) == 0 {return true}
	return false
}