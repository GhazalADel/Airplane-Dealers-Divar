CREATE TABLE IF NOT EXISTS repair_request (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    ads_id INT NOT NULL,
    status EXPERT_STATUS_TYPE,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (ads_id) REFERENCES ads(id)
);