package main

import (
	"fmt"
	"math"
)

func calOverPriceBaggage(userBaggage int, domesticPrice int, internationalPrice int, flightType string) int {
	var	freeQuotaBaggage int = 50
	overBaggage := userBaggage - freeQuotaBaggage
	if overBaggage <= 0 {
		return 0
	}

	// Harga berdasarkan jenis penerbangan
	var pricePerKg int
	if flightType == "domestic" || flightType == "domestik" {
		pricePerKg = domesticPrice
	} else if flightType == "international" || flightType == "internasional" {
		pricePerKg = internationalPrice
	}

	// Hitung harga bagasi
	totalPrice := overBaggage * pricePerKg
	return totalPrice
}

func calFlightTime(fromCountry string, destinationCountry string) int {
	fromIndex := -1
	toIndex := -1
	worldMap := []string{"PH", "HD", "ID", "SG", "MY", "VN", "KH"}

	// Cari indeks negara asal dan negara tujuan
	for i, country := range worldMap {
		if country == fromCountry {
			fromIndex = i
		}
		if country == destinationCountry {
			toIndex = i
		}
	}

	// Jika tidak ditemukan
	if fromIndex == -1 || toIndex == -1 {
		return 0
	}

	// Hitung lama penerbangan
	distance := toIndex - fromIndex
	if distance < 0 {
		distance = int(math.Abs(float64(distance)))
	}

	return distance
}

func main() {
	var userBaggage, domesticPrice, internationalPrice int
	var flightType string

	// Input bagasi dari user
	fmt.Print("Masukkan berat bagasi Anda (KG): ")
	fmt.Scan(&userBaggage)
	fmt.Print("Harga per KG untuk penerbangan domestik: ")
	fmt.Scan(&domesticPrice)
	fmt.Print("Harga per KG untuk penerbangan internasional: ")
	fmt.Scan(&internationalPrice)
	fmt.Print("Jenis penerbangan ('domestic' atau 'international'): ")
	fmt.Scan(&flightType)
	
	// Hitung bagasi
	totalPrice := calOverPriceBaggage(userBaggage, domesticPrice, internationalPrice, flightType)
	fmt.Print("Total biaya bagasi: Rp ", totalPrice)
	fmt.Print("\n")
	
	// Input negara dari user
	var fromCountry, destinationCountry string
	fmt.Print("Masukkan negara asal (kode): ")
	fmt.Scan(&fromCountry)
	fmt.Print("Masukkan negara tujuan (kode): ")
	fmt.Scan(&destinationCountry)

	// Hitung waktu penerbangan
	flightTime := calFlightTime(fromCountry, destinationCountry)
	if flightTime == 0 {
		fmt.Println("Negara tidak ditemukan dalam peta penerbangan.")
	} else {
		fmt.Print("Penerbangan dari " ,fromCountry, " ke ", destinationCountry, " adalah ",flightTime," jam")
	}}
