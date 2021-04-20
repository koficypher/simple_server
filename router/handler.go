package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koficypher/simple_server/controllers"
	"github.com/koficypher/simple_server/util"
	"gorm.io/gorm"
)

// controllers definition
// var articleController controllers.ArticleHandler

// Handler stores a pointer to mux dependency
type Handler struct {
	Router *mux.Router
}

// NewHandler returns a pointer to the Handler struct
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets up all routes in the application
func (h *Handler) SetupRoutes(db *gorm.DB) {
	// print message to console
	fmt.Println("Setting up routes....")

	// initialize new mux router
	h.Router = mux.NewRouter()

	//initialize controllers
	articleController := controllers.NewArticleController(db)

	// define handler for health check endpoin
	h.Router.HandleFunc("/api/health", func(res http.ResponseWriter, req *http.Request) {
		util.JSON(res, http.StatusOK, map[string]string{"message": "All services green...."})
	})

	h.Router.HandleFunc("/api/articles", articleController.GetAllArticles).Methods("GET")
}
