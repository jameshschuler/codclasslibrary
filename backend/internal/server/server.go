package server

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"

	"backend/internal/database"
	"backend/internal/service/loadout"
)

type Server struct {
	port int

	db database.Service
}

var validate *validator.Validate

func NewServer() *http.Server {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})

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
	loadoutHandler := loadout.NewHandler(loadoutStore, validate)
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
