package two_sum

// Brute Force
func TwoSum(nums []int, target int) []int {
	var result []int
	for i, v1 := range nums {
		for j, v2 := range nums[i+1:] {
			if v1+v2 == target {
				result = []int{i, j + i + 1}
			}
		}
	}
	return result
}

// Two-pass Hash Table
func TwoSumTwoPass(nums []int, target int) []int {
	var result []int
	hashTable := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hashTable[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		v, ok := hashTable[complement]
		if ok && v != i {
			if v < i {
				result = []int{v, i}
			} else {
				result = []int{i, v}
			}

		}
	}
	return result
}
