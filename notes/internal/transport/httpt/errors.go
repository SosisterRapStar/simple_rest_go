package httpt

// import (
// 	"errors"
// 	"first-proj/domain"
// 	"first-proj/services"
// 	"net/http"
// )

// type HttpApiError struct {
// 	Status  int
// 	Details string
// }

// func NewHttpApiError(status int, details string) *HttpApiError {
// 	return &HttpApiError{
// 		Status:  status,
// 		Details: details,
// 	}
// }

// func HandleServiceError(err error) HttpApiError {
// 	var apiError HttpApiError
// 	var serviceError services.ServiceError
// 	if errors.As(err, &serviceError) {
// 		appError := serviceError.GetAppError()
// 		apiError.Details = appError.Error()
// 		switch appError {
// 		case services.ErrNoteNotFound:
// 			apiError.Status = http.StatusNotFound
// 		case domain.ErrNoteValidation:
// 			apiError.Status = http.StatusBadRequest
// 		case services.ErrTimeOutExceeded:
// 			apiError.Status = http.StatusRequestTimeout
// 		default:
// 			apiError.Status = http.StatusBadRequest

// 		}
// 	}
// 	return apiError
// }
