CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    status SMALLINT NOT NULL,
    dob DATE NOT NULL,
    gender char(1) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

ALTER TABLE users
    

CREATE TABLE payment_methods (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE product_types (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE operators (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    product_type_id BIGINT REFERENCES product_types (id) NOT NULL,
    operator_id BIGINT REFERENCES operators (id) NOT NULL,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE product_descriptions (
    id BIGSERIAL PRIMARY KEY,
    description text NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

ALTER TABLE product_descriptions
        ADD FOREIGN KEY (id) REFERENCES products (id)
                DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE products
        ADD FOREIGN KEY (id) REFERENCES product_descriptions (id)
                DEFERRABLE INITIALLY DEFERRED;

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users (id) NOT NULL,
    payment_method_id BIGINT REFERENCES operators (id) NOT NULL,
    status VARCHAR(10) NOT NULL,
    total_qty INT NOT NULL,
    total_price NUMERIC(25, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE transaction_details (
    id BIGSERIAL PRIMARY KEY,
    transaction_id BIGINT REFERENCES transactions (id) NOT NULL,
    product_id BIGINT REFERENCES products (id) NOT NULL,
    status VARCHAR(10) NOT NULL,
    qty INT NOT NULL,
    price NUMERIC(25, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);