/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
package main

import "fmt"

func CheckUnique(s string) bool {
	m := make(map[rune]bool)
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			ch = ch - 'A' + 'a'
		}
		if m[ch] == false {
			m[ch] = true
		} else {
			return false
		}
	}
	return true
}

func TestCheckUnique() {
	testcases := map[string]struct {
		input string
		want  bool
	}{
		"Right case 1": {"abcd", true},
		"Right case 2": {"1d^fcs", true},
		"Right case 3": {"", true},
		"Right case 4": {" -+()", true},
		"Wrong case 1": {"aabcd", false},
		"Wrong case 2": {"abCdefAaf", false},
		"Wrong case 3": {"Abcda", false},
		"Wrong case 4": {"  ", false},
	}

	errorCounter := 0
	for testName, testcase := range testcases {
		if CheckUnique(testcase.input) != testcase.want {
			fmt.Printf("Wrong answer in test\"%s\": expected %v, result is %v\n", testName, testcase.want, !testcase.want)
			errorCounter++
		}
	}
	if errorCounter == 0 {
		println("All test passed!")
	}
}

func main() {
	TestCheckUnique()
}
