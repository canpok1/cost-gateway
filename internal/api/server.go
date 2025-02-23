package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	db *sql.DB
}

func NewServer(database *sql.DB) ServerInterface {
	return &Server{
		db: database,
	}
}

// GetApiV1CostsMonthly implements ServerInterface.
func (s *Server) GetApiV1CostsMonthly(w http.ResponseWriter, r *http.Request, params GetApiV1CostsMonthlyParams) {
	ctx := context.Background()

	log.Println("called GetApiV1CostsMonthly()")

	resp, err := s.getApiV1CostsMonthly(ctx, &params)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorObject{
			Message: "internal server error",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// GetApiV1CostsTypes implements ServerInterface.
func (s *Server) GetApiV1CostsTypes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	log.Println("called GetApiV1CostsTypes()")

	resp, err := s.getApiV1CostsTypes(ctx)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorObject{
			Message: "internal server error",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// PostApiV1CostsMonthly implements ServerInterface.
func (s *Server) PostApiV1CostsMonthly(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	log.Println("called PostApiV1CostsMonthly()")

	var body PostApiV1CostsMonthlyJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("error occured: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorObject{
			Message: fmt.Sprintf("failed to parse request body, %v", err),
		})
		return
	}

	resp, err := s.postApiV1CostsMonthly(ctx, &body)
	if err != nil {
		log.Printf("error occured: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorObject{
			Message: "internal server error",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
