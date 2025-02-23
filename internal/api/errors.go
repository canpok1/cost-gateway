package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleClientError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("error occured: %v\n", err)

	resp := ErrorObject{
		Message: fmt.Sprintf("%v", err),
	}
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(resp)
}
