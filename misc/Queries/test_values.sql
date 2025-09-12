INSERT INTO inventario_db.inventario_salidas(departamento, cantidad, descripcion, fecha, entregado, recibio)
VALUES 
('Ventas', 10, 'Producto A', '2025-08-14', 'si', 'Juan Perez'),
('Logística', 15, 'Producto B', '2025-08-13', 'no', 'Maria Garcia'),
('Marketing', 8, 'Producto C', '2025-08-12', 'si', 'Carlos Fernandez'),
('Compras', 5, 'Producto D', '2025-08-11', 'no', 'Ana Lopez'),
('Soporte', 12, 'Producto E', '2025-08-10', 'si', 'Luis Martinez'),
('Ventas', 20, 'Producto F', '2025-08-09', 'no', 'Pedro Sanchez'),
('Logística', 7, 'Producto G', '2025-08-08', 'si', 'Laura Torres'),
('Marketing', 25, 'Producto H', '2025-08-07', 'no', 'Juan Carlos Ramirez'),
('Compras', 18, 'Producto I', '2025-08-06', 'si', 'Eva Ruiz'),
('Soporte', 10, 'Producto J', '2025-08-05', 'no', 'Fernando Gonzalez');

INSERT INTO inventario_entradas (fecha, nombre_proveedor, descripcion_material, cantidad, nota)
VALUES
('2025-08-14', 'Proveedor A', 'Material de construcción', 100, 'Entrega programada para la próxima semana'),
('2025-08-13', 'Proveedor B', 'Pintura industrial', 50, 'Entrega urgente por alta demanda'),
('2025-08-12', 'Proveedor C', 'Cables eléctricos', 200, 'Material recibido con retraso'),
('2025-08-11', 'Proveedor D', 'Lentes de seguridad', 150, 'Pedido completo, en perfecto estado'),
('2025-08-10', 'Proveedor E', 'Herramientas manuales', 75, 'Entrega programada para mañana'),
('2025-08-09', 'Proveedor F', 'Aceite lubricante', 300, 'Material usado en mantenimiento de equipos'),
('2025-08-08', 'Proveedor G', 'Tornillos y clavos', 500, 'Entrega parcial, falta un lote'),
('2025-08-07', 'Proveedor H', 'Material de oficina', 30, 'Producto de reposición para oficina'),
('2025-08-06', 'Proveedor I', 'Placas metálicas', 120, 'Pedido entregado a tiempo'),
('2025-08-05', 'Proveedor J', 'Cemento y concreto', 250, 'Material entregado para proyecto de construcción');

INSERT INTO control_inventario (material_oficina, descripcion_oficina, cantidad_oficina, 
                                material_intendencia, descripcion_intendencia, cantidad_intendencia, 
                                consumibles, descripcion_consumibles, cantidad_consumibles, 
                                material_mtto, descripcion_mtto, cantidad_mtto)
VALUES
    ('Silla', 'Silla ergonómica de oficina', 10, 'Limpieza', 'Detergente multiusos', 50, 
     'Papel higiénico', 'Papel higiénico de 2 capas', 100, 'Destornillador', 'Destornillador plano', 20),
    ('Escritorio', 'Escritorio de madera', 15, 'Limpieza', 'Desinfectante', 30, 
     'Toallas de papel', 'Paquete de toallas absorbentes', 200, 'Martillo', 'Martillo de 500g', 25),
    ('Computadora', 'PC de escritorio', 5, 'Limpieza', 'Limpiador de pantallas', 60, 
     'Papel toalla', 'Papel toalla industrial', 150, 'Taladro', 'Taladro eléctrico', 10),
    ('Cajón', 'Cajón de almacenamiento', 8, 'Limpieza', 'Limpiador de muebles', 40, 
     'Gafas de seguridad', 'Gafas de protección para trabajo', 20, 'Llave inglesa', 'Llave ajustable', 12),
    ('Cama', 'Cama de descanso', 3, 'Limpieza', 'Limpiador de pisos', 25, 
     'Máscara facial', 'Máscara de protección facial', 30, 'Sierra', 'Sierra de mano', 5),
    ('Archivador', 'Archivador metálico', 10, 'Limpieza', 'Limpiador de vidrios', 70, 
     'Guantes desechables', 'Guantes de látex', 100, 'Alicates', 'Alicates de corte', 15),
    ('Silla de ruedas', 'Silla de ruedas hospitalaria', 2, 'Limpieza', 'Limpiador antibacterial', 15, 
     'Botellas de agua', 'Agua mineral', 200, 'Cinta métrica', 'Cinta métrica de 5m', 8),
    ('Cortina', 'Cortina para oficina', 5, 'Limpieza', 'Limpiador de alfombras', 20, 
     'Cartuchos de tinta', 'Cartuchos de tinta para impresora', 40, 'Compás', 'Compás de dibujo', 18),
    ('Lámpara', 'Lámpara de escritorio', 12, 'Limpieza', 'Limpiador de teclados', 100, 
     'Papel kraft', 'Papel kraft para envolver', 150, 'Cúter', 'Cúter industrial', 22),
    ('Impresora', 'Impresora láser', 4, 'Limpieza', 'Limpiador de superficies', 35, 
     'Carteles informativos', 'Carteles para avisos', 50, 'Escuadra', 'Escuadra de acero', 10);

INSERT INTO pedidos (fecha, descripcion_material_oficina, cantidad_material_oficina, descripcion_material_intendencia, cantidad_material_intendencia)
VALUES
('2025-08-01', 'Papel Bond A4', 100, 'Desinfectante', 50),
('2025-08-02', 'Lápices de madera', 200, 'Jabón Líquido', 30),
('2025-08-03', 'Carpetas tamaño oficio', 150, 'Toallas de papel', 40),
('2025-08-04', 'Marcadores de colores', 80, 'Papel higiénico', 100),
('2025-08-05', 'Grapas', 500, 'Limpiador multiusos', 60),
('2025-08-06', 'Papel para impresora', 120, 'Desinfectante de manos', 45),
('2025-08-07', 'Resaltadores', 100, 'Toallitas húmedas', 75),
('2025-08-08', 'Cuadernos', 50, 'Tijeras', 20),
('2025-08-09', 'Cinta adhesiva', 60, 'Escoba', 15),
('2025-08-10', 'Sillas ergonómicas', 10, 'Limpiador de vidrios', 25);

INSERT INTO material_pendiente_requisicion (nombre_material, solicitante, departamento, fecha)
VALUES
('Computadora portátil', 'Juan Pérez', 'TI', '2025-08-10'),
('Proyector', 'María López', 'Recursos Humanos', '2025-08-12'),
('Silla ergonómica', 'Carlos Gómez', 'Oficina', '2025-08-15'),
('Impresora láser', 'Ana Torres', 'Administración', '2025-08-17'),
('Escáner de documentos', 'Luis Rodríguez', 'TI', '2025-08-20'),
('Pantalla de proyección', 'Pedro Sánchez', 'Educación', '2025-08-22'),
('Teclado mecánico', 'Sofía Herrera', 'Oficina', '2025-08-23'),
('Cable de red', 'Ricardo Fernández', 'TI', '2025-08-24'),
('Cámara web', 'Lucía Ramírez', 'Marketing', '2025-08-25'),
('Micrófono', 'Carlos Díaz', 'Recursos Humanos', '2025-08-26');

