package errors

type ErrorResponse struct{
  Code string `json:"code,omitempty"`
  Message string `json:"error,omitempty"`
}

type ErrorDetails struct{
	Error ErrorResponse `json:"error,omitempty"` 
}

func MakeErrorMessage(code string,message string)*ErrorDetails{
	return &ErrorDetails{ErrorResponse{Code: code,Message: message}}
}