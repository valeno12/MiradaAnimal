CREATE TABLE formulario (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    animal_id INT NOT NULL,
    persona_id INT NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_animal FOREIGN KEY (animal_id) REFERENCES animal (id) ON DELETE CASCADE,
    CONSTRAINT fk_persona FOREIGN KEY (persona_id) REFERENCES persona (id) ON DELETE CASCADE
);
