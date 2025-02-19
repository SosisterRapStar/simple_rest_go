package server

import (
	"context"
	"first-proj/domain"
	"net/http"
)

type  HttpApi struct {
	NoteService domain.NoteService
}

type httpHandler func(http.ResponseWriter, *http.Request) error 



func (api *HttpAppi)  CreateNote(w ResponseWriter, r *Request){
	
}


func (api *HttpAppi)  DeleteNote(w ResponseWriter, r *Request) error{
	id := r.PathValue("id")
	context := context.Background()
	id, err := fabric.NoteService.DeleteNote(context, id)
	if err != nil {
		return err
	} ?
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(id))

}


func (api *HttpAppi)  GetNoteById(w ResponseWriter, r *Request){
	id := r.PathValue("id")
}


func (fabric *HttpApi) UpdateNote(w ResponseWriter, r *Request){
	id := r.PathValue("id")
}
