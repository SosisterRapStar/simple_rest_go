package httpt

import (
	"net/http"
)

func EndpointRegistration(mux *http.ServeMux, handler *HttpApi) {
	mux.HandleFunc("GET /api/v1/notes/{id}", handler.GetNoteById)
	mux.HandleFunc("POST /api/v1/notes", handler.CreateNote)
	mux.HandleFunc("PATCH /api/v1/notes/{id}", handler.UpdateNote)
	mux.HandleFunc("DELETE /api/v1/notes/{id}", handler.DeleteNote)
	mux.HandleFunc("GET /api/v1/notes/", handler.FindNotes)
}
