package main

import (
	c "belajar/basics/for"
	a "belajar/basics/penambahan"
	b "belajar/basics/perkalian"

	"fmt"
)

func callPerkalian() {
	hasilPerkalian := b.HitungPerkalian(10,15)
	fmt.Println(hasilPerkalian)
}

func callPenambahan() {
	HasilPenambahan := a.HitungPenambahan(14,15)
	fmt.Println(HasilPenambahan)
}

func callForLoop(x int) {
	c.PrintTriangleLeft(x)
}

func main() {
	fmt.Println("A")
	
}