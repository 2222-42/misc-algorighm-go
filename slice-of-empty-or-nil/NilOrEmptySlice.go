package slice_of_empty_or_nil

// nil slice
func ReturnNilSlice() []string {
	return nil
}

// empty slice
func ReturnEmptySlice() []string {
	return []string{}
}

// nil slice
func ReturnInitialSlice() []string {
	var result []string
	return result
}

// empty slice
func ReturnMadeSlice() []string {
	return make([]string, 0, 0)
}
