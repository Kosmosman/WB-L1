/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

package main

import (
	"errors"
	"log"
)

// Не поощрается использование глобальных переменных, поскольку они усложняют отладку кода.
// Рекомендуется максимально возможно сужать область видимости переменных
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// Неизвестно, что какой длины возвращается строка из функции createHugeString. Возможен выброс паники вследствие out of rage.
	// Также используется только 100 символов из 1024, что приводит к нерациональному использованию памяти
	justString = v[:100]
	println(justString)
}

func createHugeString(length int) string {
	if length < 0 {
		log.Fatal("Wrong length")
	}
	var result string
	for i := 0; i < length; i++ {
		result += string(rune(i + 33))
	}
	return result

}

// Добавим проверку и возврат ошибки
func myImplementationSomeFunc() (string, error) {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		return v[:100], nil
	} else {
		return "", errors.New("length is less than 100")
	}
}

func main() {
	someFunc()
	res, err := myImplementationSomeFunc()
	if err != nil {
		log.Fatal(err)
	}
	println(res)
}
