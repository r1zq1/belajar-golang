package main

import (
	"fmt"

	"github.com/r1zq1/belajar-golang/internal/pointer"
	"github.com/r1zq1/belajar-golang/pkg/calculator"
	"github.com/r1zq1/belajar-golang/pkg/greetings"
	"github.com/r1zq1/belajar-golang/pkg/structs"
)

func main1() {
	fmt.Println(greetings.GreetingText1)
	fmt.Println(greetings.GreetingText2)
	fmt.Println(greetings.GreetingText3)

	a := 5
	b := 10
	zero := 0.0

	sum := calculator.Add(a, b)
	sub := calculator.Substract(a, b)
	div, _ := calculator.Divide(float64(a), float64(b))

	fmt.Println(sum, sub, div)

	divByZero, err := calculator.Divide(float64(a), zero)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Hasil: %.2f\n", divByZero)
	}

	p, q := calculator.AddSub(a, b)
	fmt.Println(p, q)

	x := 10
	pointer.WithoutPointer(x)
	fmt.Println("X setelah memanggil WithoutPointer: ", x)

	pointer.WithPointer(&x)
	fmt.Println("X setelah memanggil WithPointer: ", x)

	fmt.Println(x)
	fmt.Println(&x)

	var y *int = &x
	fmt.Println(y)
	fmt.Println(*y)

	var p1 structs.Person
	p1.Name = "John"
	p1.Age = 34
	p1.Alamat = structs.Address{"Jl. Bangau", "Bogor"}

	fmt.Println(p1, p1.Name, p1.Alamat)

	p2 := structs.Person{"Jane", 40,
		structs.Address{"Jl. Sudirman", "Jakarta"}}
	p3 := structs.Person{Name: "Anton",
		Alamat: structs.Address{City: "Cirebon"}}

	fmt.Println(p2)

	p2.SayHello()
	p2.ChangeName("Joko")

	fmt.Println(p2)

	fmt.Println(p3)

	p4 := &structs.Person{Name: "Budi", Age: 25,
		Alamat: structs.Address{City: "Bandung"}}

	fmt.Println(p4, p4.Alamat)

	p5 := structs.NewPerson("Melly", 35,
		structs.Address{City: "Surabaya"})

	fmt.Println(p5)

	point1 := Point{10, 12.5}
	fmt.Println(point1)

}

type Point struct {
	int
	float64
}
