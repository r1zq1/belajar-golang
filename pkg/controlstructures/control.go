package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 15 {
			fmt.Println("Angka di atas 15, keluar loop")
			break
		}
		if i == 13 {
			fmt.Println("We got special number", i)
		} else if i%5 == 0 {
			fmt.Println(i, "adalah Kelipatan 5")
		} else {
			fmt.Println("Angka ganjil:", i)
		}
	}
}
