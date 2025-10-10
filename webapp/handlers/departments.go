package handlers

import (
	"net/http"
	"regexp"
	"strconv"
	"webapp/config"
	"webapp/models"
)

func ShowDepartments(w http.ResponseWriter, r *http.Request) {
	var departmentItems []models.Department
	err := config.DB.Find(&departmentItems).Error // Pass a pointer to the slice
	if err != nil {
		http.Error(w, "Error fetching departments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Items []models.Department
	}{
		Title: "Departamentos",
		Items: departmentItems,
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
			http.Error(w, "Department name must contain only alphanumeric characters", http.StatusBadRequest)
			return
		}

		department := models.Department{
			DepartmentName: departmentName,
		}

		err := config.DB.Create(&department).Error
		if err != nil {
			http.Error(w, "Error adding department: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/departments", http.StatusSeeOther)
		return
	}

	Templates.ExecuteTemplate(w, "add_department", nil)
}

func DeleteDepartmentByName(w http.ResponseWriter, r *http.Request) {
	departmentName := r.URL.Query().Get("department_name")

	if departmentName == "" {
		http.Error(w, "Department name must be provided", http.StatusBadRequest)
		return
	}

	var department models.Department
	err := config.DB.Where("nombre_departamento = ?", departmentName).First(&department).Error
	if err != nil {
		http.Error(w, "Department not found: "+err.Error(), http.StatusNotFound)
		return
	}

	err = config.DB.Delete(&department).Error
	if err != nil {
		http.Error(w, "Error deleting department: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/departments", http.StatusSeeOther)
}

func EditDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Department ID not provided", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var department models.Department
		err = config.DB.First(&department, id).Error
		if err != nil {
			http.Error(w, "Department not found: "+err.Error(), http.StatusNotFound)
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
			http.Error(w, "Department ID not provided", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		departmentName := r.FormValue("department_name")

		var department models.Department
		err = config.DB.First(&department, id).Error
		if err != nil {
			http.Error(w, "Department not found: "+err.Error(), http.StatusNotFound)
			return
		}

		department.DepartmentName = departmentName

		err = config.DB.Save(&department).Error
		if err != nil {
			http.Error(w, "Error updating department: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/departments", http.StatusSeeOther)
		return
	}
}
