CREATE TABLE animal_atributos (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    animal_id INT NOT NULL,
    atributo_id INT NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_animal FOREIGN KEY (animal_id) REFERENCES animal (id) ON DELETE CASCADE,
    CONSTRAINT fk_atributo FOREIGN KEY (atributo_id) REFERENCES atributo (id) ON DELETE CASCADE
);
