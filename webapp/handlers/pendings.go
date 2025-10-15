package handlers

import (
	"net/http"
	"strconv"
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
	}{
		Title:    "Materiales Pendientes de Requisición",
		Pendings: pendings,
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
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		materialName := r.PostFormValue("material_name")
		employeeName := r.PostFormValue("employee_name")
		departmentName := r.PostFormValue("department_name")
		dateStr := r.PostFormValue("date")

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
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
			http.Error(w, "Error saving pending material: "+result.Error.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/pendings", http.StatusSeeOther)
		return
	}

	err := Templates.ExecuteTemplate(w, "add_pending", nil)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func DeletePending(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	materialName := r.URL.Query().Get("material_name")

	if idStr == "" || materialName == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	err = config.DB.Where("id = ? AND nombre_material = ?", id, materialName).Delete(&models.Pendings{}).Error
	if err != nil {
		http.Error(w, "Error deleting pending material: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pendings", http.StatusSeeOther)
}

func EditPending(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		materialName := r.FormValue("material_name")
		employeeName := r.FormValue("employee_name")
		departmentName := r.FormValue("department_name")
		dateStr := r.FormValue("date")

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}

		if idStr == "" || materialName == "" || employeeName == "" || departmentName == "" || dateStr == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		err = config.DB.Model(&models.Pendings{}).Where("id = ?", id).Updates(models.Pendings{
			MaterialName:   materialName,
			EmployeeName:   employeeName,
			DepartmentName: departmentName,
			Date:           date,
		}).Error
		if err != nil {
			http.Error(w, "Error updating pending material: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/pendings", http.StatusSeeOther)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing required parameter: id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	var pending models.Pendings
	err = config.DB.First(&pending, id).Error
	if err != nil {
		http.Error(w, "Pending material not found: "+err.Error(), http.StatusNotFound)
		return
	}

	err = Templates.ExecuteTemplate(w, "edit_pending", pending)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
