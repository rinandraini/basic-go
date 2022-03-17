package main

import (
	"fmt"
	"math"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012
func volume(radius, height float64) float64 {
	return math.Pi * math.Pow(radius, 2) * height
}
func main() {
	// TODO: answer here
	fmt.Println(volume(3.0, 4.0))

}
