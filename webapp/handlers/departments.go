package handlers

import (
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"webapp/config"
	"webapp/models"
)

func ShowDepartments(w http.ResponseWriter, r *http.Request) {
	var departmentItems []models.Department
	err := config.DB.Find(&departmentItems).Error
	if err != nil {
		http.Error(w, "Error fetching departments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Items   []models.Department
		Error   string
		Success string
	}{
		Title:   "Departamentos",
		Items:   departmentItems,
		Error:   r.URL.Query().Get("error"),
		Success: r.URL.Query().Get("success"),
	}

	err = Templates.ExecuteTemplate(w, "departments", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func AddDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		departmentName := r.FormValue("department_name")
		isAlnum := regexp.MustCompile(`^[a-zA-Z0-9 ]+$`).MatchString
		if !isAlnum(departmentName) {
			http.Redirect(w, r, "/departments?error="+url.QueryEscape("El nombre del departamento debe contener solo caracteres alfanuméricos"), http.StatusSeeOther)
			return
		}

		department := models.Department{
			DepartmentName: departmentName,
		}

		err := config.DB.Create(&department).Error
		if err != nil {
			http.Redirect(w, r, "/departments?error="+url.QueryEscape("Error al agregar departamento: "+err.Error()), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/departments?success="+url.QueryEscape("Departamento agregado exitosamente"), http.StatusSeeOther)
		return
	}

	Templates.ExecuteTemplate(w, "add_department", nil)
}

func DeleteDepartmentByName(w http.ResponseWriter, r *http.Request) {
	departmentName := r.URL.Query().Get("department_name")

	if departmentName == "" {
		http.Redirect(w, r, "/departments?error="+url.QueryEscape("Debe proporcionar el nombre del departamento"), http.StatusSeeOther)
		return
	}

	var department models.Department
	err := config.DB.Where("nombre_departamento = ?", departmentName).First(&department).Error
	if err != nil {
		http.Redirect(w, r, "/departments?error="+url.QueryEscape("Departamento no encontrado"), http.StatusSeeOther)
		return
	}

	err = config.DB.Delete(&department).Error
	if err != nil {
		errorMsg := err.Error()
		if strings.Contains(errorMsg, "foreign key constraint") || strings.Contains(errorMsg, "1451") {
			http.Redirect(w, r, "/departments?error="+url.QueryEscape("No se puede eliminar el departamento porque tiene empleados o registros asociados"), http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/departments?error="+url.QueryEscape("Error al eliminar departamento: "+errorMsg), http.StatusSeeOther)
		}
		return
	}

	http.Redirect(w, r, "/departments?success="+url.QueryEscape("Departamento eliminado exitosamente"), http.StatusSeeOther)
}

func EditDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "ID del departamento no proporcionado", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		var department models.Department
		err = config.DB.First(&department, id).Error
		if err != nil {
			http.Error(w, "Departamento no encontrado: "+err.Error(), http.StatusNotFound)
			return
		}

		data := struct {
			Department models.Department
		}{
			Department: department,
		}
		Templates.ExecuteTemplate(w, "edit_department", data)
		return
	}

	if r.Method == http.MethodPost {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "ID del departamento no proporcionado", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		departmentName := r.FormValue("department_name")

		var department models.Department
		err = config.DB.First(&department, id).Error
		if err != nil {
			http.Redirect(w, r, "/departments?error="+url.QueryEscape("Departamento no encontrado"), http.StatusSeeOther)
			return
		}

		department.DepartmentName = departmentName
		err = config.DB.Save(&department).Error
		if err != nil {
			http.Redirect(w, r, "/departments?error="+url.QueryEscape("Error al actualizar departamento: "+err.Error()), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/departments?success="+url.QueryEscape("Departamento actualizado exitosamente"), http.StatusSeeOther)
		return
	}
}
