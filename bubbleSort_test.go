package misc_sort_go

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	 testCases := []struct{
	 	in []int
	 	want []int
	}{
	 	{
			in: []int{2,4,5,1,3},
			want: []int{1,2,3,4,5},
		},
		 {
			 in: []int{3, 2, 0, 4, 1, 4,5,100,10,6, 1, 2, 2, 0 },
			 want: []int{0,0,1,1,2,2,2,3,4,4,5,6,10,100},
		 },
	}

	for _, test := range testCases {
		result := BubbleSort(test.in)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("expected: %v, result: %v", test.want, result)
		}
	}
}
