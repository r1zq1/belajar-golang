package intf

import (
	"fmt"

	"github.com/r1zq1/belajar-golang/pkg/unittesting"
)

type Transaction interface {
	Execute() bool
}

type Transfer struct {
	From   unittesting.Rekening
	To     unittesting.Rekening
	Amount float64
}

func (t Transfer) Execute() bool {
	fmt.Println("Transfer", t.Amount, "from", t.From.Nama,
		"to", t.To.Nama)
	return true
}

type Deposit struct {
	unittesting.Rekening
	Amount float64
}

func (d Deposit) Execute() bool {
	fmt.Println("Deposit to ", d.Rekening.Nama,
		", amount: ", d.Amount)
	return true
}

func test() {
	var tx Transaction = Transfer{}
	fmt.Println(tx.Execute())
}

func Process(t Transaction) {
	if t.Execute() {

	} else {

	}
}
