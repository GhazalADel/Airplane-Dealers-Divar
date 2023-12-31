CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    transaction_type VARCHAR(255) NOT NULL,
    object_id BIGINT NOT NULL,
    amount BIGINT,
    status VARCHAR(255),
    authority VARCHAR(255),
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);