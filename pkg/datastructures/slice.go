package main

import "fmt"

func potonganBiaya(tx []int, biaya int) {
	for i := range tx {
		tx[i] -= biaya
	}
}

func main2() {
	tx := []int{10000, 15000, 20000}
	potonganBiaya(tx, 500)
	fmt.Println(tx)

	mutasi := [][]float64{
		{1000.0, 2000.0, 1500.0},         // Cabang A
		{500.0, 1750.0},                  // Cabang B
		{1250.0, 1700.0, 2000.0, 3000.0}, // Cabang C
	}
	fmt.Println(mutasi)
}
