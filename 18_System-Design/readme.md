# System Design

## Diagram
Penggambaran suatu informasi menggunakan sebuah model sebagai representasi.  
Menggunakan beberapa model, yaitu:

### Flowchart
diagram yang menggambarkan proses/alur

### Use Case
diagram yang menggambarkan interaksi yang dapat dilakukan actor (pengguna/user) dengan sistem

### Entity Relationship
diagram yang menggambarkan relasi antara 1 entitas/objek dengan entitas lainnya dalam sistem tersebut

### High level Architecture
Menggambarkan desain sistem secara keseluruhan, seperti aplikasi/teknologi yang digunakan, dll

## Vertical dan Horizontal Scaling
Konsiderasi pada perancangan suatu sistem agar saat di masa depan ingin dikembangkan lebih lanjut, tidak perlu perombakan desain secara keseluruhan. Terdapat 2 Metode yaitu:  

### Vertical
Meningkatkan kapasitas dari sebuah sistem (menigkatkan ram, processing power, dsb)

### Horizontal
Meningkatkan jumlah infrastruktur yang dapat menjalankan sistem tersebut (mesin/komputer baru yang dapat melakukan host infrastruktur sistem)

### Karakteristik Scaling
- Reliability (Seberapa sering sistem mengalami kegagalan)
- Availablity (seberapa lama sistem dapat berjalan lancar berbanding dengan jumlah maintenance)
- Efficiency (Seberapa efisien sistem dalam mencapai tujuan, diukur menggunakan metriks waktu dan output per satuan waktu)
- Maintainability (Seberapa sulit sistem untuk diperbaiki apabila terjadi kegagalan)

## Job/Work queue
### Job queue
struktur data yang berisi daftar pekerjaan yang harus dilakukan sebuah sistem yang dimanage menggunakan job scheduler app
### Work queue
Kerangka yang digunakan untuk membangun kelompok worker dalam scala besar

## Load Balancing
tools yang dapat membantu mendistribusikan traffic pada kumpulan komponen sistem.  
dapat diletakkan di: 
- antara kumpulan web server dengan client
- antara kumpulan web server dengan app server
- antara kumpulan app server dengan db server

## Monolithic dan Microservice
Metode dalam membuat struktur aplikasi, monolithic menggabungkan sebuah aplikasi dalam satu codebase besar yang terdiri dari beberapa modul. sedangkan microservice memecah modul2 tersebut, dan melakukan komunikasi antara satu sama lain lewat network

## SQL dan NoSQL
### Relational (SQL)
- bersifat relasional
- terstruktur dengan schema2 tertentu
- punya standar yang jelas (SQL)
- banyak tool
- bersifat ACID (Atomic, Consistent, Isolated, Durable)
- cocok untuk data yang krusial/penting bagi business domain
- contoh: postgreSQL, mySQL, DB2, etc

### Schema-less (NoSQL)
- lebih fleksible
- mengurangi effort pada ACID, design schema, kompleksitas dari SQL
- transaksi ditangani aplikasi
- cocok untuk fast development, data logging, cache
- contoh: redis(key/value), cassandra(column - family), neo4j(graph), mongoDB(document)

## Caching
menyimpan data di memori untuk akses data yang lebih cepat dibanding data di storage. 

## Database replication
menduplikasi komponen kritis dari sistem sebagai bentuk backup apabila terjadi masalah/kerusakan

## Database indexing
metode untuk mengurangi banyaknya data yang diakses untuk mendapat data yang diquery/dicari, kebanyakan menggunakan b-tree indexing