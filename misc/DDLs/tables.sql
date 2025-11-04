USE inventario_db;

-- Tabla de departamentos
CREATE TABLE departamentos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_departamento VARCHAR(100) UNIQUE
);

-- Tabla de empleados (con departamento por nombre)
-- Agregar columna que indique si el empleado está activo
-- Agregar correo empleado
CREATE TABLE empleados (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_empleado VARCHAR(100) UNIQUE,
    departamento_nombre VARCHAR(100),
    FOREIGN KEY (departamento_nombre) REFERENCES departamentos(nombre_departamento)
);

-- Tabla de inventario (inventario general)
CREATE TABLE inventario (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_material VARCHAR(100),
    cantidad INT,
    descripcion VARCHAR(255)
);

-- Tabla de inventario de salidas
CREATE TABLE inventario_salidas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_material VARCHAR(100),
    departamento_nombre VARCHAR(100), -- Relación por nombre
    cantidad INT,
    descripcion VARCHAR(255),
    fecha DATE,
    entregado ENUM('si', 'no') DEFAULT 'no',
    empleado_nombre VARCHAR(100), -- Relación por nombre
    FOREIGN KEY (departamento_nombre) REFERENCES departamentos(nombre_departamento),
    FOREIGN KEY (empleado_nombre) REFERENCES empleados(nombre_empleado)
);

-- Tabla de inventario de entradas
CREATE TABLE inventario_entradas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fecha_entrada DATE,
    nombre_material VARCHAR(100),
    cantidad INT,
    descripcion_material VARCHAR(255) NULL,
    nombre_proveedor VARCHAR(100),
    nota TEXT NULL
);

-- Tabla de pedidos
CREATE TABLE pedidos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_material VARCHAR(100),
    nombre_proveedor VARCHAR(255),
    descripcion_material VARCHAR(255) NULL,
    cantidad_material INT,
    estado ENUM('Pendiente', 'Aprobado', 'Enviado', 'Recibido') DEFAULT 'Pendiente',
    nota TEXT NULL,
    fecha_pedido DATE,
    fecha_entrega DATE NULL
);

-- Tabla de material pendiente requisición
CREATE TABLE material_pendiente_requisicion (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_material VARCHAR(255),
    empleado_nombre VARCHAR(100), -- Relación por nombre
    departamento_nombre VARCHAR(100), -- Relación por nombre
    fecha DATE,
    FOREIGN KEY (empleado_nombre) REFERENCES empleados(nombre_empleado),
    FOREIGN KEY (departamento_nombre) REFERENCES departamentos(nombre_departamento)
);

DROP VIEW IF EXISTS vw_inventario;

CREATE VIEW vw_inventario AS
SELECT
    TRIM(i.nombre_material) as nombre_material,
    IFNULL(SUM(i.cantidad), 0) - IFNULL(SUM(s.cantidad), 0) AS total
FROM inventario_entradas i
INNER JOIN inventario_salidas s 
    ON LOWER(TRIM(i.nombre_material)) = LOWER(TRIM(s.nombre_material))
GROUP BY TRIM(i.nombre_material);

CREATE VIEW vw_inventario AS
SELECT
    TRIM(i.nombre_material) AS nombre_material,
    IFNULL(SUM(i.cantidad), 0) - IFNULL(SUM(s.cantidad), 0) AS total
FROM inventario_entradas i
LEFT JOIN inventario_salidas s 
    ON LOWER(TRIM(i.nombre_material)) = LOWER(TRIM(s.nombre_material))
GROUP BY TRIM(i.nombre_material);


SELECT * FROM vw_inventario;
SELECT * FROM inventario_entradas;
SELECT * FROM inventario_salidas;


SELECT id, fecha_pedido, fecha_entrega FROM pedidos WHERE id = 31;

SELECT
    LOWER(TRIM(i.nombre_material)) AS material,
    SUM(i.cantidad) AS entradas,
    IFNULL(SUM(s.cantidad), 0) AS salidas,
    IFNULL(SUM(i.cantidad), 0) - IFNULL(SUM(s.cantidad), 0) AS total
FROM inventario_entradas i
LEFT JOIN (
    SELECT nombre_material, SUM(cantidad) AS cantidad
    FROM inventario_salidas
    GROUP BY nombre_material
) s ON LOWER(TRIM(i.nombre_material)) = LOWER(TRIM(s.nombre_material))
GROUP BY LOWER(TRIM(i.nombre_material));





DROP VIEW IF EXISTS vw_inventario;


