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

func HandleServiceError(err error) HttpApiError {
	var apiError HttpApiError
	var serviceError services.ServiceError
	if errors.As(err, &serviceError) {
		appError := serviceError.GetAppError()
		apiError.Details = appError.Error()
		switch appError {
		case services.ErrNoteNotFound:
			apiError.Status = http.StatusNotFound
			fmt.Println(serviceError.GetActualError()) // в будущем будет залогировано
		}
	}
	return apiError
}
