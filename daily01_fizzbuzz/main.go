package main

import (
	"fmt"
)

// Task: FizzBuzz
// Buat program Go yang:
// 1. Cetak angka dari 1 sampai 100
// 2. Jika angka habis dibagi 3, cetak "Fizz"
// 3. Jika angka habis dibagi 5, cetak "Buzz"
// 4. Jika angka habis dibagi 3 dan 5, cetak "FizzBuzz"
// 5. Jika tidak, cetak angka itu sendiri

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}
