// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

type SomeStruct struct {
}

func TypeDefinition(val interface{}) reflect.Type {
	return reflect.TypeOf(val)
}

func main() {
	ch := make(chan int)
	values := []interface{}{10, 2.5, true, "Hi there", ch, SomeStruct{}}
	for _, i := range values {
		fmt.Printf("Type of %v is %v\n", i, TypeDefinition(i))
	}
}
