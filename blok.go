package main

import "fmt"

func main() {
	data := map[string]int{
		"block rokan":   90000,
		"block mahakam": 50000,
	}

	for i, v := range data {
		if v > 80000 {
			fmt.Println(i, "Produksi tinggi di", v)
		} else {
			fmt.Println(i, "Produksi rendah untuk", v)
		}
	}
}
