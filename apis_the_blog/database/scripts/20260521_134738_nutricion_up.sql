-- UP: Crear tabla nutricion
CREATE TABLE IF NOT EXISTS blog.nutricion (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT,
    objetivo TEXT,
    imagen_url VARCHAR(255),
    autor_id INTEGER,
    publicado BOOLEAN DEFAULT false,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);