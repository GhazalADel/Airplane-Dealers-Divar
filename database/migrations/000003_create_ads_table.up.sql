CREATE TABLE IF NOT EXISTS ads (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    image VARCHAR(255),
    description TEXT,
    subject VARCHAR(255) NOT NULL,
    price BIGINT NOT NULL,
    category_id INT NOT NULL,
    status VARCHAR(255),
    fly_time TIMESTAMP,
    airplane_model VARCHAR(255),
    repair_check boolean,
    expert_check boolean,
    plane_age INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
