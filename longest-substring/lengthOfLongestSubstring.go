package longest_substring

import (
	"strings"
)

func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// 656 ms	7.7 MB
func LengthOfLongestSubstring(s string) int {
	var result int
	lengthOfString := len(s)
	// 0だったらskipする, 2以上だったら探索する必要がある。
	if lengthOfString == 0 {
		result = 0
	} else if lengthOfString == 1 {
		result = 1
	} else {
		// 全部の部分をとって、そこからUniqなものを探すというのは、選択肢の一つ。 -> 効率が悪い
		// 頭から探す
		// abcabcda
		// a -> ab -> abc -> (end) => 3
		// b -> bc -> bca -> (end) => 3
		// c -> ca -> cab -> (end) => 3
		// a -> ab -> abc -> abcd -> (end) => 4
		// b -> bc -> bcd -> bcda -> (end) => 4

		// 2重ループ回していくか
		sliceOfString := strings.Split(s, "")
		//fmt.Println(sliceOfString)

		for outerPosition := 0; outerPosition < lengthOfString; outerPosition++ {
			currentStr := []string{}
			for currentPosition := outerPosition; currentPosition < lengthOfString; currentPosition++ {
				//fmt.Printf("currentPositon: %d\n", currentPosition)
				targetStr := sliceOfString[currentPosition]
				if arrayContains(currentStr, targetStr) {
					break
				} else {
					currentStr = append(currentStr, targetStr)
				}
				innerResult := len(currentStr)
				if innerResult > result {
					result = innerResult
				}
				//fmt.Printf("break! innerResult: %v, result: %v\n", innerResult, result)
			}

		}
	}

	return result
}
