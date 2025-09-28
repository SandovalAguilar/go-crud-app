package models

type Department struct {
	ID             uint   `gorm:"primaryKey"`
	DepartmentName string `gorm:"type:varchar(100);unique;not null;column:nombre_departamento"`
}

func (Department) TableName() string {
	return "departamentos"
}

type Employee struct {
	ID             uint       `gorm:"primaryKey"`
	EmployeeName   string     `gorm:"type:varchar(100);unique;not null;column:nombre_empleado"`
	DepartmentName string     `gorm:"type:varchar(100);column:departamento_nombre"`        // FK a Department
	Department     Department `gorm:"foreignKey:DepartmentName;references:DepartmentName"` // relación
}

func (Employee) TableName() string {
	return "empleados"
}

type Inventory struct {
	ID           uint   `gorm:"primaryKey"`
	MaterialName string `gorm:"type:varchar(100);not null;column:nombre_material"`
	Quantity     int    `gorm:"not null;column:cantidad"`
	Description  string `gorm:"type:varchar(255);column:descripcion"`
}

func (Inventory) TableName() string {
	return "inventario"
}

type Order struct {
	ID                  uint    `gorm:"primaryKey;autoIncrement"`
	MaterialName        string  `gorm:"type:varchar(100);column:nombre_material"`
	SupplierName        string  `gorm:"type:varchar(255);column:nombre_proveedor"`
	MaterialDescription *string `gorm:"type:varchar(255);column:descripcion_material"`
	MaterialQuantity    int     `gorm:"column:cantidad_material"`
	Status              string  `gorm:"type:enum('Pendiente', 'Aprobado', 'Enviado', 'Recibido');default:'Pendiente';column:estado"`
	Note                *string `gorm:"type:text;column:nota"`
	RequestDate         string  `gorm:"type:date;column:fecha_pedido"`
	DeliveryDate        *string `gorm:"type:date;column:fecha_entrega"`
}

func (Order) TableName() string {
	return "pedidos"
}

type InventoryEntry struct {
	ID                  uint    `gorm:"primaryKey;autoIncrement"`
	EntryDate           string  `gorm:"type:date;column:fecha_entrada"`
	MaterialName        string  `gorm:"type:varchar(100);column:nombre_material"`
	Quantity            int     `gorm:"column:cantidad"`
	MaterialDescription *string `gorm:"type:varchar(255);column:descripcion_material"`
	SupplierName        string  `gorm:"type:varchar(100);column:nombre_proveedor"`
	Note                *string `gorm:"type:text;column:nota"`
}

func (InventoryEntry) TableName() string {
	return "inventario_entradas"
}

type InventoryOutput struct {
	ID             uint    `gorm:"primaryKey;autoIncrement"`
	MaterialName   string  `gorm:"type:varchar(100);column:nombre_material"`
	DepartmentName string  `gorm:"type:varchar(100);column:departamento_nombre"`
	Quantity       int     `gorm:"column:cantidad"`
	Description    *string `gorm:"type:varchar(255);column:descripcion"`
	Date           string  `gorm:"type:date;column:fecha"`
	Delivered      string  `gorm:"type:enum('si','no');default:'no';column:entregado"`
	EmployeeName   string  `gorm:"type:varchar(100);column:empleado_nombre"`
}

func (InventoryOutput) TableName() string {
	return "inventario_salidas"
}
