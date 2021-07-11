package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		in     []int
		target int
		want   []int
	}{
		{
			in:     []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			in:     []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			in:     []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, test := range testCases {
		result1 := TwoSum(test.in, test.target)
		if !reflect.DeepEqual(test.want, result1) {
			t.Errorf("by TwoSum, expected: %v, result: %v", test.want, result1)
		}

		result2 := TwoSumTwoPass(test.in, test.target)
		if !reflect.DeepEqual(test.want, result2) {
			t.Errorf("by TwoSumTwoPass, expected: %v, result: %v", test.want, result2)
		}

		result3 := TwoSumOnePass(test.in, test.target)
		if !reflect.DeepEqual(test.want, result3) {
			t.Errorf("by TwoSumOnePass, expected: %v, result: %v", test.want, result3)
		}
	}
}
