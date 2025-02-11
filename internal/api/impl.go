package api

import (
	"log"
	"net/http"
)

type Server struct{}

func NewServer() ServerInterface {
	return Server{}
}

// GetApiV1CostsMonthly implements ServerInterface.
func (s Server) GetApiV1CostsMonthly(w http.ResponseWriter, r *http.Request, params GetApiV1CostsMonthlyParams) {
	// TODO 月次コスト取得処理を実装
	log.Println("called GetApiV1CostsMonthly()")
	panic("unimplemented")
}

// GetApiV1CostsTypes implements ServerInterface.
func (s Server) GetApiV1CostsTypes(w http.ResponseWriter, r *http.Request) {
	// TODO コスト種別取得処理を実装
	log.Println("called GetApiV1CostsTypes()")
	panic("unimplemented")
}

// PostApiV1CostsMonthly implements ServerInterface.
func (s Server) PostApiV1CostsMonthly(w http.ResponseWriter, r *http.Request) {
	// TODO 月次コスト登録処理を実装
	log.Println("called PostApiV1CostsMonthly()")
	panic("unimplemented")
}
