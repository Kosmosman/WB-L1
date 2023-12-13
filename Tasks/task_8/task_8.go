// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

// Используем побитовые сдвиги для установки бита
func SetBit(value *int64, pos int64, bitValue int64) {
	// Проверяем на корректный ввод позиции и бита
	if pos >= 64 {
		fmt.Printf("%d is not correct position.\n", pos)
		return
	}

	if bitValue != 1 && bitValue != 0 {
		fmt.Printf("%d is not correct bit value.\n", bitValue)
		return
	}

	// Меняем на противоположный знак, если последний бит не совпадает с предложенным
	if pos == 63 && (*value>>63) != bitValue {
		*value = -(*value)
	} else {
		var mask int64 = 1 << pos
		if bitValue == 1 {
			*value |= mask
		} else {
			*value &= ^mask
		}
	}
}

func main() {
	var num, pos, bitValue int64 = 150, 12, 1

	fmt.Printf("%.64b\t%d\n", num, num)
	SetBit(&num, pos, bitValue)
	fmt.Printf("%.64b\t%d\n", num, num)
}
