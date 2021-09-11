package zigzag_conversion

import (
	"reflect"
	"testing"
)

func TestConvert2(t *testing.T) {
	type inputStruct struct {
		s      string
		number int
	}
	testCases := []struct {
		in   inputStruct
		want string
	}{
		{
			in:   inputStruct{s: "PAYPALISHIRING", number: 3},
			want: "PAHNAPLSIIGYIR",
		},
		{
			in:   inputStruct{s: "PAYPALISHIRING", number: 4},
			want: "PINALSIGYAHRPI",
		},
		{
			in:   inputStruct{s: "A", number: 1},
			want: "A",
		},
	}

	for _, test := range testCases {
		result1 := Convert(test.in.s, test.in.number)
		if !reflect.DeepEqual(test.want, result1) {
			t.Errorf("in: %v, expected: %v, result: %v", test.in, test.want, result1)
		}

	}
}
