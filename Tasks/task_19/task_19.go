// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

package main

import (
	"fmt"
)

func Reverse(str string) string {
	revString := make([]rune, len(str))
	for i, r := range str {
		revString[len(str)-i-1] = r
	}
	return string(revString)
}

func main() {
	str := []string{
		"",
		"Amogus",
		"䚵䛃䚼䚵䙰䚷䚾䚹䚸䛄䚵䚽䚿䛃䙰䚴䚾䚱䙰䙼䪀䪊䪎䪐䪒䩱",
		"Строка, and something else",
		"А роза упала на лапу Азора",
	}
	for _, i := range str {
		fmt.Printf("String before reverse: %s\n", i)
		revStr := Reverse(i)
		fmt.Printf("String after reverse:  %s\n\n", revStr)
	}
}
