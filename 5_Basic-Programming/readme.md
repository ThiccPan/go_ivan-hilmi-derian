# Basic Programming Summary

## 1. Golang Introduction
Golang adalah bahasa pemrograman multi purpose yang dikembangkan dengan tujuan agar memiliki performa yang cepat seperti c/c++, dengan sintaks yang mudah dipahami seperti python. Golang dapat digunakan untuk mengembangkan program aplikasi ataupun program sistem (system program)

## 2. Sintaks utama
- package : untuk mendeklarasikan package file
- import : daftar package yang diimport dalam file tsb
- func main : function utama yang pertama kali dijalankan saat running aplikasi 

## 3. Golang terminal command 
- go run : menjalankan program
- go build : compile program
- go install : compile dan install program
- go test : melakukan testing program
- go get : download package go

## 4. Variable and type
var adalah cara untuk menyimpan sebuah nilai dengan tipe data tertentu

### Tipe data golang
- boolean : true or false
- numeric : int, float, complex
- string

### Deklarasi variable
var - [nama variable] - [tipe data]  
(eg: var age int)

### Zero value in variable
merupakan nilai default dari suatu variable ketika pertama kali dideklarasi tanpa melakukan assign nilai ke var tersebut
- boolean : false
- numeric : 0/0.0
- string : ""
- lainnya : nil

### constant
menyimpan nilai konstant yang tidak dapat diubah setelah dideklarasikan  
const - [nama constant] - [tipe constant]

## 5. operator
### aritmathic
- (+) : addition
- (-) : subtraction
- (/) : division
- (*) : multiplication
- (%) : modulo
- (++) : increment
- (--) : decrement

### comparison
==, !=, >, <, >=, <=

### logical
&&, ||, !

### bitwise
&, !, ^, etc

### assignment
operator lain ditambah dengan equal/sama dengan (+=, -=, etc)

### pointer
&, *

## 6. branching
bisa menggunakan if atau switch case  

if .... {

}

switch ...{  
case ... :  
    ........  
case ... :  
    ........  
default :  
    ........  
}

## 7. Looping
hanya dapat menggunakan for, untuk alternatif while loop dapat menggunakan for tanpa parameter
