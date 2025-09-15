package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LSUDOKO/RestAPI/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Write([]byte("welcome to student api"))
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	//setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}
	fmt.Printf("sever started %s", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("failed to start server %s", err.Error())
	}

}
