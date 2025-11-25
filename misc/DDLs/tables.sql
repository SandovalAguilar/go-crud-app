USE inventario_db;

CREATE TABLE departamentos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_departamento VARCHAR(100) UNIQUE
);

CREATE TABLE empleados (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_empleado VARCHAR(100) UNIQUE,
    departamento_nombre VARCHAR(100),
    FOREIGN KEY (departamento_nombre) REFERENCES departamentos(nombre_departamento)
);

CREATE TABLE inventario (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_material VARCHAR(100),
    cantidad INT,
    descripcion VARCHAR(255)
);

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

CREATE TABLE material_pendiente_requisicion (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_material VARCHAR(255),
    departamento_nombre VARCHAR(100),
    cantidad INT,
    descripcion VARCHAR(255),
    fecha DATE,
    empleado_nombre VARCHAR(100), 
    requisicion ENUM('Pendiente', 'Recibida') DEFAULT 'Pendiente',
    FOREIGN KEY (empleado_nombre) REFERENCES empleados(nombre_empleado),
    FOREIGN KEY (departamento_nombre) REFERENCES departamentos(nombre_departamento)
);
	
DROP TABLE pedidos;

CREATE TABLE pedidos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre_material VARCHAR(100),
    nombre_proveedor VARCHAR(255) NULL,
    descripcion_material VARCHAR(255) NULL,
    cantidad_material INT,
    estado ENUM('Pendiente', 'Recibido') DEFAULT 'Pendiente',
    nota TEXT NULL,
    fecha_pedido DATE,
    fecha_entrega DATE NULL
);

CREATE TABLE inventario_entradas (
  id INT NOT NULL AUTO_INCREMENT,
  fecha_entrada DATE DEFAULT NULL,
  nombre_material VARCHAR(100) DEFAULT NULL,
  cantidad INT DEFAULT NULL,
  descripcion_material VARCHAR(255) DEFAULT NULL,
  nombre_proveedor VARCHAR(100) DEFAULT NULL,
  nota TEXT,
  PRIMARY KEY (id)
);