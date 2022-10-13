CREATE TABLE IF NOT EXISTS servers (
    server_id SERIAL PRIMARY KEY,
    server_name VARCHAR(20) UNIQUE NOT NULL,
    status BOOLEAN NOT NULL,
	created_time TIMESTAMP NOT NULL,
	last_updated TIMESTAMP NOT NULL,
	ipv4 VARCHAR(50) UNIQUE NOT NULL,
	user_id INTEGER REFERENCES users(id)
);

INSERT INTO servers (server_name, status, created_time, last_updated, ipv4, user_id) VALUES
('Cycir', 'ON', LOCALTIMESTAMP, LOCALTIMESTAMP, '192.168.1.1', 1)

SELECT * FROM servers

-- DROP TABLE users

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
);

SELECT * FROM users WHERE users.username = 'admin'
