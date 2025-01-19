CREATE TABLE animal_estado (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    persona_id INT,
    animal_id INT NOT NULL,
    estado_id INT NOT NULL,
    final DATE,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_persona FOREIGN KEY (persona_id) REFERENCES persona (id) ON DELETE CASCADE,
    CONSTRAINT fk_animal FOREIGN KEY (animal_id) REFERENCES animal (id) ON DELETE CASCADE,
    CONSTRAINT fk_estado FOREIGN KEY (estado_id) REFERENCES estado (id) ON DELETE CASCADE
);
