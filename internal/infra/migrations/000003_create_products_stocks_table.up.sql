CREATE TABLE product_stock (
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES products(id),
    stock_id UUID REFERENCES stocks(id),
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);