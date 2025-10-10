package handlers

import (
	"net/http"
	"regexp"
	"strconv"
	"webapp/config"
	"webapp/models"
)

func ShowEmployees(w http.ResponseWriter, r *http.Request) {
	var employeeItems []models.Employee
	err := config.DB.Find(&employeeItems).Error // Pass a pointer to the slice
	if err != nil {
		http.Error(w, "Error fetching employees: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Items []models.Employee
	}{
		Title: "Empleados",
		Items: employeeItems,
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

		isAlnum := regexp.MustCompile(`^[a-zA-Z0-9 ]+$`).MatchString
		if !isAlnum(departmentName) {
			http.Error(w, "Department name must contain only alphanumeric characters", http.StatusBadRequest)
			return
		}

		employee := models.Employee{
			EmployeeName:   employeeName,
			DepartmentName: departmentName,
		}

		err := config.DB.Create(&employee).Error
		if err != nil {
			http.Error(w, "Error adding employee: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employees", http.StatusSeeOther)
		return
	}

	Templates.ExecuteTemplate(w, "add_employee", nil)
}

func DeleteEmployeeByName(w http.ResponseWriter, r *http.Request) {
	employeeName := r.URL.Query().Get("employee_name")
	departmentName := r.URL.Query().Get("department_name")

	if employeeName == "" || departmentName == "" {
		http.Error(w, "Both employee_name and department_name must be provided", http.StatusBadRequest)
		return
	}

	var employee models.Employee
	err := config.DB.Where("nombre_empleado = ? AND departamento_nombre = ?", employeeName, departmentName).First(&employee).Error
	if err != nil {
		http.Error(w, "Employee not found: "+err.Error(), http.StatusNotFound)
		return
	}

	err = config.DB.Delete(&employee).Error
	if err != nil {
		http.Error(w, "Error deleting employee: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/employees", http.StatusSeeOther)
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

		data := struct {
			Employee models.Employee
		}{
			Employee: employee,
		}

		Templates.ExecuteTemplate(w, "edit_employee", data)
		return
	}

	if r.Method == http.MethodPost {
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

		employeeName := r.FormValue("employee_name")
		departmentName := r.FormValue("department_name")

		var employee models.Employee
		err = config.DB.First(&employee, id).Error
		if err != nil {
			http.Error(w, "Empleado no encontrado: "+err.Error(), http.StatusNotFound)
			return
		}

		employee.EmployeeName = employeeName
		employee.DepartmentName = departmentName

		err = config.DB.Save(&employee).Error
		if err != nil {
			http.Error(w, "Error al actualizar empleado: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employees", http.StatusSeeOther)
		return
	}
}
