package main

import (
	"fmt"
	"time"
)

type Keuangan struct {
	masuk      int
	keluar     int
	created_at time.Time
}

func (s Keuangan) sumMasuk(database []Keuangan) {
	var sum int
	for _, value := range database {
		sum = sum + value.masuk
	}
	fmt.Println(sum)
}

func (s Keuangan) sumKeluar(database []Keuangan) {
	var sum int
	for _, value := range database {
		sum = sum + value.keluar
	}
	fmt.Println(sum)
}

var database []Keuangan

func main() {
	input := 99
	var uang Keuangan
	for input != 0 {
		fmt.Println("==============")
		fmt.Println("Aplikasi Keuangan")
		fmt.Println("==============")
		fmt.Println("1 pemasukan")
		fmt.Println("2 pengeluaran")
		fmt.Println("3 seluruh data")
		fmt.Println("4 sum pemasukan")
		fmt.Println("5 sum pengeluaran")
		fmt.Println("0 exit")
		fmt.Println("==============")
		var inputs int
		fmt.Scanln(&inputs)
		switch inputs {
		case 1:
			fmt.Scanln(&inputs)
			database = append(database,
				Keuangan{
					inputs,
					0,
					time.Now(),
				})
		case 2:
			fmt.Scanln(&inputs)
			database = append(database,
				Keuangan{
					0,
					inputs,
					time.Now(),
				})
		case 3:
			fmt.Println(database)
		case 4:
			uang.sumMasuk(database)
		case 5:
			uang.sumKeluar(database)
		}
	}
}
