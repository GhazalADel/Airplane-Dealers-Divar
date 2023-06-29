CREATE TABLE IF NOT EXISTS expert_ads (
    id SERIAL PRIMARY KEY,
    expert_id INT,
    ads_id INT NOT NULL,
    report VARCHAR,
    status boolean,
    FOREIGN KEY (expert_id) REFERENCES users(id),
    FOREIGN KEY (ads_id) REFERENCES ads(id)
);