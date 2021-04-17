package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler stores a pointer to mux dependency
type Handler struct {
	Router *mux.Router
}

// NewHandler returns a pointer to the Handler struct
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets up all routes in the application
func (h *Handler) SetupRoutes() {
	// print message to console
	fmt.Println("Setting up routes....")

	// initialize new mux router
	h.Router = mux.NewRouter()

	// define handler for router
	h.Router.HandleFunc("/api/health", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "All systems green!")
	})
}
