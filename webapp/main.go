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
	http.HandleFunc("/entries", handlers.EntriesHandler)
	http.HandleFunc("/employees", handlers.ShowEmployees)
	http.HandleFunc("/employees/add", handlers.AddEmployee)
	http.HandleFunc("/employees/delete", handlers.DeleteEmployeeByName)
	http.HandleFunc("/employees/edit", handlers.EditEmployee)               // To render the edit form
	http.HandleFunc("/departments", handlers.ShowDepartments)               // Show all departments
	http.HandleFunc("/departments/add", handlers.AddDepartment)             // Add a new department
	http.HandleFunc("/departments/edit", handlers.EditDepartment)           // Edit a department
	http.HandleFunc("/departments/delete", handlers.DeleteDepartmentByName) // Delete a department
	http.HandleFunc("/orders", handlers.ShowOrders)
	http.HandleFunc("/orders/add", handlers.AddOrder)
	http.HandleFunc("/orders/delete", handlers.DeleteOrder)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
