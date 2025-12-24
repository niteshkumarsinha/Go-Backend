CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
); 
INSERT INTO users (name) VALUES ('John Doe');
INSERT INTO users (name) VALUES ('Jane Doe');
INSERT INTO users (name) SELECT 'User ' || i FROM generate_series(1, 100) AS i;
