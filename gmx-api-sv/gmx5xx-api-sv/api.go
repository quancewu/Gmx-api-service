package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func Gmx5xxAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/version", makeHTTPHandleFunc(s.handleGmxVerison))
	router.HandleFunc("/api/v1/gmxdata", makeHTTPHandleFunc(s.handleGmxData))
	router.HandleFunc("/api/v1/gmxdata/latest", makeHTTPHandleFunc(s.handleGmxDataLatest))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleGmxVerison(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGmxGetVersion(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGmxGetVersion(w http.ResponseWriter, r *http.Request) error {

	datas := map[string]interface{}{
		"api_verison": "v1.1.0",
		"build_at":    "2023-11-02T15:00:00Z",
	}

	return WriteJSON(w, http.StatusOK, datas)
}

func (s *APIServer) handleGmxData(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGmxGetData(w, r)
	}
	if r.Method == "POST" {
		return s.handleGmxPostData(w, r)

	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGmxGetData(w http.ResponseWriter, r *http.Request) error {

	dataRequest := new(DataReqestDate)
	dataAfter := r.URL.Query().Get("after")
	dataBefore := r.URL.Query().Get("before")
	if dataAfter == "" || dataBefore == "" {
		log.Println("No dataAfter or data Before")
		dataAfter = time.Now().UTC().Format(time.RFC3339)
		dataRequest.After = time.Now().UTC().Add(-1 * time.Hour)
		dataRequest.Before = time.Now().UTC()
	} else {
		response, err := validateRequestDate(dataAfter, dataBefore)
		dataRequest = response
		if err != nil {
			return err
		}
		// log.Println(dataRequest.After, dataRequest.Before)
	}

	datas, err := s.store.GetDatas(dataRequest)
	if err != nil {
		return err
	}
	// datas := map[string]interface{}{
	// 	"timestamp": "2023-10-06T00:00:00z",
	// 	"temperture": 23.23,
	// }

	return WriteJSON(w, http.StatusOK, datas)
}

func (s *APIServer) handleGmxPostData(w http.ResponseWriter, r *http.Request) error {
	req := new(InsertDataRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	data, err := NewData(*req)
	if err != nil {
		return err
	}
	if err := s.store.InsertData(data); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, data)
}

func (s *APIServer) handleGmxDataLatest(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGmxGetDataLatest(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGmxGetDataLatest(w http.ResponseWriter, r *http.Request) error {
	datas, err := s.store.GetLatestData()
	if err != nil {
		return err
	}
	// datas := map[string]interface{}{
	// 	"timestamp": "2023-10-06T00:00:00z",
	// 	"temperture": 23.23,
	// }

	return WriteJSON(w, http.StatusOK, datas)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func validateRequestDate(after string, before string) (*DataReqestDate, error) {
	dataRequest := new(DataReqestDate)
	t1, err := time.Parse(time.RFC3339, after)
	if err != nil {
		return nil, err
	}
	dataRequest.After = t1
	t2, err := time.Parse(time.RFC3339, before)
	if err != nil {
		return nil, err
	}
	dataRequest.Before = t2
	diff := t2.Sub(t1)
	if diff > 24*time.Hour {
		return dataRequest, fmt.Errorf("time delta %v lager than 1 day", diff)
	}
	if diff < 0 {
		return dataRequest, fmt.Errorf("time delta %v you need learn English", diff)
	}
	return dataRequest, nil
}
