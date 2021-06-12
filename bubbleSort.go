package misc_algorithm_go

func BubbleSort(slice []int) []int {
	// forを回す(長さnだから)
	//　iで配列にアクセスする(n - 1 まで) 
	for i := 0; i < len(slice) - 1; i++ {
		// 後ろから見ていく
		// iまでは整列済みなので
		for j := len(slice) - 1; j > i; j -- {
			// 交換をする [1,3,2]
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1]  = slice[j-1], slice[j]
			}
		}
		
	}
	return slice
}
