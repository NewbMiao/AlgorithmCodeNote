package leetcode

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//go test -v src/leetcode/* -run TestRemoveNthFromEnd
func TestRemoveNthFromEnd(t *testing.T) {
	ds := [][]int{
		{1, 2, 3, 4, 5},
		{1},
	}
	ns := []int{
		2,
		1,
	}
	ws := [][]int{
		{1, 2, 3, 5},
		{},
	}
	for k, v := range ds {
		t.Log(fmt.Sprintf("RemoveNthFromEnd, input list %v, del last %dth node , want: %v", v, ns[k], ws[k]))
		r := removeNthFromEnd(genListNode(v), ns[k])
		tmp := cvtListNode2arr(r)
		if fmt.Sprint(tmp) == fmt.Sprint(ws[k]) {
			t.Log("RemoveNthFromEnd is ok")
		} else {
			t.Error("RemoveNthFromEnd is not ok, result:", fmt.Sprintf("%v", tmp))
		}

	}
}
