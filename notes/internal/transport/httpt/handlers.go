package httpt

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"first-proj/domain"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// var MAXREQUESTIME = 10 * time.Second

// type HttpApi interface {
// 	CreateNote(w http.ResponseWriter, r *http.Request)
// 	DeleteNote(w http.ResponseWriter, r *http.Request)
// 	GetNoteById(w http.ResponseWriter, r *http.Request)
// 	FindNotes(w http.ResponseWriter, r *http.Request)
// 	UpdateNote(w http.ResponseWriter, r *http.Request)
// }

// type HttpApiHandlers struct {
// 	noteService domain.NoteService
// }

// func NewHttpApiHandlers(noteService domain.NoteService) *HttpApiHandlers {
// 	return &HttpApiHandlers{noteService: noteService}
// }

// func (api *HttpApiHandlers) setCommonHeaders(w http.ResponseWriter) {
// 	w.Header().Set("Cache-Control", "no-store, max-age=0")
// 	w.Header().Set("Pragma", "no-cache")
// }

// func (api *HttpApiHandlers) checkContentType(r *http.Request) error {
// 	ct := r.Header.Get("Content-Type")
// 	if ct != "" {
// 		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
// 		if mediaType != "application/json" {
// 			return errors.New("content type header is not application/json")
// 		}
// 	}
// 	return nil
// }

// func (api *HttpApiHandlers) decodeJson(decoder *json.Decoder, v any) *HttpApiError {
// 	if err := decoder.Decode(v); err != nil {
// 		var syntaxError *json.SyntaxError
// 		var unmarshalTypeError *json.UnmarshalTypeError
// 		var maxBytesError *http.MaxBytesError

// 		switch {
// 		case errors.As(err, &syntaxError):
// 			return NewHttpApiError(http.StatusBadRequest, fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset))

// 		case errors.Is(err, io.ErrUnexpectedEOF):
// 			return NewHttpApiError(http.StatusBadRequest, "Request body contains badly-formed JSON")

// 		case errors.As(err, &unmarshalTypeError):
// 			return NewHttpApiError(http.StatusBadRequest, fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset))

// 		case errors.Is(err, io.EOF):
// 			return NewHttpApiError(http.StatusBadRequest, "Request body must not be empty")

// 		case errors.As(err, &maxBytesError):
// 			return NewHttpApiError(http.StatusRequestEntityTooLarge, fmt.Sprintf("Request body must not be larger than %d bytes", maxBytesError.Limit))
// 		default:
// 			return NewHttpApiError(http.StatusInternalServerError, "Internal server error")
// 		}
// 	}
// 	return nil
// }

// func (api *HttpApiHandlers) CreateNote(w http.ResponseWriter, r *http.Request) {
// 	if err := api.checkContentType(r); err != nil {
// 		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
// 	}
// 	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
// 	dec := json.NewDecoder(r.Body)
// 	dec.DisallowUnknownFields()
// 	var note domain.Note
// 	if err := api.decodeJson(dec, &note); err != nil {
// 		http.Error(w, err.Details, err.Status)
// 		return
// 	}
// 	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
// 		http.Error(w, "Request should contain only one json structure", http.StatusBadRequest)
// 		return
// 	}

// 	serviceCTX, cancel := context.WithTimeout(r.Context(), MAXREQUESTIME)
// 	defer cancel()
// 	response, err := api.noteService.CreateNote(serviceCTX, &note)
// 	if err != nil {
// 		sendError := HandleServiceError(err)
// 		http.Error(w, sendError.Details, sendError.Status)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")

// 	// w.WriteHeader(http.StatusOK)
// 	api.writeJSON(w, http.StatusOK, response)
// }

// func (api *HttpApiHandlers) DeleteNote(w http.ResponseWriter, r *http.Request) {
// 	id := r.PathValue("id")
// 	serviceCTX, cancel := context.WithTimeout(r.Context(), MAXREQUESTIME)
// 	defer cancel()
// 	deleted_id, err := api.noteService.DeleteNote(serviceCTX, id)
// 	if err != nil {
// 		errorToSend := HandleServiceError(err)
// 		http.Error(w, errorToSend.Details, errorToSend.Status)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")

// 	w.WriteHeader(http.StatusOK)
// 	response := map[string]string{"message": "Note was deleted", "id": deleted_id}
// 	api.writeJSON(w, http.StatusOK, response)
// }

// func (api *HttpApiHandlers) GetNoteById(w http.ResponseWriter, r *http.Request) {
// 	id := r.PathValue("id")
// 	serviceCTX, cancel := context.WithTimeout(r.Context(), MAXREQUESTIME)
// 	defer cancel()
// 	note, err := api.noteService.GetNote(serviceCTX, id)
// 	if err != nil {
// 		errorToSend := HandleServiceError(err)
// 		http.Error(w, errorToSend.Details, errorToSend.Status)
// 		return
// 	}
// 	api.setCommonHeaders(w)
// 	w.Header().Set("Content-Type", "application/json")
// 	api.writeJSON(w, http.StatusOK, note)
// }

// func (api *HttpApiHandlers) UpdateNote(w http.ResponseWriter, r *http.Request) {
// 	if err := api.checkContentType(r); err != nil {
// 		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
// 	}
// 	id := r.PathValue("id")
// 	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
// 	dec := json.NewDecoder(r.Body)
// 	dec.DisallowUnknownFields()
// 	var updNote domain.UpdateNote
// 	if err := api.decodeJson(dec, &updNote); err != nil {
// 		http.Error(w, err.Details, err.Status)
// 		return
// 	}
// 	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
// 		http.Error(w, "Request should contain only one json structure", http.StatusBadRequest)
// 		return
// 	}
// 	serviceCTX, cancel := context.WithTimeout(r.Context(), MAXREQUESTIME)
// 	defer cancel()

// 	response, err := api.noteService.UpdateNote(serviceCTX, &updNote, id)
// 	if err != nil {
// 		sendError := HandleServiceError(err)
// 		http.Error(w, sendError.Details, sendError.Status)
// 		return
// 	}
// 	api.setCommonHeaders(w)
// 	api.writeJSON(w, http.StatusOK, response)

// }

// func (api *HttpApiHandlers) FindNotes(w http.ResponseWriter, r *http.Request) {
// 	var paginationFilter = domain.PaginateFilter{}

// 	nextPageToken := r.URL.Query().Get("token")
// 	if nextPageToken != "" {
// 		paginationFilter.NextPageToken = &nextPageToken
// 	}
// 	limit := r.URL.Query().Get("limit")
// 	if limit == "" {
// 		http.Error(w, `{"error": "Unacceptable query param value"}`, http.StatusBadRequest)
// 		return
// 	}
// 	limit_value, err := strconv.Atoi(limit)

// 	if err != nil {
// 		http.Error(w, `{"error": "Unacceptable query param value"}`, http.StatusBadRequest)
// 		return
// 	}
// 	if limit_value < 0 {
// 		http.Error(w, `{"error": "Unacceptable query param value"}`, http.StatusBadRequest)
// 		return
// 	}
// 	paginationFilter.Limit = &limit_value
// 	serviceCTX, cancel := context.WithTimeout(r.Context(), MAXREQUESTIME)
// 	defer cancel()
// 	notes, notes_num, nextPageToken, err := api.noteService.FindNotes(serviceCTX, &paginationFilter)
// 	if err != nil {
// 		errorFromService := HandleServiceError(err)
// 		http.Error(w, errorFromService.Details, errorFromService.Status)
// 		return
// 	}

// 	response := struct {
// 		Notes         []*domain.Note `json:"notes"`
// 		Notes_num     int            `json:"notes_num"`
// 		NextPageToken string         `json:"next_page_token,omitempty"`
// 	}{
// 		Notes:         notes,
// 		Notes_num:     notes_num,
// 		NextPageToken: nextPageToken,
// 	}
// 	api.setCommonHeaders(w)
// 	api.writeJSON(w, http.StatusOK, response)
// }

// func (api *HttpApiHandlers) writeJSON(w http.ResponseWriter, status int, data interface{}) {
// 	// api.setCommonHeaders(w)
// 	w.WriteHeader(status)
// 	if err := json.NewEncoder(w).Encode(data); err != nil {
// 		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
// 	}
// }
