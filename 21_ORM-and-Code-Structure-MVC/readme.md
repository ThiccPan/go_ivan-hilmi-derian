# ORM dan struktur MVC
## ORM
- Object relational mapping, layer abstraksi dari kode ke fungsionalitas database
- tools untuk membantu dalam melakukan operasi CRUD ke dalam database
- kelebihan: mempercepat proses development
- kekurangan: menambah layer abstraksi(lower performance)
- package ORM dalam bahasa go: GORM


## MVC
- Struktur kode untuk memisahkan bagian berdasarkan fungsi
- terdiri dari model, view, dan controller
- Model:
    - Representasi objek/struktur data
    - letak business logic
- View:
    - tampilan yang dilihat oleh user
    - menyajikan data yang direquest
- Controller"
    - mengirimkan data antara model dengan view
    - biasanya memiliki jumlah kode yang paling banyak
- Kelebihan: separasi kode berdasarkan concern, memudahkan refactoring, dsb
