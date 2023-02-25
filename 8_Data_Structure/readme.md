# Data Structure in Golang

## Array
Sekumpulan data yang memiliki tipe tertentu yang sama (int, string, bool, etc) yang memiliki ukuran dan kapasitas yang statis

### deklarasi
- [length]type

### inisialisasi
- arrayName[index] = 

## Slice
Memiliki struktur yang mirip seperti array, namun dengan ukuran yang dinamis. Secara low level cara kerjanya mirip dengan arraylist di java maupun array di javascript. memiliki kapasitas dan panjang

### deklarasi
- make([]Type)

### inisialisasi
- sliceName[index] = value
- append(sliceName, appendedValue...) (melakukan append apabila melebihi kapasitas)

## Map
tipe data key:value pair. Key pada map harus bersifat unique/tidak boleh ada yang sama nilainya.

### deklarasi
- make(map[keyType]valueType)

### inisialisasi
- map[key] = value (dapat menggunakan 2 assignment variable untuk mengecek value dan cek apakah map dengan key telah diassign)

## length and capacity
- length : panjang suatu tipe data
- mengecek panjang suatu array dan map menggunakan metode len()  
- capacity : reserve memory suatu data (hanya untuk slice)
- dapat mengecek capacity menggunakan method cap() (khusus tipe data slice)

## function
kumpulan kode dipisahkan ke dalam blok tertentu. function membantu dalam membuat penulisan kode yang modular dan readable  

Function dalam golang merupakan first class citizen, sehingga dapat dioper referensinya ke variable lain. contoh:
```go
func add(a, b int) int {
    return a+b
}

passAdd := add

fmt.Println(passAdd(1, 2)) //3
```

note: function dapat dipanggil secara rekursif, namun hati-hati karena dapat menyebabkan stack overflow jika function dipanggil terlalu dalam