CREATE TYPE status_type AS ENUM (
    'Wait for payment status',
    'Pending for expert',
    'Pending for matin',
    'In progress',
    'Done'
);

CREATE TABLE IF NOT EXISTS expert_ads (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    expert_id INT,
    ads_id INT NOT NULL,
    report TEXT,
    status STATUS_TYPE,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (expert_id) REFERENCES users(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (ads_id) REFERENCES ads(id)
);