package server

import (
	"errors"
	"first-proj/services"
	"fmt"
	"net/http"
)

type HttpApiError struct {
	Status  int
	Details string
}

func NewHttpApiError(status int, details string) *HttpApiError {
	return &HttpApiError{
		Status:  status,
		Details: details,
	}
}

func HandleServiceError(err error) HttpApiError {
	var apiError HttpApiError
	var serviceError services.ServiceError
	if errors.As(err, &serviceError) {
		appError := serviceError.GetAppError()
		apiError.Details = appError.Error()
		switch appError {
		case services.ErrNoteNotFound:
			apiError.Status = http.StatusNotFound
			fmt.Println(serviceError.GetActualError().Error()) // в будущем будет залогировано
		}
	}
	return apiError
}
