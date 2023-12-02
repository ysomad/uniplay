// Package writer provides utility functions for writing HTTP responses in a standardized JSON format.
//
// The primary purpose of this package is to simplify the process of responding to HTTP requests with JSON-encoded error messages.
// It includes functions to generate JSON responses for different scenarios, such as returning custom messages or standard HTTP status codes.
//
// Example usage:
//
//	// Respond with a custom error message
//	err := writer.Message(w, http.StatusBadRequest, "Invalid input data")
//
//	// Respond with a standard HTTP status code and its corresponding message
//	err := writer.Status(w, http.StatusOK)
//
//	// Respond with an error object and its corresponding status code
//	err := writer.Error(w, http.StatusInternalServerError, someError)
//
// JSON Format:
//
//	The responses are formatted as JSON objects with the following structure:
//	{
//	  "code": "HTTP Status Code or Custom Code",
//	  "msg": "Error Message or Standard HTTP Status Text"
//	}
//
// Package Naming:
//
//	To avoid naming conflicts with the standard "http" package, this package is named "writer."
//	It provides functionality to write JSON-encoded responses to an http.ResponseWriter.
//
// P.S. ChatGPT is GOAT
package writer

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func Message(w http.ResponseWriter, code int, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(httpError{
		Code: http.StatusText(code),
		Msg:  msg,
	})
}

func Status(w http.ResponseWriter, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	s := http.StatusText(code)
	return json.NewEncoder(w).Encode(httpError{
		Code: s,
		Msg:  s,
	})

}

func Error(w http.ResponseWriter, code int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(httpError{
		Code: http.StatusText(code),
		Msg:  err.Error(),
	})
}
