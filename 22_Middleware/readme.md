# Middleware
## Definisi
function yang dijalankan sebelum/setelah request/response masuk ke server melalui route

## Penggunaan
dapat digunakan di:
- root level/setelah router
- group level/kumpulan route di grup tertentu
- route level

## Fungsi
- Middleware untuk authentikasi
- Middleware untuk logging
- redirect

## Middleware di go
di framework echo terdapat beberapa middlewatre yang dapat digunakan
- auth middleware (JWT, basic auth, auth)
- logging (logger)
- dsb