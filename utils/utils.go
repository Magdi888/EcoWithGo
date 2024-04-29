package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ParseJSON parses the JSON payload from the HTTP request body into the provided payload variable.
//
// Parameters:
// - r: the HTTP request object.
// - payload: a pointer to the variable where the parsed JSON payload will be stored.
//
// Returns:
// - error: an error if the request body is empty or if there was an error decoding the JSON payload.

func ParseJSON(r *http.Request,payload any) error  {
	if r.Body  == nil{
		return fmt.Errorf( "empty request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
	}


// WriteJSON writes the provided payload as JSON to the HTTP response writer.
//
// Parameters:
// - w: the HTTP response writer.
// - statuscode: the HTTP status code to be set in the response.
// - payload: the data to be encoded as JSON and written to the response.
//
// Returns:
// - error: an error if there was a problem encoding the payload as JSON or writing it to the response.

func WriteJSON(w http.ResponseWriter, statuscode int,payload any) error{
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	// encode the response as JSON format
	return json.NewEncoder(w).Encode(payload)

}

// WriteError writes the provided error as JSON to the HTTP response writer.
//
// Parameters:
// - w: the HTTP response writer.
// - statuscode: the HTTP status code to be set in the response.
// - err: the error to be written in the response.
//
// Returns:
// - None.
func WriteError(w http.ResponseWriter, statuscode int, err error) {
	WriteJSON(w,statuscode,map[string]string{"error":err.Error()})
}

