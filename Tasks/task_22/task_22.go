// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
// значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

type BigNum struct {
}

func (num *BigNum) Add(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func (num *BigNum) Sub(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func (num *BigNum) Mul(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func (num *BigNum) Div(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {
	num := BigNum{}
	a := big.NewInt(123456789012345)
	b := big.NewInt(98765432198765)
	fmt.Printf("a = %v\nb = %v\n", a, b)
	fmt.Printf("Result of sum a and b is %v\n", num.Add(a, b))
	fmt.Printf("Result of sub a and b is %v\n", num.Sub(a, b))
	fmt.Printf("Result of mul a and b is %v\n", num.Mul(a, b))
	fmt.Printf("Result of div a and b is %v\n", num.Div(a, b))

}
