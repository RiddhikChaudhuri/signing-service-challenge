package api

import (
	"encoding/json"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/service"
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
	if signatureRequest.DeviceID == "" {
		WriteErrorResponse(response, http.StatusBadRequest, []string{"Device Id is required"})
		return
	}

	// Call the service method to create the signature device
	//createdDevice, err := service.signTransaction(&signatureRequest)
	signature, err := service.AddNewSignature(&signatureRequest, s.storage)
	if err != nil {
		// Handle error (you might want to write an error response here)
		WriteErrorResponse(response, http.StatusInternalServerError, []string{err.Error()})
		return
	}
	defer request.Body.Close()
	WriteAPIResponse(response, http.StatusCreated, signature)
}

func (s *Server) FindSignatureByDeviceID(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
	// Parse query parameters
	queryParams := request.URL.Query()
	deviceID := queryParams.Get("deviceId")

	if deviceID == "" {
		WriteErrorResponse(response, http.StatusBadRequest, []string{"Device ID is required"})
		return
	}

	// Call the persistence layer to find the device by ID
	signature, err := s.storage.FindCurrentSignatureByDeviceId(deviceID)
	if err != nil {
		WriteErrorResponse(response, http.StatusNotFound, []string{"Device not found"})
		return
	}
	WriteAPIResponse(response, http.StatusOK, signature)
}

func (s *Server) FindAllSignatureTransactions(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	// Call the persistence layer to find the device by ID
	signedTransactions := s.storage.FindAllSignatures()
	WriteAPIResponse(response, http.StatusOK, signedTransactions)
}
