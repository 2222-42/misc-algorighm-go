package median_of_two_sorted_arrays

func qSort(nums []int) (result []int) {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]

	left, right := []int{}, []int{}
	for _, v := range nums[1:] {
		if v > pivot {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	left = qSort(left)
	right = qSort(right)
	result = append(left, pivot)
	result = append(result, right...)

	return
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result = 0.0
	tmp := nums1
	tmp = append(tmp, nums2...)
	sortedArray := qSort(tmp)
	length := len(sortedArray)
	if length == 0 {
		return result
	}
	if length%2 == 0 {
		medIdx := length / 2
		result = float64(sortedArray[medIdx-1]+sortedArray[medIdx]) / 2

	} else {
		medIdx := length / 2
		result = float64(sortedArray[medIdx])
	}
	return result
}
