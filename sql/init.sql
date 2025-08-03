CREATE TABLE orders (
    id SERIAL PRIMARY KEY, -- Auto-incrementing database ID

    order_id VARCHAR(64) NOT NULL UNIQUE, -- Custom string-based order ID, indexed & unique
    phone_number VARCHAR(15) NOT NULL,

    adult_count INTEGER NOT NULL DEFAULT 0,
    child_count INTEGER NOT NULL DEFAULT 0,
    cycle_count INTEGER NOT NULL DEFAULT 0,

    total_amount DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    additional_charges DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    grand_total_amount DECIMAL(10, 2) NOT NULL DEFAULT 0.00,

    status VARCHAR(20) NOT NULL DEFAULT 'NEW', -- NEW | PENDING | SUCCESS | FAILED

    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX idx_order_order_id ON orders(order_id);


CREATE TABLE rates (
    id SERIAL PRIMARY KEY,
    item_type VARCHAR(20),
    price_per_unit DECIMAL(10, 2) NOT NULL DEFAULT 0.00
);