package helpers

import (
	"encoding/json"
	"net/http"
)

// Handler is a custom type for handler functions that return an error.
type Handler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

// Wrapper wraps the handler function, managing responses and error handling.
func ResponseWrapper(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the handler function
		data, err := h(w, r)

		// If there is an error, send an error response
		if err != nil {
			SendResponse(w, err, nil)
			return
		}

		// Send a success response if no error occurred
		SendResponse(w, nil, data)
	}
}

// SendResponse sends a structured JSON response with a dynamic status code.
func SendResponse(w http.ResponseWriter, err error, data interface{}) {
	var code int
	var msg string

	if err != nil {
		// If there is an error, set code to 500 and use the error message
		code = http.StatusInternalServerError
		msg = err.Error()
	} else {
		// If no error, set code to 200 and set the message to "success"
		code = http.StatusOK
		msg = "success"
	}

	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code": code,
			"msg":  msg,
		},
		"data": data,
	}

	// Always set HTTP status to 200 but include actual code in JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
