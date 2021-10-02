package problems

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func intSlice2ListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	l := &ListNode{}
	t := l
	for _, val := range vals {
		t.Next = &ListNode{Val: val}
		t = t.Next
	}
	return l.Next
}

func TestMergeTwoSortedLists(t *testing.T) {
	cases := []struct {
		l1     *ListNode
		l2     *ListNode
		result *ListNode
	}{
		{intSlice2ListNode([]int{1, 2, 4}), intSlice2ListNode([]int{1, 3, 4}), intSlice2ListNode([]int{1, 1, 2, 3, 4, 4})},
		{intSlice2ListNode([]int{}), intSlice2ListNode([]int{}), intSlice2ListNode([]int{})},
		{intSlice2ListNode([]int{}), intSlice2ListNode([]int{0}), intSlice2ListNode([]int{0})},
	}
	for _, c := range cases {
		if result := mergeTwoLists(c.l1, c.l2); !cmp.Equal(result, c.result) {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
