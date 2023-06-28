CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    amount BIGINT,
    status VARCHAR(255),
    authority VARCHAR(255),
    created_at TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);