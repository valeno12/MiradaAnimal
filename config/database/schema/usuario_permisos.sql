CREATE TABLE usuario_permisos (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL,
    permiso_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_usuario FOREIGN KEY (usuario_id) REFERENCES usuarios (id) ON DELETE CASCADE,
    CONSTRAINT fk_permiso FOREIGN KEY (permiso_id) REFERENCES permisos (id) ON DELETE CASCADE,
    CONSTRAINT uc_usuario_permiso UNIQUE (usuario_id, permiso_id)
);
