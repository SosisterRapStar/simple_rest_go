package server

import (
	"context"
	"first-proj/domain"
	"net/http"
)

type HttpApi struct {
	NoteService domain.NoteService
}



func (api *HttpApi) CreateNote(w http.ResponseWriter, r *http.Request) {
	return
}

func (api *HttpApi) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	deleted_id, err := api.NoteService.DeleteNote(context.Background(), id)
	if err != nil {
		errorToSend := HandleServiceError(err)
		http.Error(w, errorToSend.Details, errorToSend.Status)
		return
	}

}

func (api *HttpApi) GetNoteById(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
	return
}

func (api *HttpApi) UpdateNote(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
	return
}

func (api *HttpApi) parseJson(w http.ResponseWriter, r *http.Request)

func (api *HttpApi) writeJson(w http.ResponseWriter, r *http.Request, data string) {
	w.
}
