-- 2. Select
-- a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT id, gender FROM users WHERE gender = 'M';

-- b. Tampilkan product dengan id = 3.
SELECT * FROM products WHERE id = 3;

-- c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM users 
WHERE created_at > created_at - interval '1 days'
AND gender LIKE '%M%'; -- tidak ada field nama, diganti dengan gender yang mengandung char M

-- d. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) 
FROM users 
WHERE gender = 'F'; 

-- e. Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT *
FROM users
ORDER BY updated_at; -- tidak ada field nama, diganti dengan urutan updated_at

-- f. Tampilkan 5 data pada data product
SELECT *
FROM products
LIMIT 5;