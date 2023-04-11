# Unit Test

## Definisi
testing yang dilakukan pada level paling kecil pada aplikasi (function). 

## Kegunaan
- Meningkatkan kualitas kode
- mempermudah refactoring
- membantu dokumentasi kode
- meningkatkan kualitas kode

## Struktur
- diletakkan secara tersentralisasi di folder test
- bersamaan dengan file produksi

## Mocking
- melakukan mock pada 3rd party dependencies, karena tidak masuk ke dalam unit testing
- dengan melakukan object mock

## Coverage
- menganalisa aplikasi dan menyajikan hasil analisa menggunakan format data yang tertandarisasi
```bash
go test location/to/file
-coverprofile=cover.out && go tool cover    
-html=cover.out
```

## unit testing di go
- beri nama file dengan akhiran _test.go
- letakkan di lokasi yang sama dengan file yang di test, atau di folder yang berbeda dengan package yang sama