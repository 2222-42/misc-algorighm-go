package longest_palindrome

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "babad",
			want: "bab",
		},
		{
			in:   "cbbd",
			want: "bb",
		},
		{
			in:   "a",
			want: "a",
		},
		{
			in:   "ac",
			want: "a",
		},
	}

	for _, test := range testCases {
		result1 := LongestPalindrome(test.in)
		if !reflect.DeepEqual(test.want, result1) {
			t.Errorf("in: %v, expected: %v, result: %v", test.in, test.want, result1)
		}
	}
}
