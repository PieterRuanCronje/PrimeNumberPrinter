package main

import (
	"fmt"
	"time"
)

func main() {
	count := 1
	for {
		time.Sleep(10 * time.Millisecond)
		if isPrime(count) {
			fmt.Println(count)
		}
		count++
	}
}

func isPrime(num int) bool {
	half := num / 2
	if num == 1 || num == 2 {
		return true
	}
	for i := 2; i <= half; i++ {
		if num%i == 0 {
			return false
		}
	}
	return false
}
