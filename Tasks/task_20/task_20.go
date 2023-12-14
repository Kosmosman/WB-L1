// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func WordsReverser(str string) string {
	s := strings.Split(str, " ")
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return strings.Join(s, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println(WordsReverser(s))
}
