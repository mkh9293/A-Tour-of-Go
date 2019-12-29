package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount : 각 단어 등장 횟수 리턴
func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	result := make(map[string]int)

	for _, v := range arr {
		result[v]++
	}

	return result
}

func exerciseMaps() {
	wc.Test(WordCount)
}
