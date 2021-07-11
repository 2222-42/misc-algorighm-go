package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct{
		in []int
		target int
		want []int
	}{
		{
			in: []int{2,7,11,15},
			target: 9,
			want: []int{0,1},
		},
		{
			in: []int{3, 2, 4 },
			target: 6,
			want: []int{1,2},
		},
		{
			in: []int{3, 3 },
			target: 6,
			want: []int{0,1},
		},
	}

	for _, test := range testCases {
		result := TwoSum(test.in, test.target)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("expected: %v, result: %v", test.want, result)
		}
	}
}
