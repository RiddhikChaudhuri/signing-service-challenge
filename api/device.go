package api

import (
	"encoding/json"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/service"
)

// SignatureResponse represents the response after signing data

func (s *Server) CreateSignatureDevice(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
	// Parse request body
	var createRequest domain.CreateSignatureDeviceRequest
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&createRequest); err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"Error decoding JSON request",
		})
		return
	}
	// Validation Check For Label
	if createRequest.Label == "" {
		WriteErrorResponse(response, http.StatusBadRequest, []string{"Label is required"})
		return
	}

	// Call the service method to create the signature device
	createdDevice, err := service.AddNewSignatureDevice(&createRequest, s.storage)
	if err != nil {
		// Handle error (you might want to write an error response here)
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			"Error creating signature device",
		})
		return
	}
	defer request.Body.Close()
	WriteAPIResponse(response, http.StatusCreated, createdDevice)
}

func (s *Server) FindSignatureDeviceByID(response http.ResponseWriter, request *http.Request) {
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
	signatureDevice, err := s.storage.FindSignatureDeviceByID(deviceID)
	if err != nil {
		WriteErrorResponse(response, http.StatusNotFound, []string{"Device not found"})
		return
	}
	WriteAPIResponse(response, http.StatusOK, signatureDevice)
}

func (s *Server) FindAllSignatureDevices(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	// Call the persistence layer to find the device by ID
	signatureDevices := s.storage.FindAllSignatureDevices()
	WriteAPIResponse(response, http.StatusOK, signatureDevices)
}
