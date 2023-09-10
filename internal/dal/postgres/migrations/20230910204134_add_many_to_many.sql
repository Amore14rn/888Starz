-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.user_orders_products (
                                             user_id VARCHAR(255),
                                             order_id VARCHAR(255),
                                             product_id VARCHAR(255),
                                             PRIMARY KEY (user_id, order_id, product_id),
                                             FOREIGN KEY (user_id) REFERENCES users(id),
                                             FOREIGN KEY (order_id) REFERENCES orders(id),
                                             FOREIGN KEY (product_id) REFERENCES products(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.user_orders_products;
-- +goose StatementEnd
