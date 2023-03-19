CREATE TABLE users (
    id bigserial PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    birthdate DATE NOT NULL,
    status SMALLINT NOT NULL,
    gender char(1) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- product

CREATE TABLE product_types (
    id bigserial PRIMARY KEY,
    type varchar(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE product_description (
    id bigserial PRIMARY KEY,
    description text NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE payment_method (
    id bigserial PRIMARY KEY,
    name varchar(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE operator (
    id bigserial PRIMARY KEY,
    name varchar(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE product_descriptions (
    id bigserial PRIMARY KEY,
    description text NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE products (
    id bigserial PRIMARY KEY,
    type_id bigint REFERENCES product_types (id),
    operator_id bigint REFERENCES operator (id),
    name varchar(50) NOT NULL,
    price numeric(25, 2) NOT NULL,
    stock int NOT NULL
);

ALTER TABLE product_descriptions
        ADD FOREIGN KEY (id) REFERENCES products (id)
                DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE products
        ADD FOREIGN KEY (id) REFERENCES product_descriptions (id)
                DEFERRABLE INITIALLY DEFERRED;

CREATE TABLE transaction (
    id bigserial PRIMARY KEY,
    user_id bigint REFERENCES user (id),
    payment_method_id bigint REFERENCES operator (id),
    quantity_total int NOT NULL,
    price_total numeric(25, 2) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE transaction_detail (
    id bigserial PRIMARY KEY,
    transaction_id bigint REFERENCES transaction (id),
    product_id bigint REFERENCES products (id),
    quantity int NOT NULL,
    price_item numeric(25, 2) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);


Create tabel kurir dengan field id, name, created_at, updated_at.

CREATE TABLE kurir (
    id bigserial PRIMARY KEY,
    name varchar(50) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

-- 4. Tambahkan ongkos_dasar column di tabel kurir.
ALTER TABLE kurir
    ADD COLUMN ongkos_dasar numeric(25, 2) NOT NULL;

-- 5. Rename tabel kurir menjadi shipping.
ALTER TABLE kurir
    RENAME TO shipping;    

-- 6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
DROP TABLE shipping;

-- 7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
    -- 1. 1-to-1: payment method description.
        CREATE TABLE operator_description (
            id bigserial PRIMARY KEY,
            description text NOT NULL,
            created_at timestamp NOT NULL,
            updated_at timestamp NOT NULL
        );

        ALTER TABLE operator_description
            ADD FOREIGN KEY (id) REFERENCES operator (id)
                    DEFERRABLE INITIALLY DEFERRED;

        ALTER TABLE operator
                ADD FOREIGN KEY (id) REFERENCES operator_description (id)
                    DEFERRABLE INITIALLY DEFERRED;

    -- 2. 1-to-many: user dengan alamat.

        CREATE TABLE address_detail (
            post_code VARCHAR(5) PRIMARY KEY
            -- bisa ditambah field lain seperti kota, etc
        );

        ALTER TABLE users
            ALTER COLUMN address SET DATA TYPE VARCHAR(5);

        ALTER TABLE users
            ADD FOREIGN KEY (address) REFERENCES address_detail (post_code);
    -- 3. many-to-many: user dengan payment method menjadi user_payment_method_detail.
        CREATE TABLE payment_detail (
            id bigserial PRIMARY KEY,
            transaction_id BIGINT NOT NULL,
            payment_id BIGINT NOT NULL,
            detail VARCHAR(50) NOT NULL,
            created_at timestamp NOT NULL,
            updated_at timestamp NOT NULL
        );

        ALTER TABLE payment_detail
            ADD FOREIGN KEY (transaction_id) REFERENCES transaction (id);

        ALTER TABLE payment_detail
            ADD FOREIGN KEY (payment_id) REFERENCES payment_method (id);

