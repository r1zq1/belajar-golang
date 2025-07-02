package main

import "fmt"

func main() {
	fmt.Println("Program mulai")

	// Defer ini tetap jalan meskipun ada panic
	defer fmt.Println("Defer: ini selalu jalan di akhir")

	// Recover dari panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered dari panic:", r)
		}
		// panic("panic again")
	}()

	fmt.Println("Sebelum panic")
	panic("Ada masalah serius!")
	fmt.Println("Setelah panic (tidak akan dieksekusi)")
}
