# Docker
## Container
instance imeage yang dapat dijalankan. abstraksi app layer untuk memudahkan proses deployment aplikasi. berdiri diatas container manager

## Perbedaan dengan VM
- Virtual Machine
    - abstraksi perangkat keras/hardware
    - terdapat copy dari operating system
    - booting proses lama
- Container
    - Abstraksi pada layer aplikasi
    - ukuran lebih kecil dari VM
    - dapat menghandle lebih banyak aplikasi

## Docker infrastructure
- menggunakan client server architecture
- docker client berbicara dengan server(daemon)
- dapat menggunakan docker compose sebagai alternatif untuk menjalankan beberapa container

## Docker client
beberapa perintah docker:
```bash
docker build
docker pull
docker run
FROM
RUN
ENV
ADD
WORKDIR
ENTRYPOINT
CMD
```
dsb

## Cara kerja
1. dockerfile mendeskripsikan tata cara build image
2. image berhasil dibuat, yang akan dibungkus dalam container berisi dependencies
3. upload ke docker hub (image repo)
4. download di host lain