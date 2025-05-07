package main

import "fmt"

// Ada dua cara declare variabel
// 1. pake var, kalau ini harus declare tipe datanya. Bisa declare di luar/dalam function. Bisa declare pake value atau nggak
// 2. pake short declaration: :=, ini ga perlu declare tipe datanya. Cuma bisa declare di dalam function

// Ini declare pake var di luar function
var x int = 10 // ini pake initial value
var name string = "Yusak"
var isActive bool // ini ga pake initial value

func main() {
	fmt.Printf("Nilai var x = %d, name = %s, isActive = %t\n", x, name, isActive)

	// ini declare pake var di dalam function
	var y int = 20
	fmt.Printf("Nilai y = %d\n", y)

	// ini declare pake short declaration
	otherName := "Vio"
	fmt.Printf("Nilai otherName = %s\n", otherName)

	// Latihan
	var age int = 35
	name := "Yusak"
	height := 172
	isMarried := false

	fmt.Println("Nama:", name)
	fmt.Println("Umur:", age)
	fmt.Println("Tinggi:", height)
	fmt.Println("Status Menikah:", isMarried)
}

// Tipe data dasar di Go: int, float, string, bool
// int zero valuenya 0
// float zero valuenya 0.0
// string zero valuenya ""
// bool zero valuenya false

// klo set pake int aja, maka dia akan menyesuaikan platformnya, 32 bit atau 64 bit. ini hubungannya sama arsitektur cpu nya
