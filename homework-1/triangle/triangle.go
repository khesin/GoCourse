package triangle

import (
	"fmt"
	"math"
)

func TriangleResult() (float64, float64, float64) {
	fmt.Print("Посчитаем площадь, периметр и гипотенузу. ")
	fmt.Print("\nВведите катет стороны a: ")
	var a float64
	fmt.Scanln(&a)
	fmt.Print("Введите катет стороны b: ")
	var b float64
	fmt.Scanln(&b)
	var s = 0.5 * a * b
	var c = math.Sqrt(a*a + b*b)
	var p = a + b + c
	return s, c, p

}
