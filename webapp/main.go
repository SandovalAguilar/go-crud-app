package main

import (
	"log"
	"net/http"
	"time"
	"webapp/config"
	"webapp/handlers"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the incoming request
		log.Printf("[%s] %s %s - Started", r.Method, r.URL.Path, r.RemoteAddr)

		// Create a custom ResponseWriter to capture the status code
		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Call the next handler
		next.ServeHTTP(lrw, r)

		// Log the response
		duration := time.Since(start)
		log.Printf("[%s] %s %s - Completed %d in %v",
			r.Method, r.URL.Path, r.RemoteAddr, lrw.statusCode, duration)
	}
}

// loggingResponseWriter wraps http.ResponseWriter to capture status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func main() {
	log.Println("Initializing database connection...")
	config.InitDB()
	log.Println("Database connection established successfully")

	log.Println("Registering routes...")

	http.HandleFunc("/", LoggingMiddleware(handlers.InventoryHandler))

	// Entries routes
	http.HandleFunc("/entries", LoggingMiddleware(handlers.ShowEntries))
	http.HandleFunc("/entries/add", LoggingMiddleware(handlers.AddEntry))
	http.HandleFunc("/entries/delete", LoggingMiddleware(handlers.DeleteEntry))
	http.HandleFunc("/entries/edit", LoggingMiddleware(handlers.EditEntry))

	// Employees routes
	http.HandleFunc("/employees", LoggingMiddleware(handlers.ShowEmployees))
	http.HandleFunc("/employees/add", LoggingMiddleware(handlers.AddEmployee))
	http.HandleFunc("/employees/delete", LoggingMiddleware(handlers.DeleteEmployeeByName))
	http.HandleFunc("/employees/edit", LoggingMiddleware(handlers.EditEmployee))

	// Departments routes
	http.HandleFunc("/departments", LoggingMiddleware(handlers.ShowDepartments))
	http.HandleFunc("/departments/add", LoggingMiddleware(handlers.AddDepartment))
	http.HandleFunc("/departments/edit", LoggingMiddleware(handlers.EditDepartment))
	http.HandleFunc("/departments/delete", LoggingMiddleware(handlers.DeleteDepartmentByName))

	// Orders routes
	http.HandleFunc("/orders", LoggingMiddleware(handlers.ShowOrders))
	http.HandleFunc("/orders/add", LoggingMiddleware(handlers.AddOrder))
	http.HandleFunc("/orders/delete", LoggingMiddleware(handlers.DeleteOrder))
	http.HandleFunc("/orders/edit", LoggingMiddleware(handlers.EditOrder))

	// Outputs routes
	http.HandleFunc("/outputs", LoggingMiddleware(handlers.ShowOutputs))
	http.HandleFunc("/outputs/delete", LoggingMiddleware(handlers.DeleteOutput))
	http.HandleFunc("/outputs/add", LoggingMiddleware(handlers.AddOutput))
	http.HandleFunc("/outputs/edit", LoggingMiddleware(handlers.EditOutput))

	// Pendings routes
	http.HandleFunc("/pendings", LoggingMiddleware(handlers.ShowPendings))
	http.HandleFunc("/pendings/add", LoggingMiddleware(handlers.AddPending))
	http.HandleFunc("/pendings/delete", LoggingMiddleware(handlers.DeletePending))
	http.HandleFunc("/pendings/edit", LoggingMiddleware(handlers.EditPending))

	// Static files (no logging middleware for performance)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Routes registered successfully")
	log.Println("===========================================")
	log.Println("Server starting on http://localhost:8080")
	log.Println("===========================================")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
