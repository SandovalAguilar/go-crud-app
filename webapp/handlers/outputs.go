package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"webapp/config"
	"webapp/models"
)

func ShowOutputs(w http.ResponseWriter, r *http.Request) {
	var outputs []models.InventoryOutput
	err := config.DB.Find(&outputs).Error
	if err != nil {
		http.Error(w, "Error fetching inventory outputs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Outputs []models.InventoryOutput
		Error   string
		Success string
	}{
		Title:   "Inventario de Salidas",
		Outputs: outputs,
		Error:   r.URL.Query().Get("error"),
		Success: r.URL.Query().Get("success"),
	}

	err = Templates.ExecuteTemplate(w, "outputs", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func DeleteOutput(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	materialName := r.URL.Query().Get("material_name")

	if idStr == "" || materialName == "" {
		http.Redirect(w, r, "/outputs?error="+url.QueryEscape("Faltan parámetros requeridos"), http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/outputs?error="+url.QueryEscape("ID inválido"), http.StatusSeeOther)
		return
	}

	err = config.DB.Where("id = ? AND nombre_material = ?", id, materialName).Delete(&models.InventoryOutput{}).Error
	if err != nil {
		errorMsg := err.Error()
		if strings.Contains(errorMsg, "foreign key constraint") {
			http.Redirect(w, r, "/outputs?error="+url.QueryEscape("No se puede eliminar la salida porque tiene registros asociados"), http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/outputs?error="+url.QueryEscape("Error al eliminar salida: "+errorMsg), http.StatusSeeOther)
		}
		return
	}

	http.Redirect(w, r, "/outputs?success="+url.QueryEscape("Salida eliminada exitosamente"), http.StatusSeeOther)
}

func AddOutput(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		materialName := r.FormValue("material_name")
		departmentName := r.FormValue("department_name")
		quantity := r.FormValue("quantity")
		description := r.FormValue("description")
		date := r.FormValue("date")
		delivered := r.FormValue("delivered")
		employeeName := r.FormValue("employee_name")

		if materialName == "" || departmentName == "" || quantity == "" || date == "" || delivered == "" || employeeName == "" {
			http.Redirect(w, r, "/outputs?error="+url.QueryEscape("Faltan campos requeridos"), http.StatusSeeOther)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			http.Redirect(w, r, "/outputs?error="+url.QueryEscape("Cantidad inválida"), http.StatusSeeOther)
			return
		}

		var parsedDescription *string
		if description != "" {
			parsedDescription = &description
		}

		output := models.InventoryOutput{
			MaterialName:   materialName,
			DepartmentName: departmentName,
			Quantity:       quantityInt,
			Description:    parsedDescription,
			Date:           date,
			Delivered:      delivered,
			EmployeeName:   employeeName,
		}

		err = config.DB.Create(&output).Error
		if err != nil {
			errorMsg := err.Error()
			if strings.Contains(errorMsg, "foreign key constraint") || strings.Contains(errorMsg, "1452") {
				http.Redirect(w, r, "/outputs?error="+url.QueryEscape("El departamento o empleado especificado no existe. Por favor, verifique los datos."), http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/outputs?error="+url.QueryEscape("Error al agregar salida: "+errorMsg), http.StatusSeeOther)
			}
			return
		}

		http.Redirect(w, r, "/outputs?success="+url.QueryEscape("Salida agregada exitosamente"), http.StatusSeeOther)
		return
	}

	// Fetch departments and employees for dropdowns
	var departments []models.Department
	var employees []models.Employee
	config.DB.Find(&departments)
	config.DB.Find(&employees)

	data := struct {
		Departments []models.Department
		Employees   []models.Employee
	}{
		Departments: departments,
		Employees:   employees,
	}

	if err := Templates.ExecuteTemplate(w, "add_output", data); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func EditOutput(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		materialName := r.FormValue("material_name")
		departmentName := r.FormValue("department_name")
		quantity := r.FormValue("quantity")
		description := r.FormValue("description")
		date := r.FormValue("date")
		delivered := r.FormValue("delivered")
		employeeName := r.FormValue("employee_name")

		if idStr == "" || materialName == "" || departmentName == "" || quantity == "" || date == "" || delivered == "" || employeeName == "" {
			http.Redirect(w, r, "/outputs?error="+url.QueryEscape("Faltan campos requeridos"), http.StatusSeeOther)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			http.Redirect(w, r, "/outputs?error="+url.QueryEscape("Cantidad inválida"), http.StatusSeeOther)
			return
		}

		var parsedDescription *string
		if description != "" {
			parsedDescription = &description
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Redirect(w, r, "/outputs?error="+url.QueryEscape("ID inválido"), http.StatusSeeOther)
			return
		}

		err = config.DB.Model(&models.InventoryOutput{}).Where("id = ?", id).Updates(models.InventoryOutput{
			MaterialName:   materialName,
			DepartmentName: departmentName,
			Quantity:       quantityInt,
			Description:    parsedDescription,
			Date:           date,
			Delivered:      delivered,
			EmployeeName:   employeeName,
		}).Error
		if err != nil {
			errorMsg := err.Error()
			if strings.Contains(errorMsg, "foreign key constraint") || strings.Contains(errorMsg, "1452") {
				http.Redirect(w, r, "/outputs?error="+url.QueryEscape("El departamento o empleado especificado no existe. Por favor, verifique los datos."), http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/outputs?error="+url.QueryEscape("Error al actualizar salida: "+errorMsg), http.StatusSeeOther)
			}
			return
		}

		http.Redirect(w, r, "/outputs?success="+url.QueryEscape("Salida actualizada exitosamente"), http.StatusSeeOther)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Falta el ID de salida", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de salida inválido", http.StatusBadRequest)
		return
	}

	var output models.InventoryOutput
	err = config.DB.First(&output, id).Error
	if err != nil {
		http.Error(w, "Salida no encontrada: "+err.Error(), http.StatusNotFound)
		return
	}

	// Fetch departments and employees for dropdowns
	var departments []models.Department
	var employees []models.Employee
	config.DB.Find(&departments)
	config.DB.Find(&employees)

	data := struct {
		Output      models.InventoryOutput
		Departments []models.Department
		Employees   []models.Employee
	}{
		Output:      output,
		Departments: departments,
		Employees:   employees,
	}

	err = Templates.ExecuteTemplate(w, "edit_output", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla: "+err.Error(), http.StatusInternalServerError)
	}
}
