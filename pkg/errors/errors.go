package errors

type ApiError struct {
	Message    string   `json:"message"`
	Code       int      `json:"code"`
	Name       string   `json:"name"`
	Error      error    `json:"-"`
	Detail     string   `json:"-"`
	Validation []string `json:"validation,omitempty"`
}

func BindError() *ApiError {
	return &ApiError{Code: 400, Message: "Error processing request.", Name: "BIND_ERROR"}
}

func ValidationError(errors []string) *ApiError {
	return &ApiError{Code: 400, Name: "VALIDATION", Message: "A validation error occurred.", Validation: errors}
}
