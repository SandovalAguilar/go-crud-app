package handlers

import (
	"net/http"
	"strconv"
	"webapp/config"
	"webapp/models"
)

func ShowOrders(w http.ResponseWriter, r *http.Request) {
	var orderItems []models.Order
	err := config.DB.Find(&orderItems).Error
	if err != nil {
		http.Error(w, "Error fetching order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Items []models.Order
	}{
		Title: "Pedidos",
		Items: orderItems,
	}

	err = Templates.ExecuteTemplate(w, "orders", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		materialName := r.FormValue("material_name")
		supplierName := r.FormValue("supplier_name")
		materialDescription := r.FormValue("material_description")
		materialQuantity := r.FormValue("material_quantity")
		status := r.FormValue("status")
		requestDate := r.FormValue("request_date")
		deliveryDate := r.FormValue("delivery_date")
		note := r.FormValue("note")

		if materialName == "" || supplierName == "" || materialQuantity == "" || status == "" || requestDate == "" {
			http.Error(w, "All fields except 'Descripción del material' and 'Nota' are required", http.StatusBadRequest)
			return
		}

		// Convert materialQuantity to an integer
		quantity, err := strconv.Atoi(materialQuantity)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		var parsedDeliveryDate *string
		if deliveryDate != "" {
			parsedDeliveryDate = &deliveryDate
		}

		var parsedNote *string
		if note != "" {
			parsedNote = &note
		}

		var parsedMaterialDescription *string
		if materialDescription != "" {
			parsedMaterialDescription = &materialDescription
		}

		order := models.Order{
			MaterialName:        materialName,
			SupplierName:        supplierName,
			MaterialDescription: parsedMaterialDescription,
			MaterialQuantity:    quantity,
			Status:              status,
			RequestDate:         requestDate,
			DeliveryDate:        parsedDeliveryDate,
			Note:                parsedNote,
		}

		err = config.DB.Create(&order).Error
		if err != nil {
			http.Error(w, "Error adding order: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/orders", http.StatusSeeOther)
		return
	}

	Templates.ExecuteTemplate(w, "add_order", nil)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	materialName := r.URL.Query().Get("material_name")
	supplierName := r.URL.Query().Get("supplier_name")

	if idStr == "" || materialName == "" || supplierName == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	err = config.DB.Where("id = ? AND nombre_material = ? AND nombre_proveedor = ?", id, materialName, supplierName).Delete(&models.Order{}).Error
	if err != nil {
		http.Error(w, "Error deleting order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/orders", http.StatusSeeOther)
}

func EditOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		materialName := r.FormValue("material_name")
		supplierName := r.FormValue("supplier_name")
		materialDescription := r.FormValue("material_description")
		materialQuantity := r.FormValue("material_quantity")
		status := r.FormValue("status")
		requestDate := r.FormValue("request_date")
		deliveryDate := r.FormValue("delivery_date")
		note := r.FormValue("note")

		if idStr == "" || materialName == "" || supplierName == "" || materialQuantity == "" || status == "" || requestDate == "" {
			http.Error(w, "All required fields must be filled", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		quantity, err := strconv.Atoi(materialQuantity)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		var parsedMaterialDescription *string
		if materialDescription != "" {
			parsedMaterialDescription = &materialDescription
		}

		var parsedNote *string
		if note != "" {
			parsedNote = &note
		}

		var parsedDeliveryDate *string
		if deliveryDate != "" {
			parsedDeliveryDate = &deliveryDate
		}

		err = config.DB.Model(&models.Order{}).Where("id = ?", id).Updates(models.Order{
			MaterialName:        materialName,
			SupplierName:        supplierName,
			MaterialDescription: parsedMaterialDescription,
			MaterialQuantity:    quantity,
			Status:              status,
			RequestDate:         requestDate,
			DeliveryDate:        parsedDeliveryDate,
			Note:                parsedNote,
		}).Error
		if err != nil {
			http.Error(w, "Error updating order: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/orders", http.StatusSeeOther)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing order ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var order models.Order
	err = config.DB.First(&order, id).Error
	if err != nil {
		http.Error(w, "Order not found: "+err.Error(), http.StatusNotFound)
		return
	}

	err = Templates.ExecuteTemplate(w, "edit_order", order)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
