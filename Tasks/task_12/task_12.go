// Имеется последовательность строк - (cat, cat, dog, cat, tree), создать для нее собственное множество.

package main

import (
	"WB-L1/Tasks/task_12/set"
	"fmt"
)

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	s := set.CreateNewSet()
	s.AddStrings(words)
	fmt.Println(s.Data)

	s.AddString("cow")
	fmt.Println(s.Data)

	s.Erase("cat")
	fmt.Println(s.Data)

	fmt.Println(s.Contains("tree"))
}
