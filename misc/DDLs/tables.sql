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

CREATE VIEW vw_inventario AS
SELECT
    i.nombre_material,
    IFNULL(SUM(i.cantidad), 0) - IFNULL(SUM(s.cantidad), 0) AS total
FROM inventario_entradas i
LEFT JOIN inventario_salidas s ON i.nombre_material = s.nombre_material
GROUP BY i.nombre_material;

SELECT * FROM vw_inventario;

