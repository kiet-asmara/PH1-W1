package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// conditional 1
	var nama string
	fmt.Println("Masukkan nama:")
	fmt.Scan(&nama)

	hoki := rand.Intn(100-1) + 1
	fmt.Println(hoki)
	switch {
	case hoki > 80:
		fmt.Printf("Selamat %s, anda sangat beruntung", nama)
	case hoki <= 80 && hoki > 60:
		fmt.Printf("Selamat %s, anda beruntung", nama)
	case hoki >= 60 && hoki > 40:
		fmt.Printf("Mohon maaf %s, anda kurang beruntung", nama)
	case hoki < 40:
		fmt.Printf("Mohon maaf %s, anda sial", nama)
	}

	// conditional 2
	var nama2 string
	var umur int

	fmt.Println("\nMasukkan nama:")
	fmt.Scan(&nama2)
	fmt.Println("Masukkan umur:")
	fmt.Scan(&umur)

	if umur < 0 || umur > 100 {
		fmt.Println("ERROR invalid age")
		os.Exit(1)
	}

	switch {
	case umur > 18:
		fmt.Println("Silahkan masuk")
	case umur <= 18:
		fmt.Println("Dilarang masuk, maksimal umur 19")
	default:
		fmt.Println("ERROR bukan angka")
	}
}
