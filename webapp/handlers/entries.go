package handlers

import (
	"net/http"
	"strconv"
	"webapp/config"
	"webapp/models"
)

func ShowEntries(w http.ResponseWriter, r *http.Request) {
	var inventoryEntries []models.InventoryEntry
	err := config.DB.Find(&inventoryEntries).Error
	if err != nil {
		http.Error(w, "Error fetching inventory entries: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Entries []models.InventoryEntry
	}{
		Title:   "Inventario de Entradas",
		Entries: inventoryEntries,
	}

	err = Templates.ExecuteTemplate(w, "entries", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func AddEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle form submission to create a new entry
		materialName := r.FormValue("material_name")
		quantity := r.FormValue("material_quantity")
		description := r.FormValue("material_description")
		supplier := r.FormValue("supplier_name")
		note := r.FormValue("note")
		entryDate := r.FormValue("entry_date") // <-- added

		// Validate the required fields
		if materialName == "" || quantity == "" || supplier == "" || entryDate == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		// Convert quantity to integer
		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		// Handle optional fields as pointers
		var parsedDescription *string
		if description != "" {
			parsedDescription = &description
		}

		var parsedNote *string
		if note != "" {
			parsedNote = &note
		}

		// Create a new inventory entry including the entry date
		entry := models.InventoryEntry{
			MaterialName:        materialName,
			Quantity:            quantityInt,
			MaterialDescription: parsedDescription,
			SupplierName:        supplier,
			Note:                parsedNote,
			EntryDate:           entryDate, // <-- assign the date
		}

		// Save the new entry in the database
		err = config.DB.Create(&entry).Error
		if err != nil {
			http.Error(w, "Error creating entry: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the entries page after creating the new entry
		http.Redirect(w, r, "/entries", http.StatusSeeOther)
		return
	}

	// Render the create form
	if err := Templates.ExecuteTemplate(w, "add_entry", nil); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func DeleteEntry(w http.ResponseWriter, r *http.Request) {
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

	// Delete the entry from the database using both ID and material name
	err = config.DB.Where("id = ? AND nombre_material = ?", id, materialName).Delete(&models.InventoryEntry{}).Error
	if err != nil {
		http.Error(w, "Error deleting entry: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the entries list page
	http.Redirect(w, r, "/entries", http.StatusSeeOther)
}

func EditEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// POST: update the inventory entry
		idStr := r.FormValue("id")
		materialName := r.FormValue("material_name")
		supplierName := r.FormValue("supplier_name")
		materialDescription := r.FormValue("material_description")
		materialQuantity := r.FormValue("material_quantity")
		entryDate := r.FormValue("entry_date")
		note := r.FormValue("note")

		// Validate required fields
		if idStr == "" || materialName == "" || supplierName == "" || materialQuantity == "" || entryDate == "" {
			http.Error(w, "All required fields must be filled", http.StatusBadRequest)
			return
		}

		// Convert ID and quantity
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid entry ID", http.StatusBadRequest)
			return
		}

		quantity, err := strconv.Atoi(materialQuantity)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		// Handle optional fields as pointers
		var parsedDescription *string
		if materialDescription != "" {
			parsedDescription = &materialDescription
		}

		var parsedNote *string
		if note != "" {
			parsedNote = &note
		}

		// Update the entry in the database
		err = config.DB.Model(&models.InventoryEntry{}).Where("id = ?", id).Updates(models.InventoryEntry{
			MaterialName:        materialName,
			Quantity:            quantity,
			MaterialDescription: parsedDescription,
			SupplierName:        supplierName,
			EntryDate:           entryDate,
			Note:                parsedNote,
		}).Error
		if err != nil {
			http.Error(w, "Error updating entry: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the entries list page after updating
		http.Redirect(w, r, "/entries", http.StatusSeeOther)
		return
	}

	// GET: show the edit form with existing data
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

	// Fetch the entry from the database
	var entry models.InventoryEntry
	err = config.DB.First(&entry, id).Error
	if err != nil {
		http.Error(w, "Entry not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// Render the edit template
	err = Templates.ExecuteTemplate(w, "edit_entry", entry)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}
