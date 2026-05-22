-- UP: Crear tabla categorias
CREATE TABLE IF NOT EXISTS blog.categorias (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(80) NOT NULL,
    seccion_lugar VARCHAR(100),
    descripcion TEXT,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);