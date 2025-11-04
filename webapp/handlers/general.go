package handlers

import (
	"net/http"
	"strings"
	"webapp/config"
	"webapp/models"
)

func InventoryHandler(w http.ResponseWriter, r *http.Request) {
	var entradas []models.InventoryEntry
	err := config.DB.Find(&entradas).Error
	if err != nil {
		http.Error(w, "Error fetching entries: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var salidas []models.InventoryOutput
	err = config.DB.Find(&salidas).Error
	if err != nil {
		http.Error(w, "Error fetching outputs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Calculate balances
	balanceMap := make(map[string]*models.InventoryBalance)

	for _, entrada := range entradas {
		materialKey := strings.ToLower(strings.TrimSpace(entrada.MaterialName))

		if balanceMap[materialKey] == nil {
			balanceMap[materialKey] = &models.InventoryBalance{
				MaterialName: entrada.MaterialName, // Use original name for display
			}
		}
		balanceMap[materialKey].TotalEntries += entrada.Quantity
	}

	for _, salida := range salidas {
		materialKey := strings.ToLower(strings.TrimSpace(salida.MaterialName))

		if balanceMap[materialKey] == nil {
			balanceMap[materialKey] = &models.InventoryBalance{
				MaterialName: salida.MaterialName, // Use original name for display
			}
		}
		balanceMap[materialKey].TotalOutputs += salida.Quantity
	}

	var inventory []models.InventoryBalance
	for _, balance := range balanceMap {
		balance.AvailableStock = balance.TotalEntries - balance.TotalOutputs
		inventory = append(inventory, *balance)
	}

	data := struct {
		Title     string
		Inventory []models.InventoryBalance
	}{
		Title:     "Inventario General",
		Inventory: inventory,
	}

	err = Templates.ExecuteTemplate(w, "general", data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}
