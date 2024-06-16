package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"

	"backendv2/internal/database"
	"backendv2/internal/service/loadout"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	r := chi.NewRouter()
	NewServer.RegisterRoutes(r)

	db := NewServer.db.GetInstance()

	if db == nil {
		fmt.Print("Error getting db instance!")
		return nil
	}

	loadoutStore := loadout.NewStore(db)
	loadoutHandler := loadout.NewHandler(loadoutStore)
	loadoutHandler.RegisterRoutes(r)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
