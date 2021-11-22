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
	Error   string
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
		if err := sendOkResponse(w, Response{Message: "I'm alive"}); err != nil {
			panic(err)
		}
	})

}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse Uint from ID", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieving comment by ID", err)
		return
	}

	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}

}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {

	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Error", err)
		return
	}

	if err = sendOkResponse(w, comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {

		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	comment, err := h.Service.PostComment(comment)

	if err != nil {
		sendErrorResponse(w, "Error", err)
		return
	}

	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {

		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	comment, err := h.Service.UpdateComment(1, comment)

	if err != nil {
		sendErrorResponse(w, "Error: failed to update comment", err)
		return
	}

	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		//fmt.Fprintf(w, "Error: %s", err)
		sendErrorResponse(w, "Error: failed to delete comment", err)
		return
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Println(w, "failed to delete comment")
	}

	if err = sendOkResponse(w, Response{Message: "Comment deleted"}); err != nil {
		panic(err)
	}
}

//Send Ok response
func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

//Send Error response
func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
