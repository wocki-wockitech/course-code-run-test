CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INT NOT NULL,
    email TEXT
);

INSERT INTO users (name, age, email) VALUES
    ('Alice', 25, 'alice@example.com'),
    ('Bob', 17, 'bob@example.com'),
    ('Charlie', 30, 'charlie@example.com'),
    ('Diana', 16, 'diana@example.com'),
    ('Eve', 22, 'eve@example.com'),
    ('Frank', 18, 'frank@example.com'),
    ('Grace', 19, 'grace@example.com');
