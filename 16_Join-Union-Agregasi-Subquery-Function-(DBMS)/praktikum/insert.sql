-- DML
-- Insert
-- a. Insert 5 operators pada table operators.
INSERT INTO operators (name, created_at, updated_at) VALUES
    ('operator1', current_timestamp, current_timestamp),
    ('operator2', current_timestamp, current_timestamp),
    ('operator3', current_timestamp, current_timestamp),
    ('operator4', current_timestamp, current_timestamp),
    ('operator5', current_timestamp, current_timestamp);

-- b. Insert 3 product type.
INSERT INTO product_types (name, created_at, updated_at) VALUES
    ('type1', current_timestamp, current_timestamp),
    ('type2', current_timestamp, current_timestamp),
    ('type3', current_timestamp, current_timestamp),
    ('type4', current_timestamp, current_timestamp),
    ('type5', current_timestamp, current_timestamp);

-- c. Insert 2 product dengan product type id = 1, dan operators id = 3.
BEGIN;
INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at) VALUES
    (1, 1, 3, '23032', 'product1', 1, current_timestamp, current_timestamp),
    (2, 1, 3, '23033', 'product2', 1, current_timestamp, current_timestamp);
INSERT INTO product_descriptions (id, description, created_at, updated_at) VALUES
    (1, 'description prod 1', current_timestamp, current_timestamp),
    (2, 'description prod 2', current_timestamp, current_timestamp);
COMMIT;

-- d. Insert 3 product dengan product type id = 2, dan operators id = 1.
BEGIN;
INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at) VALUES
    (3, 2, 1, '23033', 'product3', 1, current_timestamp, current_timestamp),
    (4, 2, 1, '23034', 'product4', 1, current_timestamp, current_timestamp),
    (5, 2, 1, '23035', 'product5', 1, current_timestamp, current_timestamp);
INSERT INTO product_descriptions (id, description, created_at, updated_at) VALUES
    (3, 'description prod 3', current_timestamp, current_timestamp),
    (4, 'description prod 4', current_timestamp, current_timestamp),
    (5, 'description prod 5', current_timestamp, current_timestamp);
COMMIT;

-- e. Insert 3 product dengan product type id = 3, dan operators id = 4.
BEGIN;
INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at) VALUES
    (6, 3, 4, '23036', 'product6', 1, current_timestamp, current_timestamp),
    (7, 3, 4, '23037', 'product7', 1, current_timestamp, current_timestamp),
    (8, 3, 4, '23038', 'product8', 1, current_timestamp, current_timestamp);
INSERT INTO product_descriptions (id, description, created_at, updated_at) VALUES
    (6, 'description prod 6', current_timestamp, current_timestamp),
    (7, 'description prod 7', current_timestamp, current_timestamp),
    (8, 'description prod 8', current_timestamp, current_timestamp);
COMMIT;

-- f. Insert product description pada setiap product.
-- sudah ditambahkan bersamaan dengan tugas sebelumnya
SELECT products.id, products.name, product_descriptions.description
FROM products, product_descriptions
WHERE products.id = product_descriptions.id;

-- g. Insert 3 payment methods.
INSERT INTO payment_methods (name, status, created_at, updated_at) VALUES
    ('method1', '1', current_timestamp, current_timestamp),
    ('method2', '1', current_timestamp, current_timestamp),
    ('method3', '1', current_timestamp, current_timestamp);

-- h. Insert 5 user pada tabel user.
INSERT INTO users (status, dob, gender, created_at, updated_at) VALUES
    ('1', current_date, 'M', current_timestamp, current_timestamp),
    ('2', current_date, 'M', current_timestamp, current_timestamp),
    ('3', current_date, 'M', current_timestamp, current_timestamp),
    ('4', current_date, 'M', current_timestamp, current_timestamp),
    ('5', current_date, 'M', current_timestamp, current_timestamp);

-- i. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
-- j. Insert 3 product di masing-masing transaksi.

-- trigger dan function untuk mengupdate total_qty dan total_price secara otomatis
-- ketika terdapat operasi DML di transaction_details
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



INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price, created_at, updated_at) VALUES
    (1, 1, '1', 0, 0, current_timestamp, current_timestamp),
    (2, 1, '1', 0, 0, current_timestamp, current_timestamp),
    (3, 1, '1', 0, 0, current_timestamp, current_timestamp),
    (4, 1, '1', 0, 0, current_timestamp, current_timestamp),
    (5, 1, '1', 0, 0, current_timestamp, current_timestamp);

INSERT INTO transaction_details (transaction_id, product_id, status, qty, price, created_at, updated_at) VALUES
    -- transaksi user 1
    (6, 3, '1', 1, 100, current_timestamp, current_timestamp),
    (6, 3, '1', 1, 100, current_timestamp, current_timestamp),
    (6, 3, '1', 1, 100, current_timestamp, current_timestamp),

    -- transaksi user 2
    (7, 3, '1', 1, 100, current_timestamp, current_timestamp),
    (7, 3, '1', 1, 100, current_timestamp, current_timestamp),
    (7, 3, '1', 1, 100, current_timestamp, current_timestamp),
    
    -- transaksi user 3
    (8, 3, '1', 1, 100, current_timestamp, current_timestamp),
    (8, 3, '1', 1, 100, current_timestamp, current_timestamp),
    (8, 3, '1', 1, 100, current_timestamp, current_timestamp),
    
    -- transaksi user 4
    (9, 3, '1', 1, 100, current_timestamp, current_timestamp),
    (9, 3, '1', 1, 100, current_timestamp, current_timestamp),
    (9, 3, '1', 1, 100, current_timestamp, current_timestamp),
    
    -- transaksi user 5
    (10, 3, '1', 1, 100, current_timestamp, current_timestamp),
    (10, 3, '1', 1, 100, current_timestamp, current_timestamp),
    (10, 3, '1', 1, 100, current_timestamp, current_timestamp);