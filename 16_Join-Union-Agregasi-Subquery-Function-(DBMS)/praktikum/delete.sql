-- 4. Delete
-- 1. Delete data pada tabel product dengan id = 1.
BEGIN;
DELETE FROM transaction_details WHERE product_id = 1; -- hapus karena constraint foreign key
DELETE FROM products WHERE id = 1;
DELETE FROM product_descriptions WHERE id = 1;
COMMIT;

-- 2. Delete pada pada tabel product dengan product type id 1.
BEGIN;
DELETE FROM product_descriptions USING products WHERE products.product_type_id = 1 AND product_descriptions.id = products.id;
DELETE FROM products WHERE product_type_id = 1;
COMMIT;