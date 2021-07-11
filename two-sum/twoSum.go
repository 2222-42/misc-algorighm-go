package two_sum

func TwoSum(nums []int, target int) []int {
	var result  []int
	for i, v1 := range nums {
		for j, v2 := range nums[i+1:] {
			if v1 + v2 == target {
				result = []int{i,j+i+1}
			}
		}
	}
	return result
}
