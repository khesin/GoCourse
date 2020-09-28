package converter

import (
	"fmt"
	"os"
)

func Calc() int {
	var dollarsRate int = 45 //Курс доллара
	var personRubble int

	//Почему-то при вызове выводит адрес памяти	fmt.Print("На сегодня курс рубля составляет", &dollarsRate)

	fmt.Print("Давайте посчитаем, сколько денег Вы получите при обмене. Курс доллара 45 рублей. Введите числовое кол-во рублей для обмена ")
	fmt.Fscan(os.Stdin, &personRubble)
	var result = dollarsRate * personRubble
	return result
	//fmt.Println("Вы получите ", result)
}
