package main

import "fmt"

// Demo:
// Pyramid Pattern
// - Input: Size: 10
// - Output:
//           *
//          ***
//         *****
//        *******
//       *********
//      ***********
//     *************
//    ***************
//   *****************
//  *******************

// func setWhiteSpace(i, j int) {

// }

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)
	// size = 7

	var i, j, k, l int
	for i = size; i > 0; i-- {
		for j = 1; j <= i; j++ {
			//space, indentansi
			fmt.Printf(" ")
		}
		for k = size; k > i-1; k-- {
			//
			fmt.Printf("*")
		}
		for l = size; l > i; l-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
