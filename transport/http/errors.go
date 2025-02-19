package server

import (
	"errors"
	"first-proj/services"
)







type HttpApiError struct{
	code int
	details string
}




func HandleServiceError(err error) HttpApiError{
	apiError := HttpApiError{}
	serviceErr : 
	if errors.As(err)	
}