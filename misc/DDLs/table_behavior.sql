USE inventario_db;

-- UPDATE FOREIGN KEYS TO USE CASCADE

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

-- 5. Update material_pendiente_requisicion - cascade department name updates
ALTER TABLE material_pendiente_requisicion
DROP FOREIGN KEY material_pendiente_requisicion_ibfk_2;

ALTER TABLE material_pendiente_requisicion
ADD CONSTRAINT material_pendiente_requisicion_ibfk_2
FOREIGN KEY (departamento_nombre)
REFERENCES departamentos(nombre_departamento)
ON UPDATE CASCADE
ON DELETE RESTRICT;

USE inventario_db;

-- TRIGGER 1: Registrar entrada automáticamente cuando un pedido es recibido

DELIMITER $$

CREATE TRIGGER trg_pedido_recibido_registrar_entrada
AFTER UPDATE ON pedidos
FOR EACH ROW
BEGIN
    -- Solo ejecutar si el estado cambió a 'Recibido'
    IF NEW.estado = 'Recibido' AND OLD.estado != 'Recibido' THEN
        INSERT INTO inventario_entradas (
            fecha_entrada,
            nombre_material,
            cantidad,
            descripcion_material,
            nombre_proveedor,
            nota
        ) VALUES (
            CURDATE(),
            NEW.nombre_material,
            NEW.cantidad_material,
            NEW.descripcion_material,
            NEW.nombre_proveedor,
            CONCAT('Entrada automática desde pedido #', NEW.id,
                   CASE WHEN NEW.nota IS NOT NULL
                        THEN CONCAT(' - ', NEW.nota)
                        ELSE ''
                   END)
        );
    END IF;
END$$

DELIMITER ;

----

	DELIMITER $$
	DROP TRIGGER IF EXISTS trg_pedido_recibido_registrar_entrada$$
	CREATE TRIGGER trg_pedido_recibido_registrar_entrada
	AFTER UPDATE ON pedidos
	FOR EACH ROW
	BEGIN
		-- Solo ejecutar si:
		-- 1. El estado cambió a 'Recibido'
		-- 2. La nota NO contiene la marca especial de entrada manual
		IF NEW.estado = 'Recibido'
		   AND OLD.estado != 'Recibido'
		   AND (NEW.nota IS NULL OR NEW.nota NOT LIKE '%# Pedido parcial recibido%') THEN

			INSERT INTO inventario_entradas (
				fecha_entrada,
				nombre_material,
				cantidad,
				descripcion_material,
				nombre_proveedor,
				nota
			) VALUES (
				CURDATE(),
				NEW.nombre_material,
				NEW.cantidad_material,
				NEW.descripcion_material,
				NEW.nombre_proveedor,
				CONCAT('Entrada automática desde pedido #', NEW.id,
					   CASE WHEN NEW.nota IS NOT NULL
							AND NEW.nota NOT LIKE '%# Pedido parcial recibido%'
							THEN CONCAT(' - ', NEW.nota)
							ELSE ''
					   END)
			);
		END IF;
	END$$
	DELIMITER ;
