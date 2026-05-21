CREATE TABLE IF NOT EXISTS gimnasios.resenas_gimnasio (
    id SERIAL PRIMARY KEY,
    gimnasio_id INTEGER REFERENCES gimnasios.Sedes_gimnasios(id) ON DELETE CASCADE,
    usuario_id INTEGER,
    calificacion INTEGER CHECK (calificacion >= 1 AND calificacion <= 5),
    comentario TEXT,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);