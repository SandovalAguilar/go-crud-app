package main

import (
	"log"
	"net/http"
	"webapp/config"
	"webapp/handlers"
)

func main() {
	config.InitDB()

	http.HandleFunc("/", handlers.InventoryHandler)
	http.HandleFunc("/entries", handlers.ShowEntries)
	http.HandleFunc("/entries/add", handlers.AddEntry)
	http.HandleFunc("/entries/delete", handlers.DeleteEntry)
	http.HandleFunc("/entries/edit", handlers.EditEntry)
	http.HandleFunc("/employees", handlers.ShowEmployees)
	http.HandleFunc("/employees/add", handlers.AddEmployee)
	http.HandleFunc("/employees/delete", handlers.DeleteEmployeeByName)
	http.HandleFunc("/employees/edit", handlers.EditEmployee)
	http.HandleFunc("/departments", handlers.ShowDepartments)
	http.HandleFunc("/departments/add", handlers.AddDepartment)
	http.HandleFunc("/departments/edit", handlers.EditDepartment)
	http.HandleFunc("/departments/delete", handlers.DeleteDepartmentByName)
	http.HandleFunc("/orders", handlers.ShowOrders)
	http.HandleFunc("/orders/add", handlers.AddOrder)
	http.HandleFunc("/orders/delete", handlers.DeleteOrder)
	http.HandleFunc("/orders/edit", handlers.EditOrder)
	http.HandleFunc("/outputs", handlers.ShowOutputs)
	http.HandleFunc("/outputs/delete", handlers.DeleteOutput)
	http.HandleFunc("/outputs/add", handlers.AddOutput)
	http.HandleFunc("/outputs/edit", handlers.EditOutput)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
