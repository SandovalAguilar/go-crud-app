USE inventario_db;

-- =====================================================
-- UPDATE FOREIGN KEYS TO USE CASCADE
-- =====================================================

-- 1. Update empleados table - cascade department name updates
ALTER TABLE empleados 
DROP FOREIGN KEY empleados_ibfk_1;

ALTER TABLE empleados 
ADD CONSTRAINT empleados_ibfk_1 
FOREIGN KEY (departamento_nombre) 
REFERENCES departamentos(nombre_departamento) 
ON UPDATE CASCADE 
ON DELETE RESTRICT;

-- 2. Update inventario_salidas - cascade department name updates
ALTER TABLE inventario_salidas 
DROP FOREIGN KEY inventario_salidas_ibfk_1;

ALTER TABLE inventario_salidas 
ADD CONSTRAINT inventario_salidas_ibfk_1 
FOREIGN KEY (departamento_nombre) 
REFERENCES departamentos(nombre_departamento) 
ON UPDATE CASCADE 
ON DELETE RESTRICT;

-- 3. Update inventario_salidas - cascade employee name updates
ALTER TABLE inventario_salidas 
DROP FOREIGN KEY inventario_salidas_ibfk_2;

ALTER TABLE inventario_salidas 
ADD CONSTRAINT inventario_salidas_ibfk_2 
FOREIGN KEY (empleado_nombre) 
REFERENCES empleados(nombre_empleado) 
ON UPDATE CASCADE 
ON DELETE RESTRICT;

-- 4. Update material_pendiente_requisicion - cascade employee name updates
ALTER TABLE material_pendiente_requisicion 
DROP FOREIGN KEY material_pendiente_requisicion_ibfk_1;

ALTER TABLE material_pendiente_requisicion 
ADD CONSTRAINT material_pendiente_requisicion_ibfk_1 
FOREIGN KEY (empleado_nombre) 
REFERENCES empleados(nombre_empleado) 
ON UPDATE CASCADE 
ON DELETE RESTRICT;

