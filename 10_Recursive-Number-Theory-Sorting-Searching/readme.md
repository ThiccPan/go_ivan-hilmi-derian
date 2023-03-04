## Recursive
function yang memanggil dirinya sendiri. Terdiri dari base case dan recursive case
```go
func recurse(a int) int {
	sum := 0
	iter := a
	// base case
	if sum == a {
		return 1
	}
	// pre
	iter--
	// recursion
	sum += recurse(a - 1)
	// post
	sum += 3
	fmt.Print(iter)
	return sum
}

func main() {
    fmt.Println(recurse(5))
}
// output: 
// 0 1 2 3 4
// 15
```
## number theory
cabang matematika yang mempelajari integer. beberapa contohnya adalah kpk, fpb, bilangan prima, dsb

## searching
metode pencarian data dalam suatu kelompok data (array, list, etc)  
### metode searching:
- linear search: mencari secara linear keseluruhan list data(O(N) time complexity)
- binary search: mulai dari tengah list/array, dan mengecilkan scope pencarian berdasarkan data yang dicari  
	- jika n > list[length/2], list = list[length/2:length], 
	- jika n < list[length/2], list = list[length/2:length]
	- scope diperkecil terus hingga n dapat ditemukan
- builtin search: dari std lib sort package (menggunakan binary search)

## sorting
metode pengurutan data dalam sebuah kelompok data.
### metode sorting
- selection
- counting
- merge
- builtin
- dsb