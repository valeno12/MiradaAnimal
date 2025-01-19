CREATE TABLE atributo (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    descripcion VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);
