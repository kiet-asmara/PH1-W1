package main

import "fmt"

func main() {
	// variabel 1
	var myNum int32 = 50
	fmt.Println(myNum)

	var myNum2 float32 = 51
	fmt.Println(myNum2)

	var myNumStr string = "50"
	fmt.Println(myNumStr)

	// variabel 2
	var x int32 = 5
	var t int32 = 10
	z := x + t
	fmt.Println(z)

	// CLI
	var input string
	fmt.Println("please input your name")
	fmt.Scan(&input)

	fmt.Println("hello", input)

	// Array & Slice 1
	var people = []string{"Walt", "Jesse", "Skyler", "Saul"}

	fmt.Println("panjang slice people", len(people))
	fmt.Println(people)

	fmt.Println("Menambahkan Hank dan Marie...")
	people = append(people, "Hank", "Marie")
	fmt.Println("panjang slice people", len(people))
	fmt.Println(people)

	// Array & Slice 2
	var peopleMaps []map[string]string
	peopleMaps = append(peopleMaps, map[string]string{"name": "Hank", "gender": "M"})
	peopleMaps = append(peopleMaps, map[string]string{"name": "Heisenberg", "gender": "M"})
	peopleMaps = append(peopleMaps, map[string]string{"name": "Skyler", "gender": "F"})
	for _, person := range peopleMaps {
		if person["gender"] == "M" {
			person["name"] = "Mr." + person["name"]
		} else {
			person["name"] = "Mrs." + person["name"]
		}
	}
	fmt.Println(peopleMaps)
}
