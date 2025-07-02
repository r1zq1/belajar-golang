package main

import "fmt"

func tambahSaldo(s map[string]int, norek string, jumlah int) {
	s[norek] += jumlah
}

func main() {
	rekening := map[string]int{
		"123-ABC": 1000000,
		"456-DEF": 1250000,
	}
	fmt.Println(rekening)
	tambahSaldo(rekening, "123-ABC", 500000)
	fmt.Println(rekening["123-ABC"])

	delete(rekening, "123-ABC")
	fmt.Println(rekening)

	emptyMap := make(map[string]float64)
	fmt.Println(emptyMap)
	emptyMap["XYZ"] = 123.456
	emptyMap["XYZ"] = 890.456
	fmt.Println(emptyMap)
	emptyMap["PQR"] = 321.654

	for k, v := range emptyMap {
		fmt.Printf("Key : %s, Value: %.2f\n", k, v)
	}

}
