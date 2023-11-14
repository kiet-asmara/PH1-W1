package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// // Hello world
	// fmt.Println("Hello World")
	// fmt.Println("Motivasi saya masuk bootcamp adalah:\nmendapat kerja")

	// switch http response
	// response, err := http.Get("https://wijyasali.com")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// switch response.StatusCode {
	// case 200:
	// 	fmt.Println(http.StatusOK)
	// case 404:
	// 	fmt.Println(http.StatusNotFound)
	// case 500:
	// 	fmt.Println(http.StatusInternalServerError)
	// default:
	// 	fmt.Println(response.StatusCode)
	// }

	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
