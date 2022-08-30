package biodata

import "fmt"

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func Murid(nama, alamat, pekerjaan, alasan string) {
	var m1 = Biodata{nama, alamat, pekerjaan, alasan}
	fmt.Println("Nama \t  :", m1.Nama)
	fmt.Println("Alamat \t  :", m1.Alamat)
	fmt.Println("Pekerjaan :", m1.Pekerjaan)
	fmt.Println("Alasan \t  :", m1.Alasan)
}
