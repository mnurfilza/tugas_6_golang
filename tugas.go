package tugas6

import "fmt"

type Mahasiswa struct {
	Nama string
	Umur int
}

func (m *Mahasiswa) Tugas6() {
	fmt.Println("Nama :", m.Nama)
	fmt.Println("Umur :", m.Umur, "Tahun")
}
