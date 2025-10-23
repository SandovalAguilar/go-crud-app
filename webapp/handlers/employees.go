package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"webapp/config"
	"webapp/models"
)

func ShowEmployees(w http.ResponseWriter, r *http.Request) {
	var employeeItems []models.Employee
	err := config.DB.Find(&employeeItems).Error
	if err != nil {
		http.Error(w, "Error fetching employees: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Items   []models.Employee
		Error   string
		Success string
	}{
		Title:   "Empleados",
		Items:   employeeItems,
		Error:   r.URL.Query().Get("error"),
		Success: r.URL.Query().Get("success"),
	}

	err = Templates.ExecuteTemplate(w, "employees", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		employeeName := r.FormValue("employee_name")
		departmentName := r.FormValue("department_name")

		if employeeName == "" || departmentName == "" {
			http.Redirect(w, r, "/employees?error="+url.QueryEscape("Todos los campos son requeridos"), http.StatusSeeOther)
			return
		}

		employee := models.Employee{
			EmployeeName:   employeeName,
			DepartmentName: departmentName,
		}

		err := config.DB.Create(&employee).Error
		if err != nil {
			errorMsg := err.Error()
			// Check for foreign key constraint error (department doesn't exist)
			if strings.Contains(errorMsg, "foreign key constraint") || strings.Contains(errorMsg, "1452") {
				http.Redirect(w, r, "/employees?error="+url.QueryEscape("No se puede agregar el empleado porque el departamento especificado no existe. Por favor, cree el departamento primero."), http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/employees?error="+url.QueryEscape("Error al agregar empleado: "+errorMsg), http.StatusSeeOther)
			}
			return
		}

		http.Redirect(w, r, "/employees?success="+url.QueryEscape("Empleado agregado exitosamente"), http.StatusSeeOther)
		return
	}

	// Fetch all departments for the dropdown
	var departments []models.Department
	config.DB.Find(&departments)

	data := struct {
		Departments []models.Department
	}{
		Departments: departments,
	}

	Templates.ExecuteTemplate(w, "add_employee", data)
}

func DeleteEmployeeByName(w http.ResponseWriter, r *http.Request) {
	employeeName := r.URL.Query().Get("employee_name")
	departmentName := r.URL.Query().Get("department_name")

	if employeeName == "" || departmentName == "" {
		http.Redirect(w, r, "/employees?error="+url.QueryEscape("Debe proporcionar el nombre del empleado y el departamento"), http.StatusSeeOther)
		return
	}

	var employee models.Employee
	err := config.DB.Where("nombre_empleado = ? AND departamento_nombre = ?", employeeName, departmentName).First(&employee).Error
	if err != nil {
		http.Redirect(w, r, "/employees?error="+url.QueryEscape("Empleado no encontrado"), http.StatusSeeOther)
		return
	}

	err = config.DB.Delete(&employee).Error
	if err != nil {
		// Check if it's a foreign key constraint error
		errorMsg := err.Error()
		if strings.Contains(errorMsg, "foreign key constraint") || strings.Contains(errorMsg, "1451") {
			http.Redirect(w, r, "/employees?error="+url.QueryEscape("No se puede eliminar el empleado porque tiene registros asociados en el inventario"), http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/employees?error="+url.QueryEscape("Error al eliminar empleado: "+errorMsg), http.StatusSeeOther)
		}
		return
	}

	http.Redirect(w, r, "/employees?success="+url.QueryEscape("Empleado eliminado exitosamente"), http.StatusSeeOther)
}

func EditEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "ID del empleado no proporcionado", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		var employee models.Employee
		err = config.DB.First(&employee, id).Error
		if err != nil {
			http.Error(w, "Empleado no encontrado: "+err.Error(), http.StatusNotFound)
			return
		}

		// Fetch all departments for the dropdown
		var departments []models.Department
		config.DB.Find(&departments)

		data := struct {
			Employee    models.Employee
			Departments []models.Department
		}{
			Employee:    employee,
			Departments: departments,
		}
		Templates.ExecuteTemplate(w, "edit_employee", data)
		return
	}

	if r.Method == http.MethodPost {
		idStr := r.URL.Query().Get("id")
		page := r.URL.Query().Get("page")

		if idStr == "" {
			http.Error(w, "ID del empleado no proporcionado", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		employeeName := r.FormValue("employee_name")
		departmentName := r.FormValue("department_name")

		var employee models.Employee
		err = config.DB.First(&employee, id).Error
		if err != nil {
			http.Redirect(w, r, "/employees?page="+page+"&error="+url.QueryEscape("Empleado no encontrado"), http.StatusSeeOther)
			return
		}

		employee.EmployeeName = employeeName
		employee.DepartmentName = departmentName
		err = config.DB.Save(&employee).Error
		if err != nil {
			errorMsg := err.Error()
			// Check for foreign key constraint error (department doesn't exist)
			if strings.Contains(errorMsg, "foreign key constraint") || strings.Contains(errorMsg, "1452") {
				http.Redirect(w, r, "/employees?page="+page+"&error="+url.QueryEscape("No se puede actualizar el empleado porque el departamento especificado no existe. Por favor, cree el departamento primero."), http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/employees?page="+page+"&error="+url.QueryEscape("Error al actualizar empleado: "+errorMsg), http.StatusSeeOther)
			}
			return
		}

		http.Redirect(w, r, "/employees?page="+page+"&success="+url.QueryEscape("Empleado actualizado exitosamente"), http.StatusSeeOther)
		return
	}
}
