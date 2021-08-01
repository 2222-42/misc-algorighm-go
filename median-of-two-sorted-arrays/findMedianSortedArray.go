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

func findKth(nums1 []int, nums2 []int, k int) float64 {
	// for loopを回して、欲しい値にたどり着くまで、削っていく。
	for {
		l1, l2 := len(nums1), len(nums2)
		m1, m2 := l1/2, l2/2

		if l1 == 0 {
			return float64(nums2[k])
		}

		if l2 == 0 {
			return float64(nums1[k])
		}

		if k == 0 {
			if n1, n2 := nums1[0], nums2[0]; n1 <= n2 {
				return float64(n1)
			} else {
				return float64(n2)
			}
		}

		// どんどん小さくしていく
		if k <= m1+m2 {
			// kがm1とm2未満の和以下だったら、欲しい値は前方にあるから、前方を取り出す
			if n1, n2 := nums1[m1], nums2[m2]; n1 <= n2 {
				// それぞれの配列の中央値n1, n2で、n1のほうが小さければ、nums2を削り、n2が大きければnums1を削る
				nums2 = nums2[:m2]
			} else {
				nums1 = nums1[:m1]
			}
		} else {
			// kがm1とm2未満の和より大きければだったら、欲しい値は後方にあるから、後方を取り出す。
			// また、求めるkの値も、削る量に合わせてoffsetしなkればならない。
			if n1, n2 := nums1[m1], nums2[m2]; n1 <= n2 {
				nums1 = nums1[m1+1:]
				k -= m1 + 1
			} else {
				nums2 = nums2[m2+1:]
				k -= m2 + 1
			}
		}
	}
}

func FindMedianSortedArraysRefactored(nums1 []int, nums2 []int) float64 {
	var result float64
	if l := len(nums1) + len(nums2); l%2 == 0 {
		result = (findKth(nums1, nums2, l/2-1) + findKth(nums1, nums2, l/2)) / 2.0
	} else {
		result = findKth(nums1, nums2, l/2)
	}

	return result
}
