package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/canpok1/code-gateway/internal/api"
	"github.com/canpok1/code-gateway/internal/db"
	"github.com/canpok1/code-gateway/internal/environment"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	env, err := environment.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DB HOST: %s\n", env.MysqlHost)
	log.Printf("DB PORT: %v\n", env.MysqlPort)
	log.Printf("DB USER: %v\n", env.MysqlUser)

	database, err := db.Open(env.MysqlHost, env.MysqlPort, env.MysqlUser, env.MysqlPassword, env.MysqlDatabase)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	server := api.NewServer(database)

	r := http.NewServeMux()

	h := api.HandlerWithOptions(server, api.StdHTTPServerOptions{
		BaseRouter:       r,
		ErrorHandlerFunc: api.HandleClientError,
	})

	addr := fmt.Sprintf("0.0.0.0:%d", env.ServerPort)
	log.Printf("listen : %s\n", addr)

	s := &http.Server{
		Handler: h,
		Addr:    addr,
	}

	log.Fatal(s.ListenAndServe())
}
