CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);

INSERT INTO users (username, password)
VALUES 
('user', encode(digest('user123', 'sha256'), 'hex')),
('admin', encode(digest('admin123', 'sha256'), 'hex'));
