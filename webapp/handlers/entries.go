package handlers

import (
	"bytes"
	"log"
	"net/http"
)

// EntriesHandler serves the "entries" tab
func EntriesHandler(w http.ResponseWriter, r *http.Request) {
	// Example data structure
	data := struct {
		Items []struct {
			ID           int
			MaterialName string
			Description  string
			Quantity     int
			EntryDate    string
			SupplierName string
			Note         string
			Department   struct{ DepartmentName string }
		}
		TotalQuantityEntered int
	}{
		Items: []struct {
			ID           int
			MaterialName string
			Description  string
			Quantity     int
			EntryDate    string
			SupplierName string
			Note         string
			Department   struct{ DepartmentName string }
		}{
			{1, "Material A", "Desc A", 10, "2025-09-08", "Supplier X", "Note A", struct{ DepartmentName string }{"Dept 1"}},
			{2, "Material B", "Desc B", 5, "2025-09-08", "Supplier Y", "Note B", struct{ DepartmentName string }{"Dept 2"}},
		},
		TotalQuantityEntered: 15,
	}

	// Use a buffer to catch errors before writing to ResponseWriter
	var buf bytes.Buffer
	err := Templates.ExecuteTemplate(&buf, "entries", data)
	if err != nil {
		log.Println("Error rendering template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Only write to the client if template executed successfully
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to response:", err)
	}
}
