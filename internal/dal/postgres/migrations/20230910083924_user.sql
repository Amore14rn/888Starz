-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- Создание таблиц
CREATE TABLE public.products (
                                 id VARCHAR(255) PRIMARY KEY,
                                 description VARCHAR(255),
                                 quantity INTEGER,
                                 created_at TIMESTAMP,
                                 updated_at TIMESTAMP
);

CREATE TABLE public.product_history (
                                        product_id VARCHAR(255) REFERENCES products(id),
                                        price DOUBLE PRECISION,
                                        timestamp TIMESTAMP,
                                        PRIMARY KEY (product_id, timestamp)
);

CREATE TABLE public.users (
                              id VARCHAR(255) PRIMARY KEY,
                              first_name VARCHAR(255),
                              last_name VARCHAR(255),
                              full_name VARCHAR(255),
                              age INTEGER,
                              is_married BOOLEAN,
                              password VARCHAR(255),
                              order_id VARCHAR(255),
                              created_at TIMESTAMP,
                              updated_at TIMESTAMP
);

CREATE TABLE public.orders (
                               id VARCHAR(255) PRIMARY KEY,
                               user_id VARCHAR(255),
                               timestamp TIMESTAMP
);

CREATE TABLE public.order_products (
                                       order_id VARCHAR(255),
                                       product_id VARCHAR(255),
                                       quantity INTEGER,
                                       price DOUBLE PRECISION,
                                       PRIMARY KEY (order_id, product_id),
                                       FOREIGN KEY (order_id) REFERENCES orders(id),
                                       FOREIGN KEY (product_id) REFERENCES products(id)
);

-- Добавление внешних ключей
ALTER TABLE public.users ADD CONSTRAINT fk_orders FOREIGN KEY (order_id) REFERENCES orders(id);
ALTER TABLE public.orders ADD CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE public.order_products ADD CONSTRAINT fk_orders FOREIGN KEY (order_id) REFERENCES orders(id);
ALTER TABLE public.order_products ADD CONSTRAINT fk_products FOREIGN KEY (product_id) REFERENCES products(id);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- Удаление внешних ключей
ALTER TABLE public.order_products DROP CONSTRAINT fk_products;
ALTER TABLE public.order_products DROP CONSTRAINT fk_orders;
ALTER TABLE public.orders DROP CONSTRAINT fk_users;
ALTER TABLE public.users DROP CONSTRAINT fk_orders;

-- Удаление таблиц
DROP TABLE public.order_products;
DROP TABLE public.orders;
DROP TABLE public.users;
DROP TABLE public.product_history;
DROP TABLE public.products;
