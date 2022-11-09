package main

import "fmt"

func main() {
	data1 := map[string][]interface{}{
		"Nilai A": []interface{}{75, "Budi"},
		"Nilai B": []interface{}{100},
		"Nilai C": []interface{}{80, "Andi"},
	}
	data2 := map[string][]interface{}{
		"x": []interface{}{"Ulang Harian", 75, 80, 100},
	}

	m := make(map[string]map[string][]interface{})
	m["data1"] = data1
	m["data2"] = data2

	s := fmt.Sprintf("%v", m)
	fmt.Println(s)
}
