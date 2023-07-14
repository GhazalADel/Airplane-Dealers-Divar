CREATE TABLE IF NOT EXISTS activity_log (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    subject_type VARCHAR(50) NOT NULL,
    subject_id BIGINT NOT NULL,
    causer_type VARCHAR(50),
    causer_id BIGINT,
    log_id BIGINT,
    description VARCHAR(255),
    FOREIGN KEY (log_id) REFERENCES log_name(id)
);