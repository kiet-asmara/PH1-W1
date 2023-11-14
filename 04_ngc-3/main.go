package main

import (
	"fmt"
	"sort"
)

func main() {
	// looping 1
	var people []map[string]string
	people = append(people, map[string]string{"name": "Hank", "age": "50", "Job": "Polisi"})
	people = append(people, map[string]string{"name": "Heisenberg", "age": "M", "Job": "Ilmuwan"})
	people = append(people, map[string]string{"name": "Skyler", "age": "F", "Job": "Akuntan"})

	for _, maps := range people {
		fmt.Printf("Hi Perkenalkan, Nama saya %s, umur saya %s, dan saya bekerja sebagai %s\n", maps["name"], maps["age"], maps["Job"])
	}

	// looping 2
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	count(slice1)
	count(slice2)

	// logic 1 - palindrome
	kata := "door"
	fmt.Println(palindrome(kata))

	// logic 2 - xoxo
	test := "xoxoxo"
	fmt.Println(xoxo(test))

	// logic 3 - sort
	numbers := []int{9, 4, 2, 5, 4, 6, 1, 7}
	fmt.Println(bubble(numbers))

	// logic 4 - asterisk 1
	asterisk1(5)

	// logic 5 - asterisk 2
	asterisk2(5)

	// logic 5 - asterisk 3
	asterisk3(5)

	// logic 5 - asterisk 4
	asterisk4(5)
}

func count(slice []float64) {
	var avg, sum, med float64
	length := len(slice)

	for _, n := range slice {
		sum += n
	}
	avg = sum / float64(length)
	sort.Float64s(slice)
	if length%2 == 0 {
		med = (slice[length/2-1] + slice[length/2]) / 2
	} else {
		med = slice[length/2]
	}
	fmt.Printf("sum = %f, avg = %f, med = %f\n", sum, avg, med)
}

func palindrome(kata string) bool {
	result := ""

	for _, char := range kata {
		result = string(char) + result
	}
	if result == kata {
		return true
	} else {
		return false
	}
}

func xoxo(word string) bool {
	xCount := 0
	oCount := 0

	for _, l := range word {
		if l == 'x' {
			xCount += 1
		} else if l == 'o' {
			oCount += 1
		}
	}

	if xCount == oCount {
		return true
	} else {
		return false
	}
}

func bubble(s []int) []int {
	for range s {
		for j := 0; j < len(s)-1; j++ {
			if s[j] < s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func asterisk1(rows int) {
	for i := 1; i <= rows; i++ {
		fmt.Println("*")
	}
}

func asterisk2(rowcl int) {
	for i := 1; i <= rowcl; i++ {
		for j := 1; j <= rowcl; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func asterisk3(rowcl int) {
	for i := 1; i <= rowcl; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func asterisk4(rowcl int) {
	for i := 1; i <= rowcl; i++ {
		for j := 1; j <= rowcl-i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
