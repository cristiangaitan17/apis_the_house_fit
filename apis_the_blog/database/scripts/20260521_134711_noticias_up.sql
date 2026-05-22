-- UP: Crear tabla noticias
CREATE TABLE IF NOT EXISTS blog.noticias (
    id SERIAL PRIMARY KEY,
    categoria_id INTEGER REFERENCES blog.categorias(id) ON DELETE SET NULL,
    titulo VARCHAR(255) NOT NULL,
    contenido TEXT,
    encabezado VARCHAR(255),
    imagen_principal VARCHAR(255),
    autor_id INTEGER,
    estado VARCHAR(20) DEFAULT 'borrador',
    vistas INTEGER DEFAULT 0,
    publicado_en TIMESTAMP,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);