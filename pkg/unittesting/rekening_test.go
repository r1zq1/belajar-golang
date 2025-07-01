package unittesting

import (
	"testing"
)

func TestStatus(t *testing.T) {
	tests := []struct {
		nama   string
		input  Rekening
		expect string
	}{
		{
			"Aktif dan ada saldo",
			Rekening{Nama: "Andi", Saldo: 10000.0, Active: true},
			"Aktif",
		},
		{
			"Aktif tapi kosong",
			Rekening{Nama: "Budi", Saldo: 0.0, Active: true},
			"Kosong",
		},
		{
			"Tidak aktif",
			Rekening{Nama: "Citra", Saldo: 5000.0, Active: false},
			"Tidak Aktif",
		},
	}
	for _, t1 := range tests {
		t.Run(t1.nama, func(t *testing.T) {
			got := t1.input.Status()
			if got != t1.expect {
				t.Errorf("Status() = %s; ekspektasinya = %s",
					got, t1.expect)
			}
		})
	}

}

func TestTransfer(t *testing.T) {
	type testCase struct {
		nama         string
		source       Rekening
		target       Rekening
		jumlah       float64
		expectErr    bool
		expectSaldoS float64
		expectSaldoT float64
	}
	tests := []testCase{
		{
			nama:         "Transfer Sukses",
			source:       Rekening{Nama: "Andi", Saldo: 10000.0, Active: true},
			target:       Rekening{Nama: "Budi", Saldo: 6000.0, Active: true},
			jumlah:       3000.0,
			expectErr:    false,
			expectSaldoS: 7000.0,
			expectSaldoT: 9000.0,
		},
		// Saldo pengirim tidak cukup, Rekening tidak aktif, Jumlah tidak valid
	}
	for _, tc := range tests {
		t.Run(tc.nama, func(t *testing.T) {
			sumber := tc.source
			tujuan := tc.target
			err := sumber.Transfer(&tujuan, tc.jumlah)
			if tc.expectErr && err == nil {
				t.Errorf("seharusnya error tapi tidak ada error")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("tidak boleh error, tetapi dapat: %v", err)
			}
			if sumber.Saldo != tc.expectSaldoS {
				t.Errorf("saldo pengirim = %.2f, ekspektasinya %.2f",
					sumber.Saldo, tc.expectSaldoS)
			}
			if tujuan.Saldo != tc.expectSaldoT {
				t.Errorf("saldo penerima = %.2f, ekspektasinya %.2f",
					tujuan.Saldo, tc.expectSaldoT)
			}
		})
	}
}

func BenchmarkByValue(b *testing.B) {
	r := Rekening{Nama: "Andi", Saldo: 10000.0, Active: true}
	for i := 0; i < b.N; i++ {
		processByValue(r)
	}
}

func BenchmarkByPointer(b *testing.B) {
	r := &Rekening{Nama: "Andi", Saldo: 10000.0, Active: true}
	for i := 0; i < b.N; i++ {
		processByPointer(r)
	}
}
