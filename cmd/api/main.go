package main

import (
	"fmt"
	"go-landing-page/internal/database"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port string

	db database.Service
}

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	s := &Server{
		port: fmt.Sprintf(":%d", port),
		db:   database.NewConnectionDatabase(),
	}

	server := &http.Server{
		Addr: s.port,
		// Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}

	log.Println("teste")
}
