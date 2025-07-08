-- Bank
CREATE TABLE IF NOT EXISTS bank (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bank_id INT NOT NULL UNIQUE,
    slug VARCHAR(50) NOT NULL UNIQUE,
    swift VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    acct_length INT NOT NULL,
    country_id INT NOT NULL,
    is_mobilemoney INT NOT NULL,
    is_active INT NOT NULL,
    is_rtgs INT NOT NULL,
    active INT NOT NULL,
    is_24hrs INT NOT NULL,
    currency VARCHAR(25) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

-- Transaction
CREATE TABLE IF NOT EXISTS transaction (
    status VARCHAR(25) NOT NULL,
    ref_id VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    currency VARCHAR(25) NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    charge DOUBLE PRECISION NOT NULL,
    trans_id VARCHAR(255),
    payment_method VARCHAR(50) NOT NULL,
    customer_id VARCHAR(255) NOT NULL,
    customer_email VARCHAR(255),
    customer_first_name VARCHAR(50),
    customer_last_name VARCHAR(50),
    customer_mobile VARCHAR(25)
);

-- Transfer
CREATE TABLE IF NOT EXISTS transfer (
    account_name VARCHAR(50) NOT NULL,
    account_number VARCHAR(255) NOT NULL,
    currency VARCHAR(25) NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    charge DOUBLE PRECISION NOT NULL,
    transfer_type VARCHAR(50) NOT NULL,
    chapa_reference VARCHAR(255) NOT NULL,
    bank_code INT NOT NULL,
    bank_name VARCHAR(50) NOT NULL,
    bank_reference VARCHAR(255),
    status VARCHAR(50) NOT NULL,
    reference VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
