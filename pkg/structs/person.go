package structs

import "fmt"

type Address struct {
	Street string
	City   string
}

type Person struct {
	Name   string
	Age    int
	Alamat Address
}

func (p Person) SayHello() {
	fmt.Println("Hai, nama saya adalah ", p.Name)
}

func (p *Person) ChangeName(name string) {
	p.Name = name
}

func NewPerson(name string, age int, address Address) Person {
	return Person{
		Name:   name,
		Age:    age,
		Alamat: address,
	}
}
