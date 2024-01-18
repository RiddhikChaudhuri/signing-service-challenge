package api

import (
	"encoding/json"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
)

// Response is the generic API response container.
type Response struct {
	Data interface{} `json:"data"`
}

// ErrorResponse is the generic error API response container.
type ErrorResponse struct {
	Errors []string `json:"errors"`
}

// Server manages HTTP requests and dispatches them to the appropriate services.
type Server struct {
	listenAddress string
	storage       persistence.Storage
}

// NewServer is a factory to instantiate a new Server.
func NewServer(listenAddress string, storage persistence.Storage) *Server {
	return &Server{
		listenAddress: listenAddress,
		storage:       storage,
		// TODO: add services / further dependencies here ...
	}
}

// Run registers all HandlerFuncs for the existing HTTP routes and starts the Server.
func (s *Server) Run() error {
	mux := http.NewServeMux()

	mux.Handle("/api/v0/health", http.HandlerFunc(s.Health))
	mux.Handle("/api/v0/signaturedevice/add", http.HandlerFunc(s.CreateSignatureDevice))
	mux.Handle("/api/v0/signaturedevice/search", http.HandlerFunc(s.FindSignatureDeviceByID))
	mux.Handle("/api/v0/signaturedevice/findAll", http.HandlerFunc(s.FindAllSignatureDevices))
	mux.Handle("/api/v0/sign/transaction", http.HandlerFunc(s.SignTransaction))
	mux.Handle("/api/v0/sign/transaction/search", http.HandlerFunc(s.FindSignatureByDeviceID))
	mux.Handle("/api/v0/sign/transaction/findAll", http.HandlerFunc(s.FindAllSignatureTransactions))

	// TODO: register further HandlerFuncs here ...

	return http.ListenAndServe(s.listenAddress, mux)
}

// WriteInternalError writes a default internal error message as an HTTP response.
func WriteInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

// WriteErrorResponse takes an HTTP status code and a slice of errors
// and writes those as an HTTP error response in a structured format.
func WriteErrorResponse(w http.ResponseWriter, code int, errors []string) {
	w.WriteHeader(code)

	errorResponse := ErrorResponse{
		Errors: errors,
	}

	bytes, err := json.Marshal(errorResponse)
	if err != nil {
		WriteInternalError(w)
	}

	w.Write(bytes)
}

// WriteAPIResponse takes an HTTP status code and a generic data struct
// and writes those as an HTTP response in a structured format.
func WriteAPIResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json") // Set Content-Type header

	w.WriteHeader(code)

	response := Response{
		Data: data,
	}

	bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		WriteInternalError(w)
	}

	w.Write(bytes)
}
