CREATE TABLE IF NOT EXISTS servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    ip TEXT NOT NULL,
    port SMALLINT NOT NULL,
    status BOOLEAN NOT NULL
);


INSERT INTO servers (name, ip, port, status) VALUES
('server1', '192.168.1.0', 8080, 'on'),
('server2', '192.168.1.1', 3000, 'off')

SELECT * FROM servers

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIME,
    updated_at TIME,
    deleted_at TIME
);

SELECT * FROM users WHERE users.username = 'admin' AND users.deleted_at IS NULL LIMIT 1
