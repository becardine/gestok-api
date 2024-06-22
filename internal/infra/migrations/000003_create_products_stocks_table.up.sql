CREATE TABLE product_stocks (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    product_id CHAR(36) REFERENCES products(id),
    stock_id CHAR(36) REFERENCES stocks(id),
    quantity INTEGER NOT NULL,
    deleted_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);