package ihatefibonnaci

import "fmt"

// good article about fibonnaci https://www.mathsisfun.com/numbers/fibonacci-sequence.html

func fibonacci(n int) int {
	var a = 0
	var b = 1

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}

func IHateFibonnaci() {
	print("Пусть в огороженном месте имеется пара кроликов (самка и самец) в первый день января. Эта пара кроликов производит новую пару кроликов (самку и самца) в первый день февраля и затем в первый день каждого следующего месяца. Каждая новорожденная пара кроликов становится зрелой уже через месяц и затем через месяц дает жизнь новой паре кроликов. Посчитайте самостоятельно, сколько кроликов должно быть через n-е кол-во месяцев. Как только посчитаете, перепроверьте себя")
	println("\nВведите кол-во месяцев, чтобы узнать, сколько кроликов родится")
	var month int
	fmt.Scanln(&month)
	for n := 0; n < (month + 1); n++ {
		rabbit := fibonacci(n)
		fmt.Printf("%v = %v\n", n, rabbit)
	}
}
