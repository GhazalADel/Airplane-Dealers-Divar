CREATE TYPE expert_status_type AS ENUM (
    'Wait for payment status',
    'Pending for expert',
    'In progress',
    'Confirmed'
);

CREATE TABLE IF NOT EXISTS expert_ads (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    expert_id INT,
    ads_id INT NOT NULL,
    report TEXT,
    status EXPERT_STATUS_TYPE,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (expert_id) REFERENCES users(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (ads_id) REFERENCES ads(id)
);