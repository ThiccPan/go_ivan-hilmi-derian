# Compute Service

## Strategi Deployment
### Big bang
- Langsung mereplace keseluruhan aplikasi di server
- kelebihan: mudah untuk dlakukan, perubahan secara instan
- kekurangan: beresiko, downtime yang cukup lama

### Rollout
- mereplace aplikasi di server ke versi terbaru satu demi satu
- Kelebihan: lebih aman, downtime lebih cepat dari big bang
- Kekurangan: terdapat 2 versi dari aplikasi di server, tidak ada control request

### Blue/Green
- deploy aplikasi versi terbaru ke server baru
- load balancer kemudian melakukan switch dari server dengan aplikasi lama ke server dengan aplikasi baru
- kelebihan: versi aplikasi sama di seluruh server, perubahan terjadi dengan cepat
- kekurangan: resource yang dibutuhkan lebih banyak, testing aplikasi harus benar2 maksimal

### Canary
- mirip dengan blue/green, namun switch dilakukan secara progressif, mengalihkan sekian persen request ke aplikasi baru
- Kelebihan: cukup aman, rollback cukup mudah dilakukan
- kekurangan: untuk mengalihkan seluruh request hingga 100% ke aplikasi baru butuh waktu cukup lama

## AWS
### EC2
- layanan compute platform yang disediakan oleh aws
- Virtual machine yang dapat digunakan sebagai platform deployment aplikasi
- memiliki konfigurasi yang beragam
- dapat melakukan ssh ke instance ec2 untuk mengakses instance tersebut, memerlukan pem key

### RDS
- layanan database yang disediakan oleh aws
- beberapa konfigurasi telah dilakukan oleh pihak aws, sehingga memudahkan user
- memiliki berbagai macam fitur seperti high availability, scalability, dsb
- memiliki beragam jenis database

### Security group
- lingkup security yang dapat kita konfigurasi untuk mengatur akses keluar masuk dari resource yang dimiliki di AWS
- inbound rule untuk mengatur koneksi ke instance
- outbound rule untuk mengatur koneksi dari instance

## Link resource
- docker image: https://hub.docker.com/repository/docker/ivanhd/prak-26/general
- dev dan deployment docker-compose file: di folder praktikum