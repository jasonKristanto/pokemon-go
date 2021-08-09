package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SendResponse(resp http.ResponseWriter, httpCode int, responseData Response)  {
	resp.Header().Set("Content-Type", "application/json")

	jsonInBytes, err := json.Marshal(responseData)

	if err != nil {
		errorCode, errorResponse := InternalServerErrorResponse()
		httpCode = errorCode
		jsonInBytes, _ = json.Marshal(errorResponse)
	}

	resp.WriteHeader(httpCode)
	resp.Write(jsonInBytes)
	return
}

func SuccessResponse(message string, data interface{}) (int, Response) {
	return 200, Response {
		Status: "success",
		Message: message,
		Data: data,
	}
}

func ErrorResponse(httpCode int, message string) (int, Response) {
	return httpCode, Response {
		Status: "error",
		Message: message,
	}
}

func InvalidRequestResponse() (int, Response) {
	return ErrorResponse(400, "INVALID_REQUEST")
}

func PageNotFoundResponse() (int, Response) {
	return ErrorResponse(404, "PAGE_NOT_FOUND")
}

func DataNotFoundResponse() (int, Response) {
	return ErrorResponse(404, "DATA_NOT_FOUND")
}

func MethodNotAllowedResponse() (int, Response) {
	return ErrorResponse(405, "METHOD_NOT_ALLOWED")
}

func InternalServerErrorResponse() (int, Response) {
	return ErrorResponse(500, "INTERNAL_SERVER_ERROR")
}