package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/opiumated/yellowpages/mongo"
	"github.com/opiumated/yellowpages/users"
	"github.com/spf13/viper"
)

var appMode string

func init() {
	//Check what mode the app is being run on( production, development)
	flag.StringVar(&appMode, "appMode", "production", "Application mode")
	flag.Parse()

	if appMode == "production" {
		viper.SetConfigFile("./config/config.json")
	} else {
		viper.SetConfigFile("./config/config.development.json")
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error Reading Application Configurations.", err)
	}
}

func configureRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	return r
}

func main() {
	r := configureRouter()
	var port string
	if !viper.IsSet("server.address") {
		log.Fatal("Port is not initialized...")
	}
	port = viper.Get("server.address").(string)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)
	r.Mount("/users", users.Routes())

	//Database Connection
	mongo.Init()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Yellow pages"))
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Stopping server...")
		os.Exit(1)
	}()
	log.Fatal(http.ListenAndServe(":"+port, r))
}
