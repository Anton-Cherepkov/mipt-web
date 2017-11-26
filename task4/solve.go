package main

import (
	"strings"
	"unicode"
)

func RemoveEven(arr []int) (result []int) {
	for _, value := range arr {
		if value%2 != 0 {
			result = append(result, value)
		}
	}
	return result
}

func PowerGenerator(base int) func() int {
	lastNumber := base
	return func() (result int) {
		result = lastNumber
		lastNumber *= 2
		return
	}
}

func DifferentWordsCount(str string) (result int) {
	str = strings.ToLower(str)
	var word string
	cnt := make(map[string]int)

	for i := 0; i < len(str); i++ {
		if !unicode.IsLetter(rune(str[i])) {
			if len(word) > 0 {
				cnt[word] += 1
			}
			word = ""
		} else {
			word += string([]byte{str[i]})
		}
	}
	if len(word) > 0 {
		cnt[word] += 1
	}

	return len(cnt)
}
