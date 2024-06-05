CREATE TABLE product_stocks (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  stock_id UUID REFERENCES stocks(id) NOT NULL,
  product_id UUID REFERENCES products(id) NOT NULL,
  UNIQUE (stock_id, product_id),  -- Prevents duplicate entries
  deleted_at TIMESTAMP WITHOUT TIME ZONE,
  created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_date TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);