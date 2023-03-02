package main

import (
	"fmt"
	"strings"
)

// jumlah offset
const key = 5

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	// your code here
	decoder := func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			r = ((r - 'a' + key) % 26) + 'a'
			return r
		case r >= 'A' && r <= 'Z':
			r = ((r - 'A' + key) % 26) + 'A'
			return r
		}
		return r
	}
	nameEncode = strings.Map(decoder, s.name)
	s.nameEncode = nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	// your code here
	decoder := func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			r = ((r - 'a' + (26 - key)) % 26) + 'a'
			return r
		case r >= 'A' && r <= 'Z':
			r = ((r - 'A' + (26 - key)) % 26) + 'A'
			return r
		}
		return r
	}
	nameDecode = strings.Map(decoder, s.name)
	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student's name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student's name " + a.name + " is : " + c.Decode())
	}
}
