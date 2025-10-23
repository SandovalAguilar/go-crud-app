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
	var pendings []models.Pendings
	err := config.DB.Find(&pendings).Error
	if err != nil {
		http.Error(w, "Error fetching pending materials: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title    string
		Pendings []models.Pendings
		Error    string
		Success  string
	}{
		Title:    "Materiales Pendientes de Requisición",
		Pendings: pendings,
		Error:    r.URL.Query().Get("error"),
		Success:  r.URL.Query().Get("success"),
	}

	err = Templates.ExecuteTemplate(w, "pendings", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func AddPending(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Error al procesar el formulario"), http.StatusSeeOther)
			return
		}

		materialName := r.PostFormValue("material_name")
		employeeName := r.PostFormValue("employee_name")
		departmentName := r.PostFormValue("department_name")
		dateStr := r.PostFormValue("date")

		if materialName == "" || employeeName == "" || departmentName == "" || dateStr == "" {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Todos los campos son requeridos"), http.StatusSeeOther)
			return
		}

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Formato de fecha inválido"), http.StatusSeeOther)
			return
		}

		pending := models.Pendings{
			MaterialName:   materialName,
			EmployeeName:   employeeName,
			DepartmentName: departmentName,
			Date:           date,
		}

		result := config.DB.Create(&pending)
		if result.Error != nil {
			errorMsg := result.Error.Error()
			if strings.Contains(errorMsg, "foreign key constraint") || strings.Contains(errorMsg, "1452") {
				http.Redirect(w, r, "/pendings?error="+url.QueryEscape("El empleado o departamento especificado no existe. Por favor, verifique los datos."), http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Error al guardar material pendiente: "+errorMsg), http.StatusSeeOther)
			}
			return
		}

		http.Redirect(w, r, "/pendings?success="+url.QueryEscape("Material pendiente agregado exitosamente"), http.StatusSeeOther)
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

	err := Templates.ExecuteTemplate(w, "add_pending", data)
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
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("No se puede eliminar el material pendiente porque tiene registros asociados"), http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Error al eliminar material pendiente: "+errorMsg), http.StatusSeeOther)
		}
		return
	}

	http.Redirect(w, r, "/pendings?success="+url.QueryEscape("Material pendiente eliminado exitosamente"), http.StatusSeeOther)
}

func EditPending(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		materialName := r.FormValue("material_name")
		employeeName := r.FormValue("employee_name")
		departmentName := r.FormValue("department_name")
		dateStr := r.FormValue("date")

		if idStr == "" || materialName == "" || employeeName == "" || departmentName == "" || dateStr == "" {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Todos los campos son requeridos"), http.StatusSeeOther)
			return
		}

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Formato de fecha inválido"), http.StatusSeeOther)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Redirect(w, r, "/pendings?error="+url.QueryEscape("ID inválido"), http.StatusSeeOther)
			return
		}

		err = config.DB.Model(&models.Pendings{}).Where("id = ?", id).Updates(models.Pendings{
			MaterialName:   materialName,
			EmployeeName:   employeeName,
			DepartmentName: departmentName,
			Date:           date,
		}).Error
		if err != nil {
			errorMsg := err.Error()
			if strings.Contains(errorMsg, "foreign key constraint") || strings.Contains(errorMsg, "1452") {
				http.Redirect(w, r, "/pendings?error="+url.QueryEscape("El empleado o departamento especificado no existe. Por favor, verifique los datos."), http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/pendings?error="+url.QueryEscape("Error al actualizar material pendiente: "+errorMsg), http.StatusSeeOther)
			}
			return
		}

		http.Redirect(w, r, "/pendings?success="+url.QueryEscape("Material pendiente actualizado exitosamente"), http.StatusSeeOther)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Falta el parámetro requerido: id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var pending models.Pendings
	err = config.DB.First(&pending, id).Error
	if err != nil {
		http.Error(w, "Material pendiente no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	// Fetch departments and employees for dropdowns
	var departments []models.Department
	var employees []models.Employee
	config.DB.Find(&departments)
	config.DB.Find(&employees)

	data := struct {
		Pending     models.Pendings
		Departments []models.Department
		Employees   []models.Employee
	}{
		Pending:     pending,
		Departments: departments,
		Employees:   employees,
	}

	err = Templates.ExecuteTemplate(w, "edit_pending", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
