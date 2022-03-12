CREATE TABLE IF NOT EXISTS products (
    id SERIAL,
    name TEXT,
    description TEXT,
    price DOUBLE PRECISION,
    quantity INT
);

ALTER TABLE products ADD CONSTRAINT product_pk_id PRIMARY KEY (id);


INSERT INTO products (name, price, description, quantity) VALUES ('Apple', 1.99, 'A red fruit', 100);
INSERT INTO products (name, price, description, quantity) VALUES ('Notebook', 5400.30, 'A very good notebook', 5);
INSERT INTO products (name, price, description, quantity) VALUES ('Shirt', 120, 'A very good Shirt', 1);
INSERT INTO products (name, price, description, quantity) VALUES ('Headset', 500, 'A very good Headset', 3);