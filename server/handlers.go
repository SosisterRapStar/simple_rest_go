package server

import (
	"first-proj/domain"
	"net/http"
)

type HTTPHandlersFabric struct {
	NoteService domain.NoteService
}

func (fabric *HTTPHandlersFabric) CreateNote(w ResponseWriter, r *Request){
	return ...
}


func (fabric *HTTPHandlersFabric) DeleteNote(w ResponseWriter, r *Request){
	ret
}


func (fabric *HTTPHandlersFabric) GetNoteById(w ResponseWriter, r *Request){
	return ...
}


func (fabric *HTTPHandlersFabric) UpdateNote(w ResponseWriter, r *Request){
	return ...
}
