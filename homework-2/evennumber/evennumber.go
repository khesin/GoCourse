package evennumber

import "fmt"

func EvenResult() {
	fmt.Print("Мы написали функцию, которая проверяет четное ли число. Введите число: ")
	var personnum int
	fmt.Scanln(&personnum)
	if personnum%2 == 0 {
		fmt.Println(personnum, "Число четное")
	} else {
		fmt.Println(personnum, "Число не четное")
	}
}
