CREATE TABLE animal (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    nombre VARCHAR(255) NOT NULL,
    especie VARCHAR(255) NOT NULL,
    foto VARCHAR(255),
    sexo VARCHAR(255) NOT NULL,
    edad_meses INT NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);
