package main

import "fmt"

func main() {
	/**
	* Type Data
	**/

	var status string = "tidak aman"
	// status := "aman"

	if status == "aman" {
		fmt.Println("Operasi normal")
	} else {
		fmt.Println("Peringatan!")
	}

	/**
	* loop
	**/

	var data [3]string = [3]string{"minyak", "gas", "batu bara"}
	// data := []string{"minyak", "gas", "batu bara"}

	for index, value := range data {
		fmt.Println("index = ", index, ", value = ", value)
	}

	for i := 0; i < len(data); i++ {
		fmt.Println("i = ", i, ", v = ", data[i])
	}

	idx := 0
	for idx < len(data) {
		fmt.Println(data[idx])
		idx++
	}

	for i, val := range data {
		switch i {
		case 0:
			fmt.Println("bentuk ", val, " cair")
		case 1:
			fmt.Println("bentuk ", val, " ya tetep gas wkwk")
		case 2:
			fmt.Println("bentuk ", val, " keras")
		}
	}

	person := map[string]interface{}{"name": "hans", "age": 100}
	fmt.Println(person["name"], person["age"])
}
