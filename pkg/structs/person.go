package structs

import "fmt"

type Cabang struct {
	Kode        string
	Nama        string
	Alamat      string
	AnakCabang  []*Cabang
	IndukCabang *Cabang
}

type Transaksi struct {
	ID       string
	Nominal  float64
	Type     string
	Reversal *Transaksi
}

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
