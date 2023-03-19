-- ### Join, Union, Sub query, Function

-- 1. Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT * FROM transactions WHERE user_id = 1
UNION 
SELECT * FROM transactions WHERE user_id = 2;

-- 2. Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(total_price)
FROM transactions
WHERE user_id = 1;

-- 3. Tampilkan total transaksi dengan product type 2.
SELECT COUNT(transaction_details.id) 
FROM transaction_details 
INNER JOIN products 
ON products.id = transaction_details.product_id 
WHERE products.product_type_id = 2;

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT product_types.name, products.* 
FROM products RIGHT JOIN product_types
ON product_types.id = products.product_type_id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT t.id, p.name, td.* 
FROM transaction_details AS td 
INNER JOIN products AS p
ON td.product_id = p.id
INNER JOIN transactions AS t
ON td.transaction_id = t.id;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
CREATE OR REPLACE FUNCTION process_delete_all() RETURNS TRIGGER AS $delete_all$
    BEGIN
        DELETE FROM transaction_details
            WHERE transaction_id = OLD.id;
        RETURN NULL;
    END;
$delete_all$ LANGUAGE plpgsql;

CREATE TRIGGER delete_all
BEFORE DELETE ON transactions
    FOR EACH ROW EXECUTE FUNCTION process_delete_all();


-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.

-- sama dengan function di file insert.sql
CREATE OR REPLACE FUNCTION process_transaction() RETURNS TRIGGER AS $transactions$
    BEGIN
        IF (TG_OP = 'DELETE') THEN
            UPDATE transactions SET 
                total_qty = total_qty - OLD.qty,
                total_price = total_price - OLD.price,
                updated_at = current_timestamp
            WHERE transactions.id = OLD.transaction_id;

        ELSIF (TG_OP = 'UPDATE') THEN
            UPDATE transactions SET 
                total_qty = total_qty - OLD.qty + NEW.qty,
                total_price = total_price - OLD.price + NEW.price,
                updated_at = current_timestamp
            WHERE transactions.id = OLD.transaction_id;
            
        ELSEIF (TG_OP = 'INSERT') THEN
            UPDATE transactions SET 
                total_qty = total_qty + NEW.qty,
                total_price = total_price + NEW.price,
                updated_at = current_timestamp
            WHERE transactions.id = NEW.transaction_id;
        END IF;
        RETURN NULL;
    END;
$transactions$ LANGUAGE plpgsql;

CREATE TRIGGER transactions
AFTER INSERT OR UPDATE OR DELETE ON transaction_details
    FOR EACH ROW EXECUTE FUNCTION process_transaction();


-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
SELECT * 
FROM products
WHERE id NOT IN (
    SELECT product_id
    FROM transaction_details
    WHERE product_id IS NOT NULL
);