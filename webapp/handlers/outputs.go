package handlers

import (
	"net/http"
	"strconv"
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
	}{
		Title:   "Inventario de Salidas",
		Outputs: outputs,
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
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	err = config.DB.Where("id = ? AND nombre_material = ?", id, materialName).Delete(&models.InventoryOutput{}).Error
	if err != nil {
		http.Error(w, "Error deleting output: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/outputs", http.StatusSeeOther)
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
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
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
			http.Error(w, "Error adding output: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/outputs", http.StatusSeeOther)
		return
	}

	if err := Templates.ExecuteTemplate(w, "add_output", nil); err != nil {
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
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		var parsedDescription *string
		if description != "" {
			parsedDescription = &description
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
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
			http.Error(w, "Error updating output: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/outputs", http.StatusSeeOther)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing entry ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid entry ID", http.StatusBadRequest)
		return
	}

	var output models.InventoryOutput
	err = config.DB.First(&output, id).Error
	if err != nil {
		http.Error(w, "Output not found: "+err.Error(), http.StatusNotFound)
		return
	}

	err = Templates.ExecuteTemplate(w, "edit_output", output)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}
