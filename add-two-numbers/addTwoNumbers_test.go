package add_two_numbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	node0 := ListNode{
		Val:  0,
		Next: nil,
	}
	node3 := ListNode{
		Val:  3,
		Next: nil,
	}
	node43 := ListNode{
		Val:  4,
		Next: &node3,
	}
	node243 := ListNode{
		Val:  2,
		Next: &node43,
	}
	node4 := ListNode{
		Val:  4,
		Next: nil,
	}
	node64 := ListNode{
		Val:  6,
		Next: &node4,
	}
	node564 := ListNode{
		Val:  5,
		Next: &node64,
	}
	node8 := ListNode{
		Val:  8,
		Next: nil,
	}
	node08 := ListNode{
		Val:  0,
		Next: &node8,
	}
	node708 := ListNode{
		Val:  7,
		Next: &node08,
	}

	node9 := ListNode{
		Val:  9,
		Next: nil,
	}
	node99 := ListNode{
		Val:  9,
		Next: &node9,
	}
	node999 := ListNode{
		Val:  9,
		Next: &node99,
	}
	node9999 := ListNode{
		Val:  9,
		Next: &node999,
	}

	node9a := ListNode{
		Val:  9,
		Next: nil,
	}
	node99a := ListNode{
		Val:  9,
		Next: &node9a,
	}
	node999a := ListNode{
		Val:  9,
		Next: &node99a,
	}
	node9999a := ListNode{
		Val:  9,
		Next: &node999a,
	}
	node99999 := ListNode{
		Val:  9,
		Next: &node9999a,
	}
	node999999 := ListNode{
		Val:  9,
		Next: &node99999,
	}
	node9999999 := ListNode{
		Val:  9,
		Next: &node999999,
	}

	node1 := ListNode{
		Val:  1,
		Next: nil,
	}
	node01 := ListNode{
		Val:  0,
		Next: &node1,
	}
	node001 := ListNode{
		Val:  0,
		Next: &node01,
	}
	node0001 := ListNode{
		Val:  0,
		Next: &node001,
	}
	node90001 := ListNode{
		Val:  9,
		Next: &node0001,
	}
	node990001 := ListNode{
		Val:  9,
		Next: &node90001,
	}
	node9990001 := ListNode{
		Val:  9,
		Next: &node990001,
	}
	node89990001 := ListNode{
		Val:  8,
		Next: &node9990001,
	}

	testCases := []struct {
		inA  *ListNode
		inB  *ListNode
		want *ListNode
	}{
		{
			inA:  &node243,
			inB:  &node564,
			want: &node708,
		},
		{
			inA:  &node0,
			inB:  &node0,
			want: &node0,
		},
		{
			inA:  &node9999999,
			inB:  &node9999,
			want: &node89990001,
		},
	}

	for _, test := range testCases {
		result1 := AddTwoNumbers(test.inA, test.inB)
		if !reflect.DeepEqual(test.want, result1) {
			t.Errorf("by TwoSum, expected: %v, result: %v", test.want, result1)
		}
	}
}
