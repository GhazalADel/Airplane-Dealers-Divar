CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL UNIQUE,
    role VARCHAR(255) NOT NULL,
    token VARCHAR(1020),
    is_active boolean,
);