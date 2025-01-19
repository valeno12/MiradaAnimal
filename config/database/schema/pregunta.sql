CREATE TABLE pregunta (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    descripcion VARCHAR(255) NOT NULL,
    requiere_detalle BOOLEAN NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);
