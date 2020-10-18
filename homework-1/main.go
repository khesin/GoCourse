package main

import (
	"fmt"

	"github.com/khesin/GoCourse/homework-1/converter"
	"github.com/khesin/GoCourse/homework-1/debit"
	"github.com/khesin/GoCourse/homework-1/triangle"
)

func main() {
	fmt.Println("Вы получите", converter.Calc(), "рублей")
	fmt.Println(triangle.TriangleResult()) //Почему-то не получается вывести текст, как в примере выше
	fmt.Println(debit.DebitResult())
}
