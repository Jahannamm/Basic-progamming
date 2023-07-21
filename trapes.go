package main

import "fmt"

func hitungLuasTrapesium(alasAtas, alasBawah, tinggi float64) float64 {
	luas := 0.5 * (alasAtas + alasBawah) * tinggi
	return luas
}

func main() {
	var alasAtas, alasBawah, tinggi float64

	fmt.Println("Program Menghitung Luas Trapesium")
	fmt.Print("Masukkan panjang alas atas: ")
	fmt.Scan(&alasAtas)

	fmt.Print("Masukkan panjang alas bawah: ")
	fmt.Scan(&alasBawah)

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scan(&tinggi)

	luas := hitungLuasTrapesium(alasAtas, alasBawah, tinggi)
	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}
