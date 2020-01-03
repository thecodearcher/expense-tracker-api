package helpers

import (
	"encoding/json"
	"net/http"
)

//APIResponse data format for api response
type APIResponse struct {
	Data    interface{} `json:"data"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
}

//Success sends JSON response to writer
func (res *APIResponse) Success(w http.ResponseWriter) {
	res.Status = true
	if res.Message == "" {
		res.Message = "OK"
	}

	//Send header, status code and output to writer
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

//Created send back a 201
func (res *APIResponse) Created(w http.ResponseWriter) {
	res.Status = true
	if res.Message == "" {
		res.Message = "Resource created!"
	}

	//Send header, status code and output to writer
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (res *APIResponse) Error(statusCode int, w http.ResponseWriter) {
	res.Status = false
	if statusCode == http.StatusInternalServerError && res.Message == "" {
		res.Message = http.StatusText(http.StatusInternalServerError)
	}
	if res.Message == "" {
		panic("Error message not provied for Error APIResponse")
	}

	//Send header, status code and output to writer
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

//Custom sends any
func (res *APIResponse) Custom(statusCode int, w http.ResponseWriter) {
	if res.Message == "" {
		panic("message can not be blank")
	}

	//Send header, status code and output to writer
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}
