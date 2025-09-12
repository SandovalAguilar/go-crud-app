package handlers

import (
	"net/http"
	"webapp/config"
	"webapp/models"
)

func InventoryHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all inventory items from the database
	var inventoryItems []models.Inventory
	err := config.DB.Find(&inventoryItems).Error
	if err != nil {
		http.Error(w, "Error fetching inventory items: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the inventory template with the fetched data
	data := struct {
		Title string
		Items []models.Inventory
	}{
		Title: "Inventory Management",
		Items: inventoryItems,
	}

	// Execute the "general" template
	err = Templates.ExecuteTemplate(w, "general", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}
