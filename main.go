package main

import (
	"asssignment1_golang/biodata"
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "1":
		fmt.Println("Absen No  : 1")
		biodata.Murid("Fajrian Nugraha", "Bandung", "Mahasiswa", "Karena Ingin menjadi Programmer")
		// fmt.Println(os.Args)
	case "2":
		fmt.Println("Absen No  : 2")
		biodata.Murid("Asep Darso", "Garut", "Kuli", "Karena Ingin menjadi Back End Programmer")
	case "3":
		fmt.Println("Absen No  : 3")
		biodata.Murid("Jajang Nurjaman", "Makasar", "Wiraswasta", "Karena Ingin Belajar Bahasa Pemrogramman Golang")
	case "4":
		fmt.Println("Absen No  : 4")
		biodata.Murid("Kakang Rudianto", "Bandung", "Pemaen Bola", "Karena Ingin menjadi Seorang Golang Developer")
	case "5":
		fmt.Println("Absen No  : 5")
		biodata.Murid("Icih Markicih", "Papua", "Mahasiswa", "Pengen Aja")
	default:
		fmt.Println("PERHATIAN!!")
		fmt.Println("Murid di Sekolah hanya ada 5 orang")
		os.Exit(1)
	}
}
