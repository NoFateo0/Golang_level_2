package main

import (
	"Golang_GB_2/lesson2/ex1/primenumbers"
	"fmt"
)

func main() {
	var a int
	fmt.Print("Введите конечное число: ")
	fmt.Scanln(&a)

	primenumbers.PrimeNumbers(a)
}
