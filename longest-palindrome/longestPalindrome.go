package longest_palindrome

import (
	"fmt"
	"strings"
)

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

	newCandidates := []string{}
	for _, comp := range candidates {
		lenOfComp := len(comp)
		simplified := ""
		med := (lenOfComp / 2)
		if lenOfComp%2 == 0 {
			simplified = checkValidString(comp, med-1, med)
		} else {
			simplified = checkValidString(comp, med, med)
		}

		newCandidates = append(newCandidates, simplified)
	}

	result := ""
	maxLength := 0
	for _, candidate := range newCandidates {
		if len(candidate) > maxLength {
			maxLength = len(candidate)
			result = candidate
		}
	}

	return result
}

func checkValidString(s string, start, end int) string {
	fmt.Printf("start: %d, end: %d\n", start, end)
	for start >= 0 && end <= len(s)-1 && s[start] == s[end] {
		start -= 1
		end += 1
	}
	fmt.Printf("start: %d, end: %d\n", start, end)

	return s[start+1 : end]
}
