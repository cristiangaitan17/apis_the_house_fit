-- UP: Crear tabla dieta_comidas
CREATE TABLE IF NOT EXISTS blog.dieta_comidas (
    id SERIAL PRIMARY KEY,
    dieta_id INTEGER REFERENCES blog.nutricion(id) ON DELETE CASCADE,
    tiempo_comida VARCHAR(20),
    descripcion TEXT,
    orden INTEGER,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);