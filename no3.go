package main

import "fmt"

type pegawai struct {
	nama        string
	gaji_pokok  int
	masa_kerja  int
	besar_bonus float64
}

func main() {
	var p pegawai

	fmt.Scan(&p.nama, &p.gaji_pokok, &p.masa_kerja)

	hitung_bonus(&p)

	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %.0f\n", p.nama, p.besar_bonus)
}

func hitung_bonus(p *pegawai) {
	var kali float64

	if p.masa_kerja >= 10 {
		kali = 1.5
	} else if p.masa_kerja >= 5 {
		kali = 0.75
	} else {
		kali = 0.5
	}
	p.besar_bonus = float64(p.gaji_pokok) * kali
}
