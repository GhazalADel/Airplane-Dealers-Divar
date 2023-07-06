CREATE TABLE IF NOT EXISTS expert_ads (
    id SERIAL PRIMARY KEY,
    report VARCHAR(255),
    status INT,
    expert_id INT,
    ads_id INT NOT NULL,
    FOREIGN KEY (expert_id) REFERENCES users (id),
    FOREIGN KEY (ads_id) REFERENCES ads (id)
);