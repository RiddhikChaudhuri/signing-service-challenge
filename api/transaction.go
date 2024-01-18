package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
	// Parse request body
	var signatureRequest domain.SignatureRequest
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&signatureRequest); err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"Error decoding JSON request",
		})
		return
	}
	fmt.Println("Are we coming here")
	// Validation Check For Device Id
	if signatureRequest.DeviceID == "" {
		WriteErrorResponse(response, http.StatusBadRequest, []string{"Device Id is required"})
		return
	}

	// Call the service method to create the signature device
	//createdDevice, err := service.signTransaction(&signatureRequest)

	defer request.Body.Close()
	WriteAPIResponse(response, http.StatusCreated, signatureRequest)
}
