package main

import (
	"fmt"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&number)

	if isPrime(number) {
		fmt.Printf("%d adalah bilangan prima.\n", number)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", number)
	}
}
