package models

// Department representa la tabla departamentos
type Department struct {
	ID             uint   `gorm:"primaryKey"`
	DepartmentName string `gorm:"type:varchar(100);unique;not null;column:nombre_departamento"`
}

func (Department) TableName() string {
	return "departamentos"
}

// Employee representa la tabla empleados
type Employee struct {
	ID             uint       `gorm:"primaryKey"`
	EmployeeName   string     `gorm:"type:varchar(100);unique;not null;column:nombre_empleado"`
	DepartmentName string     `gorm:"type:varchar(100);column:departamento_nombre"`        // FK a Department
	Department     Department `gorm:"foreignKey:DepartmentName;references:DepartmentName"` // relación
}

// TableName sobrescribe el nombre de la tabla para Employee
func (Employee) TableName() string {
	return "empleados"
}

// Inventory representa la tabla inventario
type Inventory struct {
	ID           uint   `gorm:"primaryKey"`
	MaterialName string `gorm:"type:varchar(100);not null;column:nombre_material"`
	Quantity     int    `gorm:"not null;column:cantidad"`
	Description  string `gorm:"type:varchar(255);column:descripcion"`
}

// TableName sobrescribe el nombre de la tabla para Inventory
func (Inventory) TableName() string {
	return "inventario"
}
