package slice_of_empty_or_nil

import (
	"reflect"
	"testing"
)

func TestCheck(t *testing.T) {
	result1 := ReturnNilSlice()     // return nil
	result2 := ReturnEmptySlice()   // return []string{}
	result3 := ReturnInitialSlice() // use var
	result4 := ReturnMadeSlice()    // use make

	// nil と emptyは等しくない
	if reflect.DeepEqual(result1, result2) {
		t.Errorf("expected: %v, result: %v", result1, result2)
	}

	// nil と initialは等しい
	if !reflect.DeepEqual(result1, result3) {
		t.Errorf("expected: %v, result: %v", result3, result1)
	}

	// emptyとvarで作ったのは等しくない
	if reflect.DeepEqual(result2, result3) {
		t.Errorf("expected: %v, result: %v", result3, result2)
	}

	// emptyとmakeで作ったのは等しい
	if !reflect.DeepEqual(result2, result4) {
		t.Errorf("expected: %v, result: %v", result4, result2)
	}

	// varで作ったのとmakeで作ったのとは等しくない
	if reflect.DeepEqual(result3, result4) {
		t.Errorf("expected: %v, result: %v", result4, result2)
	}
}
