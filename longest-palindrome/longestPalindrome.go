package longest_palindrome

import (
	"strings"
)

func LongestPalindrome(s string) string {
	/* 方針:
	stringから部分を取っていく
	1文字目は必ずOK
	n文字目がこれまでの文字に含まれていたら候補に入れる、なかったら候補に入れない
	*/

	sliceOfString := strings.Split(s, "")

	candidates := []string{}

	for j := 0; j < len(s); j++ {
		targetSlice := sliceOfString[j:]
		stocks := map[string]int{}
		candidate := ""
		for i, c := range targetSlice {
			candidate = candidate + c
			if i == 0 {
				stocks[c] = i
				candidates = append(candidates, candidate)
				continue
			}

			if _, ok := stocks[c]; ok {
				candidates = append(candidates, candidate)
				continue
			}
			stocks[c] = i
		}
	}

	maxLength := 0
	result := ""
	for _, comp := range candidates {
		if comp == reverseString(comp) && len(comp) > maxLength {
			maxLength = len(comp)
			result = comp
		}
	}

	return result
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
