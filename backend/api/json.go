package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JsonWrapper(handler func(http.ResponseWriter, *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		data, err := handler(response, request)

		response.Header().Set("Content-Type", "application/json")

		if err != nil {
			errResponse := ErrorResponse{Message: "An error occurred processing your request", Error: err.Error()}
			jsonErrData, _ := json.Marshal(errResponse)
			http.Error(response, string(jsonErrData), http.StatusInternalServerError)
			return
		}

		successResponse := SuccessResponse{Message: "Request processed successfully", Data: data}
		jsonSuccessData, err := json.Marshal(successResponse)

		if err != nil {
			errResponse := ErrorResponse{Message: "An error occurred processing your request", Error: err.Error()}
			jsonErrData, _ := json.Marshal(errResponse)
			http.Error(response, string(jsonErrData), http.StatusInternalServerError)
			return
		}

		response.Write(jsonSuccessData)
	}
}
