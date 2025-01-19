CREATE TABLE persona (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    dni VARCHAR(255) UNIQUE NOT NULL,
    nombre VARCHAR(255) NOT NULL,
    apellido VARCHAR(255) NOT NULL,
    direccion VARCHAR(255) NOT NULL,
    telefono VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    estado VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);
