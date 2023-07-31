package entity

import (
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	Absen     string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (d Data) Showdata() {
	fmt.Printf("Data teman dengan absen %v :\n", d.Absen)
	fmt.Printf("Nama: %v\nAlamat: %v\nPekerjaan: %v\nAlasan: %v\n", d.Nama, d.Alamat, d.Pekerjaan, d.Alasan)
}

func Soal1() {
	for {
		inputStr := os.Args[1]
		_, err := strconv.Atoi(inputStr)

		if err != nil {
			fmt.Println("Invalid input, please enter a number")
			continue
		}

		data_output := Data{
			Absen:     inputStr,
			Nama:      "Teman " + inputStr,
			Alamat:    "Alamat Teman " + inputStr,
			Pekerjaan: "Pekerjaan Teman " + inputStr,
			Alasan:    "Alasan memilih golang karena alasan " + inputStr,
		}

		data_output.Showdata()
		break
	}

}
