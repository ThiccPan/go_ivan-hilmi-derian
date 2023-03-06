# Problem Solving Paradigm

## Brute Force
Mencoba setiap kombinasi yang ada, biasanya mudah untuk diimplementasikan

## Divide & Conquer
membagi masalah menjadi bagian yang lebih kecil. terkadang menggabungkan solusi yang dihasilkan (combine)
contoh : 
- binary search (perbandingan time complexity dengan linear O(log N) dengan O(N))
- pangkat(power)

## Greedy
mencari solusi lokal (bukan global seperti brute force) yang paling optimal. solusi yang ditawarkan situasional karena tidak selamanya merupakan solusi yang paling tepat (mungkin ada yang lebih tepat/efektif)
contoh : 
- coin change problem
- dijkstra maze solving algo

## Dynamic Programming
memecahkan masalah menjadi sub masalah yang lebih kecil. ciri dari masalah DP adalah banyak sub masalah yang diulang, sehingga bisa menggunakan cache untuk mempercepat penyelesaian masalah (dengan mengorbankan space)  
tipe :
- top down (menggunakan memoization, biasanya menggunakan recursive func)
- bottom up (menggunakan tabulation)