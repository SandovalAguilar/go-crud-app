package handlers

import (
	"net/http"
	"webapp/config"
	"webapp/models"
)

func InventoryHandler(w http.ResponseWriter, r *http.Request) {
	var inventoryItems []models.Inventory
	err := config.DB.Find(&inventoryItems).Error
	if err != nil {
		http.Error(w, "Error fetching inventory items: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Items []models.Inventory
	}{
		Title: "Inventory Management",
		Items: inventoryItems,
	}

	err = Templates.ExecuteTemplate(w, "general", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}
