package main

import (
	"fmt"
	"math"
)

type Book struct {
	judul   string
	penulis string
	harga   float64
}

func main() {
	// soal 1
	// Buatlah program Golang sederhana yang mendeklarasikan dua variabel: nama bertipe string dan umur bertipe int, kemudian cetak ke konsol.
	nama := "joni"
	umur := 20

	fmt.Println(nama)
	fmt.Println(umur)

	// soal 2
	// Buatlah sebuah map yang memiliki key berupa string (nama buah) dan value berupa float64 (harga per kilogram).
	// Isi map tersebut dengan 3 data, lalu tampilkan datanya.
	buah := map[string]float64{
		"Mangga": 20.12,
		"Jambu":  16.20,
		"Jeruk":  13.20,
	}

	fmt.Println(buah)

	// soal 3
	// Gunakan struktur pengkondisian untuk memeriksa apakah suatu bilangan adalah bilangan genap atau ganjil.
	bil := 9

	if bil%2 == 0 {
		fmt.Println("Bilangan adalah genap")
	} else {
		fmt.Println("Bilangan adalah ganjil")
	}

	// soal 4
	// Tuliskan sebuah program yang menerima input angka bulan (1-12), dan mengembalikan nama bulan tersebut.
	// Gunakan switch case untuk mencapainya.
	var bulan int
	fmt.Println("Masukkan angka bulan:")
	fmt.Scanln(&bulan)

	switch bulan {
	case 1:
		fmt.Println("Januari")
	case 2:
		fmt.Println("Februari")
	case 3:
		fmt.Println("Maret")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("Mei")
	case 6:
		fmt.Println("Juni")
	case 7:
		fmt.Println("Juli")
	case 8:
		fmt.Println("Agustus")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("Oktober")
	case 11:
		fmt.Println("Nopember")
	case 12:
		fmt.Println("Desember")
	default:
		fmt.Println("Tidak ada bulan dengan angka itu.")
	}

	// soal 5
	// Tuliskan sebuah program yang mencetak angka dari 1 hingga 10 menggunakan loop.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// soal 6
	// Buatlah program yang mencetak N suku pertama dari sebuah deret geometri dengan rasio 2.
	var n int
	fmt.Println("Masukkan panjang deret geometri 2:")
	fmt.Scanln(&n)
	var r float64
	for i := 0; i < n; i++ {
		r = math.Pow(2, float64(i))
	}
	fmt.Println(r)

	// soal 7
	// Gunakan perulangan dan pengkondisian untuk mencetak angka dari 1 hingga 20,
	// tetapi cetak "Fizz" untuk angka yang habis dibagi 3 dan "Buzz" untuk angka yang habis dibagi 5.
	for i := 1; i <= 20; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// soal 8
	// Tuliskan sebuah program yang mencetak bilangan dari 1 sampai N, tetapi untuk kelipatan 3 cetak "Fizz", untuk kelipatan 5 cetak "Buzz", dan untuk kelipatan keduanya cetak "FizzBuzz".
	var fb int
	fmt.Println("Masukkan angka fizzbuzz:")
	fmt.Scanln(&fb)
	for i := 1; i <= fb; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// soal 9
	// Buatlah sebuah program yang mencetak pola kotak dengan lapisan luar terdiri dari karakter *,
	// sedangkan bagian dalamnya kosong/hollow (menggunakan karakter spasi), dengan ukuran N x N.
	var st int
	fmt.Println("Masukkan dimensi kotak kosong:")
	fmt.Scanln(&st)

	for i := 1; i <= st; i++ {
		if i == 1 || i == st {
			for j := 1; j <= st; j++ {
				fmt.Print("*")
			}
		} else {
			for j := 1; j <= st; j++ {
				if j == 1 || j == st {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")
	}

	// soal 10
	// Tuliskan sebuah program yang menggunakan nested loop untuk menghasilkan segitiga Pascal dengan tinggi 5 baris.
	for i := 0; i < 5; i++ {
		if i == 0 {
			fmt.Println(1)
		} else if i == 1 {
			fmt.Println("1 1")
		} else {
			t := 1
			for j := 0; j <= i; j++ {
				if j == 0 {
					fmt.Print(1, " ")
				} else if j == i {
					fmt.Print(1, "\n")
				} else {
					t = t * (i - j + 1) / j
					fmt.Print(t, " ")
				}
			}

		}
	}

	// soal 11
	// Buatlah sebuah fungsi yang menerima dua bilangan dan mengembalikan jumlah dari kedua bilangan tersebut..
	fmt.Println(duaBil(9, 2))

	// soal 12
	// Buatlah sebuah fungsi yang menerima dua argumen (bilangan bulat) dan mengembalikan hasil penjumlahan dari kedua argumen tersebut.
	fmt.Println(duaBilBulat(9, 2))

	// soal 13
	// Buatlah sebuah fungsi variadic yang menerima sejumlah bilangan bulat dan mengembalikan jumlah dari semua bilangan tersebut.
	fmt.Println(jumlahSemua(1, 2, 3, 4, 5))

	// soal 14
	// Buatlah sebuah fungsi variadic yang menerima sejumlah bilangan bulat dan mengembalikan nilai maksimum dari sekumpulan bilangan tersebut.
	fmt.Println(maxSemua(1, 2, 3, 4, 5))

	// soal 15
	// Anda memiliki struktur data yang mencakup informasi tentang buku yang terdiri dari judul, penulis, dan harga.
	// Buatlah struct yang merepresentasikan struktur data ini dan buat fungsi yang akan menerima slice dari struct ini dan mengembalikan total harga dari semua buku.
	buku := []Book{
		{judul: "harry potter", penulis: "JK Rowling", harga: 100.20},
		{judul: "naruto", penulis: "sasuke", harga: 150},
		{judul: "ramayana", penulis: "wisnu", harga: 200.12},
	}

	fmt.Printf("%.2f\n", totalHarga(buku))

	// soal 16
	/*
		Sebuah perusahaan ingin memberikan bonus kepada karyawannya berdasarkan performa.
		Buatlah program yang menghitung bonus karyawan dengan menggunakan switch case yang mencakup kondisi sebagai berikut:
		A: Bonus 20% dari gaji
		B: Bonus 10% dari gaji
		C: Bonus 5% dari gaji
		D: Tidak ada bonus
	*/
	fmt.Println("Bonus anda:")
	fmt.Println(hitungBonus("A", 100))
	fmt.Println(hitungBonus("B", 100))
	fmt.Println(hitungBonus("C", 100))

	// soal 17
	// Buatlah program yang mencetak angka dari 1 hingga 100, namun untuk kelipatan 3 cetak "Fizz",
	// untuk kelipatan 5 cetak "Buzz", dan untuk kelipatan dari 3 dan 5 cetak "FizzBuzz".
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// soal 18
	// Buatlah program yang mencetak pola berlian dengan bintang (*) di tengah-tengah layar dengan menggunakan nested loop.
	// Tinggi berlian (jumlah baris) adalah n (sebuah bilangan ganjil).
	berlian := 9

	median := (berlian + 1) / 2
	for i := 0; i < median; i++ {
		for j := 1; j <= berlian; j++ {
			if j >= median-i && j <= median+i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	for i := berlian - median - 1; i >= 0; i-- {
		for j := 1; j <= berlian; j++ {
			if j >= median-i && j <= median+i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

	// soal 19
	// Buatlah fungsi variadic yang menerima sejumlah kata (string) dan mengembalikan sebuah kalimat di mana kata-kata tersebut disatukan dengan spasi.
	// Fungsi juga harus mengembalikan jumlah kata yang telah disatukan.
	fmt.Println(buildSentence("halo", "namaku", "joni", "jani"))

	// soal 20
	/*
		Seorang guru ingin membuat sebuah program yang bisa membantu dia dalam mengevaluasi kinerja siswa di kelasnya.

		Program tersebut harus mampu melakukan hal berikut:
		Iterasi dari 1 hingga n, di mana n adalah jumlah siswa di kelas.
		Jika nilai siswa i kurang dari 50, cetak "Siswa i mendapatkan predikat D".
		Jika nilai siswa i berada di antara 50 dan 69, cetak "Siswa i mendapatkan predikat C".
		Jika nilai siswa i berada di antara 70 dan 84, cetak "Siswa i mendapatkan predikat B".
		Jika nilai siswa i lebih dari atau sama dengan 85, cetak "Siswa i mendapatkan predikat A".
	*/
	nilaiSiswa := []int{85, 60, 78, 92, 45, 73}
	evaluasiKinerjaSiswa(nilaiSiswa)

}

// soal 11
func duaBil(a float64, b float64) float64 {
	return a + b
}

// soal 12
func duaBilBulat(a int, b int) int {
	return a + b
}

// soal 13
func jumlahSemua(s ...int) int {
	t := 0
	for _, v := range s {
		t += v
	}
	return t
}

// soal 14
func maxSemua(s ...int) int {
	m := 0
	for _, v := range s {
		if v > m {
			m = v
		}
	}
	return m
}

// soal 15
func totalHarga(buku []Book) float64 {
	var t float64 = 0
	for _, b := range buku {
		t += b.harga
	}
	return t
}

// soal 16
func hitungBonus(performa string, gaji float64) float64 {
	var bonus float64
	switch performa {
	case "A":
		bonus = 0.2 * gaji
	case "B":
		bonus = 0.1 * gaji
	case "C":
		bonus = 0.05 * gaji
	case "D":
		bonus = 0
	}
	return bonus
}

// soal 19
func buildSentence(words ...string) (string, int) {
	s := ""
	for _, w := range words {
		s += w + " "
	}
	return s, len(words)
}

// soal 20
func evaluasiKinerjaSiswa(nilaiSiswa []int) {
	for i, n := range nilaiSiswa {
		if n < 50 {
			fmt.Printf("Siswa %d mendapatkan predikat D.\n", i+1)
		} else if n <= 69 {
			fmt.Printf("Siswa %d mendapatkan predikat C.\n", i+1)
		} else if n <= 84 {
			fmt.Printf("Siswa %d mendapatkan predikat B.\n", i+1)
		} else if n >= 85 {
			fmt.Printf("Siswa %d mendapatkan predikat A.\n", i+1)
		}
	}
}
