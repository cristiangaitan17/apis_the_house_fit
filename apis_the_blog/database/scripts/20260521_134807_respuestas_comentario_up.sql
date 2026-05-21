-- UP: Crear tabla respuestas_comentario
CREATE TABLE IF NOT EXISTS blog.respuestas_comentario (
    id SERIAL PRIMARY KEY,
    comentario_id INTEGER REFERENCES blog.comentarios_comunidad(id) ON DELETE CASCADE,
    usuario_id INTEGER,
    contenido TEXT NOT NULL,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);