package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/canpok1/code-gateway/internal/db"
)

type Server struct {
	client *db.Queries
}

func NewServer(client *db.Queries) ServerInterface {
	return Server{
		client: client,
	}
}

// GetApiV1CostsMonthly implements ServerInterface.
func (s Server) GetApiV1CostsMonthly(w http.ResponseWriter, r *http.Request, params GetApiV1CostsMonthlyParams) {
	// TODO 月次コスト取得処理を実装
	log.Println("called GetApiV1CostsMonthly()")
	panic("unimplemented")
}

// GetApiV1CostsTypes implements ServerInterface.
func (s Server) GetApiV1CostsTypes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	log.Println("called GetApiV1CostsTypes()")

	resp, err := getApiV1CostsTypes(ctx, s.client)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Printf("%v\n", resp)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// PostApiV1CostsMonthly implements ServerInterface.
func (s Server) PostApiV1CostsMonthly(w http.ResponseWriter, r *http.Request) {
	// TODO 月次コスト登録処理を実装
	log.Println("called PostApiV1CostsMonthly()")
	panic("unimplemented")
}
