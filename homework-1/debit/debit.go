package debit

import (
	"fmt"
	"math"
)

func DebitResult() float64 {
	fmt.Print("Введите сумму вклада, по которой вы хотите получить расчет:")
	var personContributionValue float64
	fmt.Scanln(&personContributionValue)
	fmt.Print("Введите процент вклада: ")
	var personProcentValue float64
	fmt.Scanln(&personProcentValue)

	//var years[4]string["1 год", "2 год", "3 год", "4 год", "5 год"]
	//	for sum := 1; sum < 5; sum++ {
	//		fmt.Println((personContribution * personProcent) / 100)
	//	}
	//sum := 1
	//for sum < 5 {
	//	sum = personContributionValue + ((personContributionValue * personProcentValue) / 100)
	//}
	var sum = personContributionValue * math.Pow((1+0.01*personProcentValue), 5)
	return sum

}
