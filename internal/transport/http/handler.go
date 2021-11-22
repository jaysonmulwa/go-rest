package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jaysonmulwa/go-rest/internal/comment"
)

// Handler - stores pointer tou pur comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

//Response - an object to store responses from our API.
type Response struct {
	Message string
}

// NewHandler - Return pointer to our handler struct
//We can cal this fro our main.go
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

//func seutp routes takes in a pointer to our handler
func (h *Handler) SetupRoutes() {

	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "I'm alive")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I'm alive"}); err != nil {
			panic(err)
		}
	})

}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}

}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {

		fmt.Fprintf(w, "Failed to decode JSON body")
	}

	comment, err := h.Service.PostComment(comment)

	if err != nil {
		fmt.Fprintf(w, "failed to post ew comment")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {

		fmt.Fprintf(w, "Failed to decode JSON body")
	}

	comment, err := h.Service.UpdateComment(1, comment)

	if err != nil {
		fmt.Fprintf(w, "failed to update comment")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Println(w, "failed to delete comment")
	}

	if err := json.NewEncoder(w).Encode("Comment deleted"); err != nil {
		panic(err)
	}
}
