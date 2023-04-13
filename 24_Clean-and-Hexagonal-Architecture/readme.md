# 24 Clean and Hexagonal Architecture
## Definisi
Tata cara penyusunan struktur codebase untuk membuat codebase yang modular, scalable, dan mudah untuk dimaintain

## Constraint
- tidak bergantung pada framework tertentu
- business rule mudah untuk ditest
- tidak bergantung pada ui
- tidak bergantung pada database
- tidak bergantung pada external library

## Keuntungan
- navigasi codebase yang lebih mudah
- kecepatan development, terutama di project yang besar
- mocking dependencies lebih mudah
- memudahkan untuk realisasi dari prototype

## Layer
classic layer adalah sebagai berikut:
- entities layer
    - menyimpan business object
- usecase - domain layer
    - letak business logic
- controller - presentasion layer
    - presentsasi data & interaksi user
    - dapat berubah tergantung cara komunikasi (http, grpc, etc..)
- drivers - data layer
    - manage data aplikasi
