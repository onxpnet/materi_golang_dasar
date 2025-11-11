package main

import (
	"fmt"
	"time"
)

type SensorData struct {
	ID        string
	Lokasi    string
	Value     float64
	Timestamp time.Time
}

func (s SensorData) UpdateValue(newValue float64) {
	s.Value = newValue // Memodifikasi data asli
}

func (s SensorData) UpdateLokasi(newLokasi string) {
	s.Lokasi = newLokasi
}

func main() {
	sensor := SensorData{
		"1",
		"Jakarta",
		100.5,
		time.Now(),
	}

	sensor = SensorData{
		"2", "Pandeglang", 12.3, time.Now(),
	}

	sensor.UpdateValue(1.1)
	sensor.UpdateLokasi("Jepang")

	fmt.Println(sensor)

	var pemasukan string
	fmt.Scanln(&pemasukan)
	fmt.Println("input data ", pemasukan, "telah di simpan")
}
