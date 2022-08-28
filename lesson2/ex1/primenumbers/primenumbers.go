//go:build windows || linux || darwin || freebsd

// Package primenumbers can be used to calculate Prime numbers
package primenumbers

import "fmt"

// PrimeNumbers function returns the values of prime numbers up to a given number
func PrimeNumbers(a int) {
	for i := 2; i < a; i++ {
		flag := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = true
			}
		}
		if flag == true {
			continue
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
