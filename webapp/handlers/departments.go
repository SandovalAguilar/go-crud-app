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
	// Handle POST request to add a new department
	if r.Method == http.MethodPost {
		// Extract department data from the form
		departmentName := r.FormValue("department_name")

		// Verify that departmentName contains only alphanumeric characters
		isAlnum := regexp.MustCompile(`^[a-zA-Z0-9 ]+$`).MatchString
		if !isAlnum(departmentName) {
			http.Error(w, "Department name must contain only alphanumeric characters", http.StatusBadRequest)
			return
		}

		// Create a new department record
		department := models.Department{
			DepartmentName: departmentName,
		}

		// Insert the department into the database
		err := config.DB.Create(&department).Error
		if err != nil {
			http.Error(w, "Error adding department: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the department list page after adding the department
		http.Redirect(w, r, "/departments", http.StatusSeeOther)
		return
	}

	// Handle GET request to render the Add Department form
	Templates.ExecuteTemplate(w, "add_department", nil)
}

// DeleteDepartmentByName deletes a department by its name
func DeleteDepartmentByName(w http.ResponseWriter, r *http.Request) {
	departmentName := r.URL.Query().Get("department_name")

	// Validate if department_name is provided
	if departmentName == "" {
		http.Error(w, "Department name must be provided", http.StatusBadRequest)
		return
	}

	// Find the department by name
	var department models.Department
	err := config.DB.Where("nombre_departamento = ?", departmentName).First(&department).Error
	if err != nil {
		http.Error(w, "Department not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// Delete the department from the database
	err = config.DB.Delete(&department).Error
	if err != nil {
		http.Error(w, "Error deleting department: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the department list page after successful deletion
	http.Redirect(w, r, "/departments", http.StatusSeeOther)
}

// EditDepartment handles both displaying the edit form (GET) and updating the department (POST)
func EditDepartment(w http.ResponseWriter, r *http.Request) {
	// Handle GET request to display the edit form
	if r.Method == http.MethodGet {
		// Get the department ID from the URL
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Department ID not provided", http.StatusBadRequest)
			return
		}

		// Convert the department ID from string to integer
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Fetch the department from the database
		var department models.Department
		err = config.DB.First(&department, id).Error
		if err != nil {
			http.Error(w, "Department not found: "+err.Error(), http.StatusNotFound)
			return
		}

		// Render the edit form template with the department data
		data := struct {
			Department models.Department
		}{
			Department: department,
		}

		Templates.ExecuteTemplate(w, "edit_department", data)
		return
	}

	// Handle POST request to update the department data
	if r.Method == http.MethodPost {
		// Get the department ID from the URL
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Department ID not provided", http.StatusBadRequest)
			return
		}

		// Convert the department ID from string to integer
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Fetch the updated data from the form
		departmentName := r.FormValue("department_name")

		// Find the department from the database
		var department models.Department
		err = config.DB.First(&department, id).Error
		if err != nil {
			http.Error(w, "Department not found: "+err.Error(), http.StatusNotFound)
			return
		}

		// Update the department data (Only Department Name)
		department.DepartmentName = departmentName

		// Save the updated department data to the database
		err = config.DB.Save(&department).Error
		if err != nil {
			http.Error(w, "Error updating department: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect back to the department list after successful update
		http.Redirect(w, r, "/departments", http.StatusSeeOther)
		return
	}
}
