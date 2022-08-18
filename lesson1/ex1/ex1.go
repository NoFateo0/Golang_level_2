package main

import (
	"fmt"
	"log"
)

func fork(a [3]int) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("Паника: ", v)
		}
	}()

	for i := 0; i < len(a)+1; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	s := [3]int{1, 2, 3}

	fork(s)
}
