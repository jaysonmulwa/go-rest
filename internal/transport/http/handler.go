package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer tou pur comments service
type Handler struct {
	Router *mux.Router
}

// NewHandler - Return pointer to our handler struct
//We can cal this fro our main.go
func NewHandler() *Handler {
	return &Handler{}
}

//func seutp routes takes in a pointer to our handler
func (h *Handler) SetupRoutes() {

	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm alive")
	})

	/*r.HandleFunc("/comments", h.GetComments).Methods("GET")
	r.HandleFunc("/comments", h.AddComment).Methods("POST")
	r.HandleFunc("/comments/{id}", h.GetComment).Methods("GET")
	r.HandleFunc("/comments/{id}", h.UpdateComment).Methods("PUT")
	r.HandleFunc("/comments/{id}", h.DeleteComment).Methods("DELETE")*/
}
