package median_of_two_sorted_arrays

import (
	"reflect"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	type inStruct struct {
		nums1 []int
		nums2 []int
	}
	table := []struct {
		in   inStruct
		want float64
	}{
		{
			in: inStruct{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.00000,
		},
		{
			in: inStruct{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.50000,
		},
		{
			in: inStruct{
				nums1: []int{0, 0},
				nums2: []int{0, 0},
			},
			want: 0.00000,
		},
		{
			in: inStruct{
				nums1: []int{},
				nums2: []int{1},
			},
			want: 1.00000,
		},
		{
			in: inStruct{
				nums1: []int{2},
				nums2: []int{},
			},
			want: 2.00000,
		},
	}

	for _, s := range table {
		result := FindMedianSortedArrays(s.in.nums1, s.in.nums2)
		if !reflect.DeepEqual(s.want, result) {
			t.Errorf("fail the case nums1: %v, nums2: %v,\nexpected: %v, result: %v", s.in.nums1, s.in.nums2, s.want, result)
		}

		result2 := FindMedianSortedArraysRefactored(s.in.nums1, s.in.nums2)
		if !reflect.DeepEqual(s.want, result2) {
			t.Errorf("fail the case nums1: %v, nums2: %v,\nexpected: %v, result: %v", s.in.nums1, s.in.nums2, s.want, result2)
		}
	}
}
