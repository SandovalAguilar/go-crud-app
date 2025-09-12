-- Crear base de datos
CREATE DATABASE inventario_db;
USE inventario_db;

-- Tabla de departamentos
CREATE TABLE departamentos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_departamento VARCHAR(100)
);

-- Tabla de empleados (con departamento)
CREATE TABLE empleados (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_empleado VARCHAR(100),
    departamento_id INT,  -- Relación con la tabla departamentos
    FOREIGN KEY (departamento_id) REFERENCES departamentos(id)
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
    departamento_id INT, -- Relación con la tabla departamentos
    cantidad INT,
    descripcion VARCHAR(255),
    fecha DATE,
    entregado ENUM('si', 'no') DEFAULT 'no',
    empleado_id INT, -- Relación con la tabla empleados (quien recibió)
    FOREIGN KEY (departamento_id) REFERENCES departamentos(id),
    FOREIGN KEY (empleado_id) REFERENCES empleados(id)
);

-- Tabla de inventario de entradas
CREATE TABLE inventario_entradas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fecha_entrada DATE,
    nombre_material VARCHAR(100),
    cantidad INT,
    descripcion_material VARCHAR(255),
    nombre_proveedor VARCHAR(100),
    cantidad INT,
    nota TEXT,
    departamento_id INT,  -- Relación con la tabla departamentos
    FOREIGN KEY (departamento_id) REFERENCES departamentos(id)
);

-- Tabla de pedidos
CREATE TABLE pedidos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    descripcion_material VARCHAR(255),
    cantidad_material INT,
    estado ENUM('Pendiente', 'Aprobado', 'Enviado', 'Recibido') DEFAULT 'Pendiente',
    fecha_pedido DATE,
    fecha_entrega DATE
);

-- Tabla de material pendiente requisición (relacionada con salidas no entregadas)
CREATE TABLE material_pendiente_requisicion (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_material VARCHAR(255),
    empleado_id INT, -- Relación con la tabla empleados (quien solicitó)
    departamento_id INT, -- Relación con la tabla departamentos
    fecha DATE,
    salida_id INT,  -- Relación con la tabla inventario_salidas (material pendiente)
    FOREIGN KEY (empleado_id) REFERENCES empleados(id),
    FOREIGN KEY (departamento_id) REFERENCES departamentos(id),
    FOREIGN KEY (salida_id) REFERENCES inventario_salidas(id)  -- Relación con la tabla de salidas
);