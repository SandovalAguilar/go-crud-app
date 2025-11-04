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
		Error   string
		Success string
	}{
		Title:   "Inventario de Entradas",
		Entries: inventoryEntries,
		Error:   r.URL.Query().Get("error"),
		Success: r.URL.Query().Get("success"),
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
		entryDateStr := r.FormValue("entry_date")

		if materialName == "" || quantity == "" || supplier == "" || entryDateStr == "" {
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("Faltan campos requeridos"), http.StatusSeeOther)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("Cantidad inválida"), http.StatusSeeOther)
			return
		}

		// Parse entry date IN LOCAL TIMEZONE
		entryDate, err := time.ParseInLocation("2006-01-02", entryDateStr, time.Local)
		if err != nil {
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("Fecha de entrada inválida"), http.StatusSeeOther)
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
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("Error al crear entrada: "+err.Error()), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/entries?success="+url.QueryEscape("Entrada agregada exitosamente"), http.StatusSeeOther)
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
		http.Redirect(w, r, "/entries?error="+url.QueryEscape("Faltan parámetros requeridos"), http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/entries?error="+url.QueryEscape("ID inválido"), http.StatusSeeOther)
		return
	}

	err = config.DB.Where("id = ? AND nombre_material = ?", id, materialName).Delete(&models.InventoryEntry{}).Error
	if err != nil {
		errorMsg := err.Error()
		if strings.Contains(errorMsg, "foreign key constraint") {
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("No se puede eliminar la entrada porque tiene registros asociados"), http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("Error al eliminar entrada: "+errorMsg), http.StatusSeeOther)
		}
		return
	}

	http.Redirect(w, r, "/entries?success="+url.QueryEscape("Entrada eliminada exitosamente"), http.StatusSeeOther)
}

func EditEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		materialName := r.FormValue("material_name")
		supplierName := r.FormValue("supplier_name")
		materialDescription := r.FormValue("material_description")
		materialQuantity := r.FormValue("material_quantity")
		entryDateStr := r.FormValue("entry_date")
		note := r.FormValue("note")

		if idStr == "" || materialName == "" || supplierName == "" || materialQuantity == "" || entryDateStr == "" {
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("Todos los campos requeridos deben estar llenos"), http.StatusSeeOther)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("ID de entrada inválido"), http.StatusSeeOther)
			return
		}

		quantity, err := strconv.Atoi(materialQuantity)
		if err != nil {
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("Cantidad inválida"), http.StatusSeeOther)
			return
		}

		// Parse entry date IN LOCAL TIMEZONE
		entryDate, err := time.ParseInLocation("2006-01-02", entryDateStr, time.Local)
		if err != nil {
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("Fecha de entrada inválida"), http.StatusSeeOther)
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
			http.Redirect(w, r, "/entries?error="+url.QueryEscape("Error al actualizar entrada: "+err.Error()), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/entries?success="+url.QueryEscape("Entrada actualizada exitosamente"), http.StatusSeeOther)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Falta el ID de entrada", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de entrada inválido", http.StatusBadRequest)
		return
	}

	var entry models.InventoryEntry
	err = config.DB.First(&entry, id).Error
	if err != nil {
		http.Error(w, "Entrada no encontrada: "+err.Error(), http.StatusNotFound)
		return
	}

	err = Templates.ExecuteTemplate(w, "edit_entry", entry)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla: "+err.Error(), http.StatusInternalServerError)
	}
}
