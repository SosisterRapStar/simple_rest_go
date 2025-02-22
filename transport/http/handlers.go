package server

import (
	"encoding/json"
	"errors"
	"first-proj/domain"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HttpApi struct {
	NoteService domain.NoteService
}

func (api *HttpApi) setCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store, max-age=0")
	W.Header().Set("Pragma", "no-cache")
	w.Header().Set("Content-Type", "application/json")
}

func (api *HttpApi) checkContentType(w http.ResponseWriter) error {
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			return errors.New("Content type header is not application/json")
		}
	}
	return nil
}

func (api *HttpApi) decodeJson(decoder *json.Decoder, v any) *HttpApiError {
	if err := decoder.Decode(v); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return NewHttpApiError(http.StatusBadRequest, fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset))

		case errors.Is(err, io.ErrUnexpectedEOF):
			return NewHttpApiError(http.StatusBadRequest, "Request body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			return NewHttpApiError(http.StatusBadRequest, fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset))

		case errors.Is(err, io.EOF):
			return NewHttpApiError(http.StatusBadRequest, "Request body must not be empty")

		case errors.As(err, &maxBytesError):
			return NewHttpApiError(http.StatusRequestEntityTooLarge, fmt.Sprintf("Request body must not be larger than %d bytes", maxBytesError.Limit))
		default:
			return NewHttpApiError(http.StatusInternalServerError, "Internal server error")
		}
	}
	return nil
}

func (api *HttpApi) CreateNote(w http.ResponseWriter, r *http.Request) {
	if err := api.checkContentType(w); err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var note domain.Note
	if err := api.decodeJson(dec, &note); err != nil {
		http.Error(w, err.Details, err.Status)
		return
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		http.Error(w, "Request should contain only one json structure", http.StatusBadRequest)
		return
	}
	note_id, err := api.NoteService.CreateNote(r.Context(), &note)
	if err != nil {
		sendError := HandleServiceError(err)
		http.Error(w, sendError.Details, sendError.Status)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	return

}

func (api *HttpApi) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	deleted_id, err := api.NoteService.DeleteNote(r.Context(), id)
	if err != nil {
		errorToSend := HandleServiceError(err)
		http.Error(w, errorToSend.Details, errorToSend.Status)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Note was deleted", "id": deleted_id}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		return
	}
}

func (api *HttpApi) GetNoteById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	note, err := api.NoteService.GetNote(r.Context(), id)
	if err != nil {
		errorToSend := HandleServiceError(err)
		http.Error(w, errorToSend.Details, errorToSend.Status)
		return
	}
	api.setCommonHeaders(w)
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		return
	}
}

func (api *HttpApi) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var note domain.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, `{"error": "Failed to decode request"}`, http.StatusInternalServerError)
		return
	}
	note, err := api.NoteService.UpdateNote(r.Context(), note)
	if err != nil {
		sendErr := HandleServiceError(err)
		http.Error(w, sendErr.Details, sendErr.Status)
		return
	}

}

// func (api *HttpApi) parseJson(w http.ResponseWriter, r *http.Request) {

// }

// func (api *HttpApi) writeJson(w http.ResponseWriter, r *http.Request, data string) {
// 	w.
// }
