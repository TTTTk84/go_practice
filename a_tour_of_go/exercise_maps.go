package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	word_map :=  make(map[string]int)
	for _, c := range strings.Fields(s){
			_, ok := word_map[c]
			// 存在した
			if ok {
				word_map[c] += 1
				continue
			} else {
				// 存在しない
				word_map[c] = 1
			}
	}
	return word_map
}

func main() {
	//wc.Test(WordCount)

	fmt.Println(WordCount("a b c a b"))
}
