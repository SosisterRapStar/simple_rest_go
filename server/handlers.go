package server

import (
	"first-proj/domain"
	"net/http"
)

type HTTPHandlersFabric struct {
	NoteService domain.NoteService
}

func (fabric *HTTPHandlersFabric) CreateNote() http.Handler {
	return ...
}


func (fabric *HTTPHandlersFabric) DeleteNote() http.Handler {
	return ...
}


func (fabric *HTTPHandlersFabric) GetNoteById() http.Handler {
	return ...
}


func (fabric *HTTPHandlersFabric) UpdateNote() http.Handler {
	return ...
}
