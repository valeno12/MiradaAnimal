CREATE TABLE animal_turnos (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    descripcion VARCHAR(255) NOT NULL,
    fecha DATE,
    animal_id INT NOT NULL,
    veterinario_id INT NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_animal FOREIGN KEY (animal_id) REFERENCES animal (id) ON DELETE CASCADE,
    CONSTRAINT fk_veterinario FOREIGN KEY (veterinario_id) REFERENCES veterinario (id) ON DELETE CASCADE
);
