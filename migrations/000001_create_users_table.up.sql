CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    user_name TEXT NOT NULL,
    user_password TEXT NOT NULL,
    balance INT DEFAULT 0,
    status SMALLINT DEFAULT 0
);
