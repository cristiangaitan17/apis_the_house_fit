-- UP: Crear tabla comentarios_comunidad
CREATE TABLE IF NOT EXISTS blog.comentarios_comunidad (
    id SERIAL PRIMARY KEY,
    categoria_id INTEGER REFERENCES blog.categorias(id) ON DELETE SET NULL,
    usuario_id INTEGER,
    contenido TEXT NOT NULL,
    calificacion INTEGER CHECK (calificacion >= 1 AND calificacion <= 5),
    likes INTEGER DEFAULT 0,
    dislikes INTEGER DEFAULT 0,
    estado VARCHAR(20) DEFAULT 'activo',
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);