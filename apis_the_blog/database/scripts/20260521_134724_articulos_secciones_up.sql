-- UP: Crear tabla articulos_secciones
CREATE TABLE IF NOT EXISTS blog.articulos_secciones (
    id SERIAL PRIMARY KEY,
    articulo_id INTEGER REFERENCES blog.noticias(id) ON DELETE CASCADE,
    titulo_seccion VARCHAR(200),
    contenido TEXT,
    imagen_url VARCHAR(255),
    orden INTEGER DEFAULT 0,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);