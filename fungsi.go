package main

import "fmt"

func fetchData(id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("ID tidak boleh kosong")
	}
	return "Data untuk ID " + id, nil // 'nil' berarti tidak ada error
}

func main() {
	fmt.Println(fetchData("1"))
	fmt.Println(fetchData(""))
}
