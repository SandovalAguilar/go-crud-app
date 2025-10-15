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
		materialName := r.FormValue("material_name")
		quantity := r.FormValue("material_quantity")
		description := r.FormValue("material_description")
		supplier := r.FormValue("supplier_name")
		note := r.FormValue("note")
		entryDate := r.FormValue("entry_date")

		if materialName == "" || quantity == "" || supplier == "" || entryDate == "" {
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

		var parsedNote *string
		if note != "" {
			parsedNote = &note
		}

		entry := models.InventoryEntry{
			MaterialName:        materialName,
			Quantity:            quantityInt,
			MaterialDescription: parsedDescription,
			SupplierName:        supplier,
			Note:                parsedNote,
			EntryDate:           entryDate,
		}

		err = config.DB.Create(&entry).Error
		if err != nil {
			http.Error(w, "Error creating entry: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/entries", http.StatusSeeOther)
		return
	}

	if err := Templates.ExecuteTemplate(w, "add_entry", nil); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func DeleteEntry(w http.ResponseWriter, r *http.Request) {
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

	err = config.DB.Where("id = ? AND nombre_material = ?", id, materialName).Delete(&models.InventoryEntry{}).Error
	if err != nil {
		http.Error(w, "Error deleting entry: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/entries", http.StatusSeeOther)
}

func EditEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		materialName := r.FormValue("material_name")
		supplierName := r.FormValue("supplier_name")
		materialDescription := r.FormValue("material_description")
		materialQuantity := r.FormValue("material_quantity")
		entryDate := r.FormValue("entry_date")
		note := r.FormValue("note")

		if idStr == "" || materialName == "" || supplierName == "" || materialQuantity == "" || entryDate == "" {
			http.Error(w, "All required fields must be filled", http.StatusBadRequest)
			return
		}

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

		var parsedDescription *string
		if materialDescription != "" {
			parsedDescription = &materialDescription
		}

		var parsedNote *string
		if note != "" {
			parsedNote = &note
		}

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

		http.Redirect(w, r, "/entries", http.StatusSeeOther)
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

	var entry models.InventoryEntry
	err = config.DB.First(&entry, id).Error
	if err != nil {
		http.Error(w, "Entry not found: "+err.Error(), http.StatusNotFound)
		return
	}

	err = Templates.ExecuteTemplate(w, "edit_entry", entry)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}
