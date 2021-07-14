package add_two_numbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		inA  []int
		inB  []int
		want []int
	}{
		{
			inA:  []int{2, 4, 3},
			inB:  []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			inA:  []int{0},
			inB:  []int{0},
			want: []int{0},
		},
		{
			inA:  []int{9, 9, 9, 9, 9, 9, 9},
			inB:  []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, test := range testCases {
		result1 := AddTwoNumbers(test.inA, test.inB)
		if !reflect.DeepEqual(test.want, result1) {
			t.Errorf("by TwoSum, expected: %v, result: %v", test.want, result1)
		}
	}
}
