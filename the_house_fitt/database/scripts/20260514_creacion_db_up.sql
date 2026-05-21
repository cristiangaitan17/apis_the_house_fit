-- Creacion del Schema
CREATE SCHEMA IF NOT EXISTS tienda;

-- Creacion de Tabla categorias_producto
CREATE TABLE tienda.categorias_producto (
    id                  SERIAL              PRIMARY KEY,
    nombre              VARCHAR(80)         NOT NULL,
    icono               VARCHAR(50),
    "Descripcion"       VARCHAR(200),
    activo              BOOLEAN             NOT NULL DEFAULT TRUE,
    "Fecha_modificacion" TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    "Fecha_creacion"    TIMESTAMP           DEFAULT CURRENT_TIMESTAMP
);

-- Creacion de Tabla productos
CREATE TABLE tienda.productos (
    id                  SERIAL              PRIMARY KEY,
    categoria_id        INTEGER             NOT NULL,
    nombre              VARCHAR(100)        NOT NULL,
    marca               VARCHAR(100),
    sabor               VARCHAR(100),
    precio              NUMERIC(10,2)       NOT NULL,
    precio_original     NUMERIC(10,2),
    "Impuesto"          NUMERIC(5,2),
    porcion_g           INTEGER,
    porciones           INTEGER,
    descripcion         TEXT,
    imagen_url          VARCHAR(200),
    stock               INTEGER             DEFAULT 0,
    es_nuevo            BOOLEAN             DEFAULT FALSE,
    descuento_pct       INTEGER             DEFAULT 0,
    calificacion_prom   NUMERIC(2,1)        DEFAULT 0.0,
    total_resenas       INTEGER             DEFAULT 0,
    activo              BOOLEAN             NOT NULL DEFAULT TRUE,
    "Fecha_modificacion" TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    "Fecha_creacion"    TIMESTAMP           DEFAULT CURRENT_TIMESTAMP
);

-- Creacion de Tabla carrito
CREATE TABLE tienda.carrito (
    id                  SERIAL              PRIMARY KEY,
    usuario_id          INTEGER             NOT NULL,
    activo              BOOLEAN             NOT NULL DEFAULT TRUE,
    "Fecha_modificacion" TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    "Fecha_creacion"    TIMESTAMP           DEFAULT CURRENT_TIMESTAMP
);

-- Creacion de Tabla carrito_items
CREATE TABLE tienda.carrito_items (
    id_carrito          SERIAL              PRIMARY KEY,
    carrito_id          INTEGER             NOT NULL,
    producto_id         INTEGER             NOT NULL,
    cantidad            INTEGER             NOT NULL,
    precio_unitario     NUMERIC(10,2)       NOT NULL,
    activo              BOOLEAN             NOT NULL DEFAULT TRUE,
    "Fecha_modificacion" TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    "Fecha_creacion"    TIMESTAMP           DEFAULT CURRENT_TIMESTAMP
);

-- Creacion de Tabla pedidos
CREATE TABLE tienda.pedidos (
    id_pedido           SERIAL              PRIMARY KEY,
    usuario_id          INTEGER             NOT NULL,
    subtotal            NUMERIC(10,2)       DEFAULT 0,
    envio               NUMERIC(10,2)       DEFAULT 0,
    impuesto            NUMERIC(10,2)       DEFAULT 0,
    total               NUMERIC(10,2)       DEFAULT 0,
    metodo_pago         VARCHAR(50),
    estado_pago         VARCHAR(50)         DEFAULT 'pendiente',
    estado_pedido       VARCHAR(50)         DEFAULT 'procesando',
    nombre_receptor     VARCHAR(100),
    apellido_receptor   VARCHAR(100),
    email_receptor      VARCHAR(100),
    telefono            VARCHAR(20),
    direccion_envio     VARCHAR(200),
    activo              BOOLEAN             NOT NULL DEFAULT TRUE,
    "Fecha_modificacion" TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    "Fecha_creacion"    TIMESTAMP           DEFAULT CURRENT_TIMESTAMP
);

-- Creacion de Tabla pedido_items
CREATE TABLE tienda.pedido_items (
    indexes             SERIAL              PRIMARY KEY,
    pedido_id           INTEGER