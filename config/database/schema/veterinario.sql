CREATE TABLE veterinario (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    dni VARCHAR(255) NOT NULL,
    nombre VARCHAR(255) NOT NULL,
    apellido VARCHAR(255) NOT NULL,
    telefono VARCHAR(255) NOT NULL,
    clinica_id INT NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_clinica FOREIGN KEY (clinica_id) REFERENCES clinica (id)
);
