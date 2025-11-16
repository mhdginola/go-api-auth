CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  role_id numeric NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

INSERT INTO roles (id, name, description) VALUES
(1, 'admin', 'Full access'),
(2, 'user', 'Regular user')
ON CONFLICT (name) DO NOTHING;

INSERT INTO users (username, password, role_id)
VALUES 
('user', encode(digest('user123', 'sha256'), 'hex'), 2),
('admin', encode(digest('admin123', 'sha256'), 'hex'), 1)
ON CONFLICT (username) DO NOTHING;
