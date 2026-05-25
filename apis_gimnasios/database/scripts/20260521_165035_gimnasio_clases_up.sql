CREATE TABLE IF NOT EXISTS gimnasios.gimnasio_clases (
    id SERIAL PRIMARY KEY,
    gimnasio_id INTEGER REFERENCES gimnasios.Sedes_gimnasios(id) ON DELETE CASCADE,
    nombre_clase VARCHAR(100) NOT NULL,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);