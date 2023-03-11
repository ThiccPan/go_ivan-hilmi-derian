# concurrent programming

## Concurrent vs Parallel

### Concurrent :
- berjalan secara bergantian
- contoh: pada queue, saat menjalankan suatu operasi, apabila operasi tersebut membutuhkan waktu yang relatif lama untuk menyelesaikannya, maka akan mengembalikan operasi tersebut ke queue dan menjalankan operasi lainnya terlebih dulu
- golang mendukung concurrency dengan fitur goroutine

### Parallel : 
- berjalan secara bersamaan
- contoh: menjalankan 2 operasi yang berbeda dalam satu waktu
- memanfaatkan multiple thread
- golang mendukung concurrency dengan fitur goroutine

## Concurrency in golang
golang memiliki fitur concurrency, yaitu goroutine. goroutine merupakan _mini thread_ yang berada di dalam thread
```go
func async() {
	fmt.Print("second")
}

func main() {
    //memanggil goroutine
	go cowSay()
    fmt.Print("first ")
    time.Sleep(5 * time.Second)
}
```
output: first second
## Channel
channel merupakan cara untuk melempar value dari 1 goroutine ke goroutine lainnya. channel yang telah dibuat harus diisi dengan value agar tidak terjadi deadlock. channel juga sebaiknya di close setelah dipanggil agar tidak mengakibatkan memory leak
```go
func main() {
	channel := make(chan string) // deklarasi channel
	defer close(channel) // close channel setelah func selesai
	
	go func() {
		time.Sleep(3 * time.Second)
		channel <- "hello from async anon func" // assign data ke channel
		fmt.Println("data sent successfully")
	}()

	time.Sleep(3 * time.Second)

	receiveStr := <- channel // assign value channel ke variabel normal
	fmt.Println(receiveStr)
}
```
channel mirip dengan async await dari js