package main

import "fmt"

// Demo:
// Program untuk check anka Palindrome yaitu angka yang dibaca dari kanan ataupun kiri menghasilkan bilangan yang sama, contoh: 1001, 1212, 1331, 2442, dll
// - Input: Angka Positive
// - Output: Angka Positive adalah angka Palindrome atau Angka Positive bukan angka Palindrome

// Contoh:
// Input:
// - Masukan angka positive : 1001
// Output:
// - Result: 1001 adalah angka Palindrome

func main() {
	var number, remainder, temp int
	var reverse int = 0

	fmt.Print("Masukan angka positive : ")
	fmt.Scan(&number)

	temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10
		//number =number/10
		// number = 0 / 10
		// number 0
		//number = 10/0
		//undefined
		//10/0.1
		//100/1
		//10/0.000001
		//10/0
		//penentu
		if number == 0 {
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d adalah angka Palindrome", temp)
	} else {
		fmt.Printf("%d bukan angka Palindrome", temp)
	}

}
