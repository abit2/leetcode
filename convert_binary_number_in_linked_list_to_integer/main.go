package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var s []int
	for head != nil {
		s = append([]int{head.Val}, s...)
		head = head.Next
	}

	ans := 0
	for i, v := range s {
		if v == 1 {
			ans += (1 << i)
		}
	}
	return ans
}
