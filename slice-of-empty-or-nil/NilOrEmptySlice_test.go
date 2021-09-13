package slice_of_empty_or_nil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestCheck(t *testing.T) {
	s1 := ReturnNilSlice()     // return nil
	s2 := ReturnEmptySlice()   // return []string{}
	s3 := ReturnInitialSlice() // use var
	s4 := ReturnMadeSlice()    // use make

	if (len(s1) != len(s2)) || (len(s2) != len(s3)) || (len(s3) != len(s4)) {
		t.Error("every length should be equaled")
	}

	if (cap(s1) != cap(s2)) || (cap(s2) != cap(s3)) || (cap(s3) != cap(s4)) {
		t.Error("every capacity should be equaled")
	}

	// nil と emptyは等しくない
	if reflect.DeepEqual(s1, s2) {
		t.Errorf("expected: %v, result: %v", s1, s2)
	}

	// nil と initialは等しい
	if !reflect.DeepEqual(s1, s3) {
		t.Errorf("expected: %v, result: %v", s3, s1)
	}

	// emptyとvarで作ったのは等しくない
	if reflect.DeepEqual(s2, s3) {
		t.Errorf("expected: %v, result: %v", s3, s2)
	}

	// emptyとmakeで作ったのは等しい
	if !reflect.DeepEqual(s2, s4) {
		t.Errorf("expected: %v, result: %v", s4, s2)
	}

	// varで作ったのとmakeで作ったのとは等しくない
	if reflect.DeepEqual(s3, s4) {
		t.Errorf("expected: %v, result: %v", s4, s2)
	}

	if s1[:] != nil {
		t.Errorf("s1 should not be nil")
	}

	if s2[:] == nil {
		t.Errorf("s2 should be nil")
	}

	if s3[:] != nil {
		t.Errorf("s3 should not be nil")
	}

	if s4[:] == nil {
		t.Errorf("s4 should be nil")
	}

	jsonBytes1, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}
	if string(jsonBytes1) != "null" {
		t.Error("json.Marshal of s1 should be null")
	}

	jsonBytes2, err := json.Marshal(s2)
	if err != nil {
		t.Error("JSON Marshal error:", err)
	}
	if string(jsonBytes2) != "[]" {
		t.Errorf("json.Marshal of s2 should be []")
	}

	jsonBytes3, err := json.Marshal(s3)
	if err != nil {
		t.Error("JSON Marshal error:", err)
	}
	if string(jsonBytes3) != "null" {
		t.Errorf("json.Marshal of s1 should be null")
	}

	jsonBytes4, err := json.Marshal(s4)
	if err != nil {
		t.Error("JSON Marshal error:", err)
	}
	if string(jsonBytes4) != "[]" {
		t.Errorf("json.Marshal of s4 should be []")
	}

}
