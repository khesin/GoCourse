package carinitialize

import (
	"fmt"

	"github.com/khesin/GoCourse/homework-3/carsstruct"
)

func DataCar() {
	sedan := carsstruct.Carssrtuct{
		Category:    "Sedan",
		Model:       "Corolla",
		StartDate:   "2018, 01, 12",
		TrunkValue:  60,
		EngineOn:    false,
		WindowsOpen: false,
		Color:       "White",
	}
	truck := carsstruct.Carssrtuct{
		Category:    "Truck",
		Model:       "F50",
		StartDate:   "2020, 01, 12",
		TrunkValue:  80,
		EngineOn:    true,
		WindowsOpen: true,
		Color:       "Green",
	}
	fmt.Println(sedan, truck)
}
