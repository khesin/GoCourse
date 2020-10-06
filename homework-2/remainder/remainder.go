package remainder

import "fmt"

func NoRemainder() {
	fmt.Print("Введите число, а мы проверим делится ли оно на 3 без остатка: ")
	var personnum int
	fmt.Scanln(&personnum)
	if personnum%3 == 0 {
		fmt.Println(personnum, "Число поделено без остатка")
	} else {
		fmt.Println(personnum, "Число не делится без остатка")
	}
}
