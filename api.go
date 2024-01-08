package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Server struct
type ApiServer struct {
	ListenPort string
	Store      Storage
}

// Run the server
func (a *ApiServer) Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/getData", makeHttpFunc(a.GetDataFilter))
	r.HandleFunc("/getDataAll", makeHttpFunc(a.GetDataAll))

	err := http.ListenAndServe(a.ListenPort, r)
	if err != nil {
		return err
	}
	return nil
}

func NewAPIServer(port string, store Storage) *ApiServer {
	return &ApiServer{
		ListenPort: port,
		Store:      store,
	}
}

// wrapper for http.Handler
type apiFunc func(http.ResponseWriter, *http.Request) error

// error struct for handling handler's Error
type ApiError struct {
	Error string `json:"error"`
}

// Handler Get Data Filtered
func (a *ApiServer) GetDataFilter(w http.ResponseWriter, r *http.Request) error {
	request := &RecordRequest{}
	err := DecodeJsonReq(r, request)
	if err != nil {
		return err
	}

	records, err := a.Store.GetDataFilter(request)
	if err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, RecordResponse{Code: 0, Msg: "Success", Records: records})
	return nil
}

// Handler Get All Data
func (a *ApiServer) GetDataAll(w http.ResponseWriter, r *http.Request) error {
	request := &RecordRequest{}
	err := DecodeJsonReq(r, request)
	if err != nil {
		return err
	}

	records, err := a.Store.GetDataAll(request)
	if err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, records)
	return nil
}

// to return http.HandlerFunc, for easier error handling
func makeHttpFunc(a apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			// Handle Erorr handler
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

// Func to write json
func WriteJson(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

// Func to decode from json Request
func DecodeJsonReq(r *http.Request, records *RecordRequest) error {
	err := json.NewDecoder(r.Body).Decode(records)
	defer r.Body.Close()
	return err
}
