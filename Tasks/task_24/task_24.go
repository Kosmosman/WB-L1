// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package main

import (
	"WB-L1/Tasks/task_24/point"
	"fmt"
)

func main() {
	p1 := point.NewPoint(0, 0)
	p2 := point.NewPoint(5, 5)
	p3 := point.NewPoint(4, 8)
	fmt.Printf("Distance between points 1 and 2 is %f\n", point.Distance(p1, p2))
	fmt.Printf("Distance between points 1 and 3 is %f\n", point.Distance(p1, p3))
	fmt.Printf("Distance between points 2 and 3 is %f\n", point.Distance(p2, p3))
}
