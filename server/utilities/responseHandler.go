package utilities

import "encoding/json"

// ResponseWrapper wraps a response to the appropriate format with either
// error or data.
type ResponseWrapper struct {
	Error *ErrorResponse `json:"error,omitempty"`
	Data  interface{}    `json:"data,omitempty"`
}

// ErrorResponse is the generic error type that all error should contain when
// it is returned to the client
type ErrorResponse struct {
	Message     string `json:"message"`
	Endpoint    string `json:"endpoint"`
	RequestType string `json:"request_type"`
}

// CraftErrorResponse creates an error JSON response based on the ErrorResponse interface.
func CraftErrorResponse(e ErrorResponse) (response []byte, err error) {
	return json.Marshal(wrapError(ErrorResponse{
		Message:     e.Message,
		Endpoint:    e.Endpoint,
		RequestType: e.RequestType,
	}))
}

// CraftJSONResponse creates a JSON-typed response body to be returned for API responses.
func CraftJSONResponse(i interface{}, endpoint string, requestType string) (response []byte, err error) {
	jsonResponse, err := json.Marshal(wrapData(i))
	if err != nil {
		// Ignoring the error produced by this marshalling - in the off chance
		// it fails to marshal the error response, it means that something has gone
		// horribly wrong.
		errorResponseJSON, _ := CraftErrorResponse(ErrorResponse{
			Message:     err.Error(),
			Endpoint:    endpoint,
			RequestType: requestType,
		})
		return errorResponseJSON, err
	}
	return jsonResponse, nil
}

func wrapData(i interface{}) ResponseWrapper {
	return ResponseWrapper{
		Data: i,
	}
}

func wrapError(e ErrorResponse) ResponseWrapper {
	return ResponseWrapper{
		Error: &e,
	}
}
