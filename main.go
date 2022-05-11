package main

import "nisc/leetcode"

func main() {
	//sort.Binary([]int{5, 2, 3, 1})
	x := leetcode.ListNode{Val: 3, Next: &leetcode.ListNode{Val: 4, Next: &leetcode.ListNode{Val: 5}}}
	println(leetcode.HasCycle(&x))
}
