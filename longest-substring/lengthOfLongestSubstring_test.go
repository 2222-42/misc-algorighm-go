package longest_substring

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		in   string
		want int
	}{
		{
			in:   "abcabcbb",
			want: 3,
		},
		{
			in:   "bbbbb",
			want: 1,
		},
		{
			in:   "pwwkew",
			want: 3,
		},
		{
			in:   "",
			want: 0,
		},
		{
			in:   "au",
			want: 2,
		},
	}

	for _, test := range testCases {
		result1 := LengthOfLongestSubstring(test.in)
		if !reflect.DeepEqual(test.want, result1) {
			t.Errorf("in: %v, expected: %v, result: %v", test.in, test.want, result1)
		}
	}
}
