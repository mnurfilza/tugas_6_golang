package main

import (
	"belajargolang/tugas6"
	"fmt"
)

func main() {
	fmt.Println("Ini Tugas Ke 6")
	temp := &tugas6.Mahasiswa{Nama: "Filza", Umur: 23}
	temp.Tugas6()
}
