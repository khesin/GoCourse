package phonebook

import (
	"fmt"
)

func PhoneBook() {
	addressBook := make(map[string][]int)

	addressBook["Igor"] = []int{89996543210}
	addressBook["Sergey"] = []int{89167243812}
	addressBook["Sergey"] = append(addressBook["Sergey"], 89155243627)
	addressBook["Galya"] = []int{89267773223}
	addressBook["Olya"] = []int{89267273423}

	//	fmt.Println(addressBook)
	for name, numbers := range addressBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}
}
