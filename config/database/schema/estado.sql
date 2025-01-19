CREATE TABLE estado (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    nombre VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);
