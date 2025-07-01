package unittesting

import "errors"

type Rekening struct {
	Nama   string
	Norek  string
	Saldo  float64
	Active bool
}

func (r Rekening) Status() string {
	if !r.Active {
		return "Tidak Aktif"
	}
	if r.Saldo <= 0 {
		return "Kosong"
	}
	return "Aktif"
}

func (r *Rekening) Transfer(tujuan *Rekening, jumlah float64) error {
	if !r.Active || !tujuan.Active {
		return errors.New("rekening tidak aktif")
	}
	if jumlah <= 0 {
		return errors.New("jumlah tidak valid")
	}
	if r.Saldo < jumlah {
		return errors.New("saldo tidak cukup")
	}
	r.Saldo -= jumlah
	tujuan.Saldo += jumlah
	return nil
}

func processByValue(r Rekening) {
	_ = r.Saldo + 1000
}

func processByPointer(r *Rekening) {
	_ = r.Saldo + 1000
}
