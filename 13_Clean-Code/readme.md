# Clean Code
Guide/panduan dalam menuliskan kode
## Tujuan
- menghasilkan kode yang mudah dibaca
- mempermudah kolaborasi
- mudah untuk diupdate/diextend kedepannya
## Aturan
- Penamaan yang mudah dipahami
- Penamaan yang mudah dieja dan dicari
- Penamaan yang singkat namun dapat mendeskripsikan konteks dengan jelas
- konsistensi dalam pola penamaan kode (variabel, konstanta, etc)
- menghindari penambahan konteks yang tidak perlu
- penulisan komentar yang efektif
    - tidak pada setiap line
    - terdapat di atas bagian yang perlu dijelaskan
    - menghindari pemberian komentar pada kode yang tidak digunakan
- penulisan function yang efektif
    - menghindari side effect (perubahan variable diluar scope pemanggilan func)
    - menghindari jumlah parameter yang terlalu banyak
- Menggunakan konveksi yang sudah ada
## KISS dan DRY
### KISS
buat kode sesimpel mungkin
- fungsi dan class harus kecil
- fungsi digunakan untuk satu hal saja
- fungsi jangan terlalu banyak argumen
- perhatikan agar mencapai kondisi yang seimbang

### DRY
mengurangi duplikasi pada kode. kurangi melakukan copy paste pada kode

### Refactoring
menulis ulang kode tanpa mengubah logic dibaliknya. kode ditulis ulang menerapkan prinsip KISS dan DRY