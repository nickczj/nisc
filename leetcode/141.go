package leetcode

// https://leetcode.com/problems/linked-list-cycle/

func HasCycle(head *ListNode) bool {
	x := make(map[*ListNode]int)
	z := head

	for {
		if z == nil {
			return false
		}
		if _, exist := x[z]; exist {
			return true
		} else {
			x[z] = 0
		}
		if z.Next != nil {
			z = z.Next
		} else {
			return false
		}
	}
}

// https://www.geeksforgeeks.org/floyds-cycle-finding-algorithm/
func faster(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
