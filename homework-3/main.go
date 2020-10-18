package main

import (
	"github.com/khesin/GoCourse/homework-3/carinitialize"
	"github.com/khesin/GoCourse/homework-3/fifo"
	"github.com/khesin/GoCourse/homework-3/phonebook"
)

func main() {
	// Get byte data to write to file.
	// Use WriteFile to create a file with byte data.
	carinitialize.DataCar()
	fifo.Result()
	phonebook.PhoneBook(addressBook := make(map[string][]int), addressBook["Olya"] = []int{89267273423})
	//ioutil.WriteFile("example.txt", phonebook.PhoneBook(), 0664)
	//fmt.Println("DONE")
}
