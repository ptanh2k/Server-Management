CREATE TABLE IF NOT EXISTS servers (
    server_id INTEGER PRIMARY KEY,
    server_name VARCHAR(20) NOT NULL,
    status BOOLEAN NOT NULL,
	created_time TIMESTAMP,
	last_updated TIMESTAMP,
	domain TEXT UNIQUE NOT NULL,
	created_by INTEGER REFERENCES users(id)
);

-- DROP TABLE users

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
);

SELECT * FROM users WHERE users.username = 'admin'
