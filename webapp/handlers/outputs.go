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
	// Get query parameters
	idStr := r.URL.Query().Get("id")
	materialName := r.URL.Query().Get("material_name")

	// Validate parameters
	if idStr == "" || materialName == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	// Convert id to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	// Delete the output from the database using both ID and material name
	err = config.DB.Where("id = ? AND nombre_material = ?", id, materialName).Delete(&models.InventoryOutput{}).Error
	if err != nil {
		http.Error(w, "Error deleting output: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the outputs list page
	http.Redirect(w, r, "/outputs", http.StatusSeeOther)
}

func AddOutput(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle form submission to create a new output
		materialName := r.FormValue("material_name")
		departmentName := r.FormValue("department_name")
		quantity := r.FormValue("quantity")
		description := r.FormValue("description")
		date := r.FormValue("date")
		delivered := r.FormValue("delivered")
		employeeName := r.FormValue("employee_name")

		// Validate required fields
		if materialName == "" || departmentName == "" || quantity == "" || date == "" || delivered == "" || employeeName == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		// Convert quantity to integer
		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		// Handle optional description as a pointer
		var parsedDescription *string
		if description != "" {
			parsedDescription = &description
		}

		// Create a new inventory output record
		output := models.InventoryOutput{
			MaterialName:   materialName,
			DepartmentName: departmentName,
			Quantity:       quantityInt,
			Description:    parsedDescription,
			Date:           date,
			Delivered:      delivered,
			EmployeeName:   employeeName,
		}

		// Insert the output into the database
		err = config.DB.Create(&output).Error
		if err != nil {
			http.Error(w, "Error adding output: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the outputs list page
		http.Redirect(w, r, "/outputs", http.StatusSeeOther)
		return
	}

	// Render the add output form
	if err := Templates.ExecuteTemplate(w, "add_output", nil); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}
