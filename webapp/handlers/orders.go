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

func ShowOrders(w http.ResponseWriter, r *http.Request) {
	var orderItems []models.Order
	err := config.DB.Find(&orderItems).Error
	if err != nil {
		http.Error(w, "Error fetching order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Items   []models.Order
		Error   string
		Success string
	}{
		Title:   "Pedidos",
		Items:   orderItems,
		Error:   r.URL.Query().Get("error"),
		Success: r.URL.Query().Get("success"),
	}

	err = Templates.ExecuteTemplate(w, "orders", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	materialName := r.URL.Query().Get("material_name")
	supplierName := r.URL.Query().Get("supplier_name")

	if idStr == "" || materialName == "" || supplierName == "" {
		http.Redirect(w, r, "/orders?error="+url.QueryEscape("Faltan parámetros requeridos"), http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/orders?error="+url.QueryEscape("ID inválido"), http.StatusSeeOther)
		return
	}

	err = config.DB.Where("id = ? AND nombre_material = ? AND nombre_proveedor = ?", id, materialName, supplierName).Delete(&models.Order{}).Error
	if err != nil {
		errorMsg := err.Error()
		if strings.Contains(errorMsg, "foreign key constraint") {
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("No se puede eliminar el pedido porque tiene registros asociados"), http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("Error al eliminar pedido: "+errorMsg), http.StatusSeeOther)
		}
		return
	}

	http.Redirect(w, r, "/orders?success="+url.QueryEscape("Pedido eliminado exitosamente"), http.StatusSeeOther)
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		materialName := r.FormValue("material_name")
		supplierName := r.FormValue("supplier_name")
		materialDescription := r.FormValue("material_description")
		materialQuantity := r.FormValue("material_quantity")
		status := r.FormValue("status")
		requestDateStr := r.FormValue("request_date")
		deliveryDateStr := r.FormValue("delivery_date")
		note := r.FormValue("note")

		if materialName == "" || supplierName == "" || materialQuantity == "" || status == "" || requestDateStr == "" {
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("Todos los campos excepto 'Descripción del material' y 'Nota' son requeridos"), http.StatusSeeOther)
			return
		}

		// Convert materialQuantity to an integer
		quantity, err := strconv.Atoi(materialQuantity)
		if err != nil {
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("Cantidad inválida"), http.StatusSeeOther)
			return
		}

		// Parse request date IN LOCAL TIMEZONE
		requestDate, err := time.ParseInLocation("2006-01-02", requestDateStr, time.Local)
		if err != nil {
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("Fecha de pedido inválida"), http.StatusSeeOther)
			return
		}

		// Parse delivery date (optional) IN LOCAL TIMEZONE
		var parsedDeliveryDate *time.Time
		if deliveryDateStr != "" {
			deliveryDate, err := time.ParseInLocation("2006-01-02", deliveryDateStr, time.Local)
			if err != nil {
				http.Redirect(w, r, "/orders?error="+url.QueryEscape("Fecha de entrega inválida"), http.StatusSeeOther)
				return
			}

			// Validate: request date should not be after delivery date
			if requestDate.After(deliveryDate) {
				http.Redirect(w, r, "/orders?error="+url.QueryEscape("La fecha de pedido no puede ser posterior a la fecha de entrega"), http.StatusSeeOther)
				return
			}

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
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("Error al agregar pedido: "+err.Error()), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/orders?success="+url.QueryEscape("Pedido agregado exitosamente"), http.StatusSeeOther)
		return
	}

	Templates.ExecuteTemplate(w, "add_order", nil)
}

func EditOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		materialName := r.FormValue("material_name")
		supplierName := r.FormValue("supplier_name")
		materialDescription := r.FormValue("material_description")
		materialQuantity := r.FormValue("material_quantity")
		status := r.FormValue("status")
		requestDateStr := r.FormValue("request_date")
		deliveryDateStr := r.FormValue("delivery_date")
		note := r.FormValue("note")

		if idStr == "" || materialName == "" || supplierName == "" || materialQuantity == "" || status == "" || requestDateStr == "" {
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("Todos los campos requeridos deben estar llenos"), http.StatusSeeOther)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("ID de pedido inválido"), http.StatusSeeOther)
			return
		}

		quantity, err := strconv.Atoi(materialQuantity)
		if err != nil {
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("Cantidad inválida"), http.StatusSeeOther)
			return
		}

		// Parse request date IN LOCAL TIMEZONE
		requestDate, err := time.ParseInLocation("2006-01-02", requestDateStr, time.Local)
		if err != nil {
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("Fecha de pedido inválida"), http.StatusSeeOther)
			return
		}

		// Parse delivery date (optional) IN LOCAL TIMEZONE
		var parsedDeliveryDate *time.Time
		if deliveryDateStr != "" {
			deliveryDate, err := time.ParseInLocation("2006-01-02", deliveryDateStr, time.Local)
			if err != nil {
				http.Redirect(w, r, "/orders?error="+url.QueryEscape("Fecha de entrega inválida"), http.StatusSeeOther)
				return
			}

			// Validate: request date should not be after delivery date
			if requestDate.After(deliveryDate) {
				http.Redirect(w, r, "/orders?error="+url.QueryEscape("La fecha de pedido no puede ser posterior a la fecha de entrega"), http.StatusSeeOther)
				return
			}

			parsedDeliveryDate = &deliveryDate
		}

		var parsedMaterialDescription *string
		if materialDescription != "" {
			parsedMaterialDescription = &materialDescription
		}

		var parsedNote *string
		if note != "" {
			parsedNote = &note
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
			http.Redirect(w, r, "/orders?error="+url.QueryEscape("Error al actualizar pedido: "+err.Error()), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/orders?success="+url.QueryEscape("Pedido actualizado exitosamente"), http.StatusSeeOther)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Falta el ID del pedido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de pedido inválido", http.StatusBadRequest)
		return
	}

	var order models.Order
	err = config.DB.First(&order, id).Error
	if err != nil {
		http.Error(w, "Pedido no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	err = Templates.ExecuteTemplate(w, "edit_order", order)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
