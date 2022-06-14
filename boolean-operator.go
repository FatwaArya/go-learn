package main

import (
	"fmt"
)

func main() {
	var nilaiAkhir = 80
	var absensi = 80

	var lulusUjian = nilaiAkhir >= 80
	var hadir = absensi >= 75

	var lulus = lulusUjian && hadir
	fmt.Println("Lulus = ", lulus)
}
