package leetcode

// https://leetcode.com/problems/swap-nodes-in-pairs/

func swapPairs(head *ListNode) *ListNode {
	pos := *head

	for pos.Next != nil {
		left := pos
		curr := pos.Next
		if curr != nil {
			right := curr.Next
			curr.Next = &left
			left.Next = right
			if right != nil {
				pos = *pos.Next
			}
		}
	}

	return head
}
