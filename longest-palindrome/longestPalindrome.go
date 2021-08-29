package longest_palindrome

func LongestPalindrome(s string) string {
	/* 方針1:
	stringから部分を取っていく
	1文字目は必ずOK
	n文字目がこれまでの文字に含まれていたら候補に入れる、なかったら候補に入れない
	*/

	/* 方針2:
	stringから部分を取っていく
	1文字目は必ずOK
	n文字目を含めて、確認する
		確認する方法は、真ん中の文字をとって、そこから比較していく。
	*/

	/* 方針3:
	for分回して、真ん中の文字をとって、そこから比較していく。
	*/

	result := ""

	for i := 0; i < len(s); i++ {
		even := checkValidString(s, i, i+1)
		odd := checkValidString(s, i, i)
		candidate := ""
		if len(odd) > len(even) {
			candidate = odd
		} else {
			candidate = even
		}

		if len(candidate) > len(result) {
			result = candidate
		}
	}

	return result
}

func checkValidString(s string, start, end int) string {

	for start >= 0 && end <= len(s)-1 && s[start] == s[end] {
		start -= 1
		end += 1
	}

	return s[start+1 : end]
}
