CREATE TABLE animal_vacunas (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    animal_id INT NOT NULL,
    vacuna_id INT NOT NULL,
    fecha DATE NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_animal FOREIGN KEY (animal_id) REFERENCES animal (id) ON DELETE CASCADE,
    CONSTRAINT fk_vacuna FOREIGN KEY (vacuna_id) REFERENCES vacuna (id) ON DELETE CASCADE
);
