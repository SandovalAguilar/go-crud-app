package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"webapp/config"
	"webapp/models"
)

func ShowPendings(w http.ResponseWriter, r *http.Request) {
	var outputs []models.Pendings
	err := config.DB.Find(&outputs).Error
	if err != nil {
		http.Error(w, "Error fetching pendings: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Outputs []models.Pendings
		Error   string
		Success string
	}{
		Title:   "Material Pendiente de Requisición",
		Outputs: outputs,
		Error:   r.URL.Query().Get("error"),
		Success: r.URL.Query().Get("success"),
	}

	err = Templates.ExecuteTemplate(w, "pendings", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func DeletePending(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	materialName := r.URL.Query().Get("material_name")

	if idStr == "" || materialName == "" {
		http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Faltan parámetros requeridos"), http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/pendings?error="+url.QueryEscape("ID inválido"), http.StatusSeeOther)
		return
	}

	err = config.DB.Where("id = ? AND nombre_material = ?", id, materialName).Delete(&models.Pendings{}).Error
	if err != nil {
		errorMsg := err.Error()
		if strings.Contains(errorMsg, "foreign key constraint") {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("No se puede eliminar el pendiente porque tiene registros asociados"), http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Error al eliminar pendiente: "+errorMsg), http.StatusSeeOther)
		}
		return
	}

	http.Redirect(w, r, "/pendings?success="+url.QueryEscape("Pendiente eliminado exitosamente"), http.StatusSeeOther)
}

func AddPending(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		materialName := r.FormValue("material_name")
		departmentName := r.FormValue("department_name")
		quantity := r.FormValue("quantity")
		description := r.FormValue("description")
		dateStr := r.FormValue("date")
		requisition := r.FormValue("requisition")
		employeeName := r.FormValue("employee_name")

		if materialName == "" || departmentName == "" || quantity == "" || dateStr == "" || requisition == "" || employeeName == "" {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Faltan campos requeridos"), http.StatusSeeOther)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Cantidad inválida"), http.StatusSeeOther)
			return
		}

		// Parse date IN LOCAL TIMEZONE
		date, err := time.ParseInLocation("2006-01-02", dateStr, time.Local)
		if err != nil {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Fecha inválida"), http.StatusSeeOther)
			return
		}

		var parsedDescription *string
		if description != "" {
			parsedDescription = &description
		}

		output := models.Pendings{
			MaterialName:   materialName,
			DepartmentName: departmentName,
			Quantity:       quantityInt,
			Description:    parsedDescription,
			Date:           date,
			Requisition:    requisition,
			EmployeeName:   employeeName,
		}

		err = config.DB.Create(&output).Error
		if err != nil {
			errorMsg := err.Error()
			if strings.Contains(errorMsg, "foreign key constraint") || strings.Contains(errorMsg, "1452") {
				http.Redirect(w, r, "/pendings?error="+url.QueryEscape("El departamento o empleado especificado no existe. Por favor, verifique los datos."), http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Error al agregar pendiente: "+errorMsg), http.StatusSeeOther)
			}
			return
		}

		http.Redirect(w, r, "/pendings?success="+url.QueryEscape("Pendiente agregado exitosamente"), http.StatusSeeOther)
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

	if err := Templates.ExecuteTemplate(w, "add_pending", data); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func EditPending(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		materialName := r.FormValue("material_name")
		departmentName := r.FormValue("department_name")
		quantity := r.FormValue("quantity")
		description := r.FormValue("description")
		dateStr := r.FormValue("date")
		requisition := r.FormValue("requisition")
		employeeName := r.FormValue("employee_name")

		if idStr == "" || materialName == "" || departmentName == "" || quantity == "" || dateStr == "" || requisition == "" || employeeName == "" {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Faltan campos requeridos"), http.StatusSeeOther)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Cantidad inválida"), http.StatusSeeOther)
			return
		}

		// Parse date IN LOCAL TIMEZONE
		date, err := time.ParseInLocation("2006-01-02", dateStr, time.Local)
		if err != nil {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Fecha inválida"), http.StatusSeeOther)
			return
		}

		var parsedDescription *string
		if description != "" {
			parsedDescription = &description
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("ID inválido"), http.StatusSeeOther)
			return
		}

		err = config.DB.Model(&models.Pendings{}).Where("id = ?", id).Updates(models.Pendings{
			MaterialName:   materialName,
			DepartmentName: departmentName,
			Quantity:       quantityInt,
			Description:    parsedDescription,
			Date:           date,
			Requisition:    requisition,
			EmployeeName:   employeeName,
		}).Error
		if err != nil {
			errorMsg := err.Error()
			if strings.Contains(errorMsg, "foreign key constraint") || strings.Contains(errorMsg, "1452") {
				http.Redirect(w, r, "/pendings?error="+url.QueryEscape("El departamento o empleado especificado no existe. Por favor, verifique los datos."), http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Error al actualizar pendiente: "+errorMsg), http.StatusSeeOther)
			}
			return
		}

		http.Redirect(w, r, "/pendings?success="+url.QueryEscape("Pendiente actualizado exitosamente"), http.StatusSeeOther)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Falta el ID de pendiente", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de pendiente inválido", http.StatusBadRequest)
		return
	}

	var output models.Pendings
	err = config.DB.First(&output, id).Error
	if err != nil {
		http.Error(w, "Pendiente no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	// Fetch departments and employees for dropdowns
	var departments []models.Department
	var employees []models.Employee
	config.DB.Find(&departments)
	config.DB.Find(&employees)

	data := struct {
		Output      models.Pendings
		Departments []models.Department
		Employees   []models.Employee
	}{
		Output:      output,
		Departments: departments,
		Employees:   employees,
	}

	err = Templates.ExecuteTemplate(w, "edit_pending", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla: "+err.Error(), http.StatusInternalServerError)
	}
}
