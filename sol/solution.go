package sol

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	k := len(lists)
	if k == 0 {
		return head
	}
	if k == 1 && lists[0] == nil {
		return head
	}
	finish := 0
	min := 0
	minIdx := 0
	is_min_set := false
	var current *ListNode
	for finish < k {
		for idx, _ := range lists {
			if !is_min_set && lists[idx] != nil {
				min = lists[idx].Val
				is_min_set = true
			}
			if lists[idx] != nil && lists[idx].Val <= min {
				min = lists[idx].Val
				minIdx = idx
			}
		}
		if lists[minIdx] == nil {
			return head
		}
		node := &ListNode{Val: lists[minIdx].Val}
		is_min_set = false
		if head == nil {
			head = node
			current = node
		} else {
			current.Next = node
			current = current.Next
		}
		if lists[minIdx] != nil {
			lists[minIdx] = lists[minIdx].Next
		} else {
			finish += 1
		}
	}
	return head
}
