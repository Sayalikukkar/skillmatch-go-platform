package routes

import (
	// "net/http"
	"taskify/handlers"
	"github.com/gorilla/mux"
)

// SetupRoutes initializes the routes for the application
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// User-related routes
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")            // Create User
	// router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")           // Create Task
	// router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")       // Update Task
	// router.HandleFunc("/offers/{id}/accept", handlers.AcceptOffer).Methods("PUT") // Accept Offer
	// router.HandleFunc("/offers/{id}/reject", handlers.RejectOffer).Methods("PUT") // Reject Offer
	// router.HandleFunc("/tasks/{id}/accept", handlers.AcceptTaskCompletion).Methods("PUT") // Accept Task Completion
	// router.HandleFunc("/tasks/{id}/reject", handlers.RejectTaskCompletion).Methods("PUT") // Reject Task Completion

	// // Provider-related routes
	router.HandleFunc("/providers", handlers.CreateProvider).Methods("POST")    // Create Provider
	// router.HandleFunc("/skills", handlers.CreateSkill).Methods("POST")          // Create Skill
	// router.HandleFunc("/skills/{id}", handlers.UpdateSkill).Methods("PUT")      // Update Skill
	// router.HandleFunc("/tasks/{id}/offer", handlers.MakeOffer).Methods("POST")  // Make an Offer
	// router.HandleFunc("/tasks/{id}/progress", handlers.UpdateTaskProgress).Methods("PUT") // Update Task Progress
	// router.HandleFunc("/tasks/{id}/complete", handlers.MarkTaskCompleted).Methods("PUT")  // Mark Task as Completed

	return router
}
