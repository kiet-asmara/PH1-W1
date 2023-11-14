package main

import "fmt"

func main() {
	AlayGen("hello", "world", "internationalsx")
	fmt.Println(fibonacci(10))
}

func AlayGen(kata ...string) {
	kalimat := ""

	for _, k := range kata {
		kalimat += k + " "
	}

	kalimat2 := ""
	for _, h := range kalimat {
		switch h {
		case 'a':
			kalimat2 += "4"
		case 'e':
			kalimat2 += "3"
		case 'i':
			kalimat2 += "!"
		case 'l':
			kalimat2 += "1"
		case 'n':
			kalimat2 += "N"
		case 's':
			kalimat2 += "5"
		case 'x':
			kalimat2 += "*"
		default:
			kalimat2 += string(h)
		}
	}
	fmt.Println(kalimat2)
}

func fibonacci(n int) int {
	var fibo []int
	fibo = append(fibo, 0)
	fibo = append(fibo, 1)

	for i := 2; i <= n; i++ {
		fibo = append(fibo, 0)
		fibo[i] = fibo[i-1] + fibo[i-2]
	}

	// 0 termasuk
	return fibo[n-1]
}
