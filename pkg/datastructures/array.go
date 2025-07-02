package main

import "fmt"

type Nasabah struct {
	Nama  string
	Saldo *int
}

func main1() {
	// s := 100000
	// n1 := Nasabah{Nama: "Roni", Saldo: &s}

	var bankProduct = [3]string{"Tabungan", "Giro", "Deposito"}

	for i, p := range bankProduct {
		fmt.Printf("Product %d: %s\n", i+1, p)
	}

	var statusTransaksi = [...]string{"Pending", "Sukses",
		"Gagal", "Reversal"}
	for j, q := range statusTransaksi {
		fmt.Printf("Status transaksi %d: %s\n", j+1, q)
	}

	var p *[3]string = &bankProduct
	fmt.Println((*p)[0])

	saldo1, saldo2 := 100000, 200000
	saldoArr := [2]*int{&saldo1, &saldo2}

	*saldoArr[0] += 10000
	*saldoArr[1] += 500

	fmt.Println(saldo1, saldo2)

	// array multidimensi
	var kurs [2][2]float64 = [2][2]float64{
		{1.0, 16300.0}, // USD to USD, USD to IDR
		{0.00006, 1.0}, // IDR to USD, IDR to IDR
	}
	fmt.Println("USD to IDR", kurs[0][1])
	fmt.Println("IDR to USD", kurs[1][0])
}
