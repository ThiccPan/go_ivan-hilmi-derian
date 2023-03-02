# String, Advance Function, Pointer, Package, & Error Handling

## String
mengolah string dapat dilakukan dengan berbagai cara, salah satunya adalah string package. Package ini memberikan beberapa utility function untuk memproses string.

### String Helper: 
- Len: mengetahui panjang string
- Compare: komparasi 2 string dengan operator '=='
- Contains: mengetahui apakah suatu substring terdapat di string lain
- Substring: memotong string menggunakan ':' pada index yang diinginkan
- Replace: mengubah substr pada string
- Insert: memasukkan string lain ke dalam deklarasi suatu string

## Advanced Function  
### Variadic Function
menambahkan titik 3 (...) pada function parameter untuk menampung banyak variable:
```go
func sum(nums ...int) int {
    // your code here
}

// call your func
sum(a, b, c, d) // however many var you want
```

### Anonymous Function
function tanpa nama, bisa dialokasikan ke sebuah variable (Higher order function)
```go
  func() {
    fmt.Println("Welcome! to GeeksforGeeks")
  }()

  // Assigning an anonymous function to a variable
  value := func() {
    fmt.Println("Welcome! to GeeksforGeeks")
  }
  value()

  // Passing arguments in anonymous function
  func(sentence string) {
    fmt.Println(sentence)
  }("GeeksforGeeks")
```

### Closure Function
anonymous function yang merefer pada sebuah variable di luar lingkupnya
```go
func incremetCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

### Defer
perintah untuk menunda berjalannya suatu function setelah lingkup berjalannya function tersebut selesai mengembalikan nilai. urutan berjalan adalah LIFO (seperti stack)
```go
func main() {
    defer func() {
        fmt.Println("ditunda")
    }()
    fmt.Println("duluan")
    /*** 
    output: 
    duluan
    ditunda
    */
}
```

## Pointer
Mereferensi alamat memori (memory address) dari sebuah variable yang telah dideklarasikan sebelumnya.
```go
var a := "abcd"
var b := &a
fmt.Println(b) // alamat memori dari a
fmt.Println(*b) // "abcd"
// jika reassign b, maka value a juga akan berubah
*b = "dcba"
// a = "dcba"
```

## Error Handling
untuk menghandle error, dapat menggunakan tipe data error dan package "errors". Contohnya pada sebuah function, bisa mereturn 2 nilai, yaitu nilai yang ingin di return dan juga error yang tertrigger di dalam function. bisa juga menggunakan function panic() untuk menandakan terjadinya error, dan recover() untuk mengatasi panic call

## interface
kumpulan method yang dapat diimplementasikan ke dalam suatu method. salah satu usecase adalah untuk mengatur akses dari method suatu struct data.
```go
// implemntasi interface
type calculate interface {
    large() int
}

type square struct {
    side int
}

func (s square) large() int {
    return s.side * s.side
}

func main() {
    var dimResult calculate
    dimResult = square{10}
    fmt.Println("large square :", dimResult.large())
}
```
Juga dapat digunakan sebagai dynamic value dan type assertion
```go
func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
    // type assertion
    var i interface
    describe(i) // (nil, nil)
    i = 12
    describe(i) // (12, int)
    i = "abcd"
    describe(i) // (abcd, string)
}
```

## Package
cara untuk menggunakan variable/func/struct/etc dari luar lokasi file tersebut. Untuk membuat suatu function/data dapat diakses di luar file, penulisan function/data tersebut diawali dengan huruf kapital