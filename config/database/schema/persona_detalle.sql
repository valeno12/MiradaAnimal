CREATE TABLE persona_detalle (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    persona_id INT NOT NULL,
    pregunta_id INT NOT NULL,
    detalle VARCHAR(255),
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_persona FOREIGN KEY (persona_id) REFERENCES persona (id) ON DELETE CASCADE,
    CONSTRAINT fk_pregunta FOREIGN KEY (pregunta_id) REFERENCES pregunta (id) ON DELETE CASCADE
);
