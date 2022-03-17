package main

import "fmt"

// Check Point:
// Diamond Pattern
// - Input: Size: 5
// - Output:
//         *
//        ***
//       *****
//      *******
//     *********
//    ***********
//     *********
//      *******
//       *****
//        ***
//         *

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)


	var i, j, k  int
	for i := 0; i <- size; i++ {
		for j := size + 4 - i; j > i; j-- {
			//space, indentansi
			fmt.Printf("%d", "*")
		}
		for k = size; k > i-1; k++ {
			//
			fmt.Printf("*")
		}
		for l = size; l > i; l-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

	// TODO: answer here
}
