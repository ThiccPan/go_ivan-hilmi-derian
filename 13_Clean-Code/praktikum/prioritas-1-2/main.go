package main

import "fmt"

type Kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}
/*
terdiri dari anonymous field bertipe kendaraan
berisi field: 
	- totalRoda bertipe int
	- kecepatanPerJam bertipe int
*/
type Mobil struct {
	Kendaraan
}

/*
menambah kecepatanPerJam sebanyak 10
*/
func (m *Mobil) Berjalan() {
	m.Mempercepat(10)
}

/*
menambah kecepatan dari mobil berdasarkan parameter
parameter :
	- tambahan bertipe int, tambahan kecepatan dari mobil
*/
func (m *Mobil) Mempercepat(tambahan int) {
	m.kecepatanPerJam = m.kecepatanPerJam + tambahan
}

func main() {
	mobilcepat := Mobil{}
	mobilcepat.Berjalan()
	mobilcepat.Berjalan()
	mobilcepat.Berjalan()
	fmt.Println(mobilcepat)

	mobillamban := Mobil{}
	mobillamban.Berjalan()
}