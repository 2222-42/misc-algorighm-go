package slice_of_empty_or_nil

func ReturnNilSlice() []string {
	return nil
}

func ReturnEmptySlice() []string {
	return []string{}
}

func ReturnInitialSlice() []string {
	var result []string
	return result
}

func ReturnMadeSlice() []string {
	return make([]string, 0, 0)
}
