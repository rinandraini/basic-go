package main

import "fmt"

// Check Point:
// Start Pattern
// - Input: Size
// - Output: Start Pattern

// Contoh:
// - Input: Size: 10
// - Output:
//           *
//          **
//         ***
//        ****
//       *****
//      ******
//     *******
//    ********
//   *********
//  **********

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)

	for i := 1; i <= size; i++ {
		for j := i; j < size; j++ {
			fmt.Printf("%s", " ")
		}
		for k := 1; k <= i; k++ {
			fmt.Printf("%s", "*")
		}
		fmt.Println()
	}
	// TODO: answer here
}
