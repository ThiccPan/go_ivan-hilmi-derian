# intro to rest API
## API Definition
api adalah fungsi/metode yang digunakan untuk memanggil/mengakses fitur dari OS, aplikasi, maupun komponen sistem lainnya  
terdiri dari 2 bagian:
- server yang menyediakan layanan
- client yang memanggil/mengkonsumsi layanan
## REST API
- salah satu metode komunikasi yang sering diguanakan oleh webservice
- menggunakan protokol http
- terdiri dari beberapa req method:
    - GET
    - POST
    - PUT
    - DELETE
    - HEAD
    - OPTION
    - PATCH
- menggunakan format req dan res seperti JSON, XML, dan SOAP
- Memiliki response code untuk mengecek status transaksi:
    - 200: OK
    - 201: Created
    - 400: Bad req
    - 401: Unauthorized
    - 404: Not found
    - 405: Method not allowed
    - 500: server error
- dapat menggunakan api tools (eg: postman) sebagai alat testing endpoint api
- Best practice:
    - menggunakan kata benda/noun
    - menggunakan kata jamak/plural
    - menggunakan resource nesting sebagai representasi hierarki
    - menggunakan query parameter untuk filtering
