package app

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/iamanders/tdr3/internal/data"
	"gorm.io/gorm"
)

type App struct {
	Port       int
	Host       string
	Production bool

	InfoLog  *log.Logger
	ErrorLog *log.Logger

	Router *chi.Mux
	DB     *gorm.DB
}

func Run() error {

	// Create and setup application
	appInstance := newApp(data.DB)

	// Setup database
	err := data.SetupDatabase(!appInstance.Production)
	if err != nil {
		return errors.New("could not setup database, " + err.Error())
	}
	err = data.MigrateDatabase()
	if err != nil {
		return errors.New("could not migrate database tables, " + err.Error())
	}

	// Serve
	err = appInstance.serve()
	if err != nil {
		return errors.New("could not start http server, " + err.Error())
	}

	return nil
}

func newApp(db *gorm.DB) *App {
	app := &App{
		Host:       "0.0.0.0",
		Port:       8888,
		InfoLog:    log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLog:   log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Router:     chi.NewMux(),
		DB:         db,
		Production: true,
	}

	// Production mode?
	if prod := os.Getenv("PRODUCTION"); strings.ToLower(prod) != "true" {
		app.Production = false
	}

	// Port?
	if port := os.Getenv("PORT"); port != "" {
		app.Port, _ = strconv.Atoi(port)
	}

	app.setupCors()
	app.setupMiddlewares()
	app.setupRoutes()

	return app
}

func (app *App) setupMiddlewares() {
	app.Router.Use(middleware.RequestID)
	app.Router.Use(middleware.RealIP)
	app.Router.Use(middleware.Logger)
	app.Router.Use(middleware.Recoverer)
}

func (app *App) setupRoutes() {
	app.Router.Route("/api", func(r chi.Router) {
		// r.Get("/", app.handleIndex())

		// Status
		r.Get("/status", app.handleStatus())
		r.Put("/status/stop", app.handleStatusStop())

		// Times
		r.Get("/times", app.handleTimeIndex())
		r.Post("/times", app.handleTimeCreate())
		r.Get("/times/{id:[a-z0-9-]{36}}", app.handleTimeGet())
		r.Put("/times/{id:[a-z0-9-]{36}}", app.updateTime())
		r.Delete("/times/{id:[a-z0-9-]{36}}", app.deleteTime())
		r.Get("/times/search", app.handleTimeSearch())
	})

	// File server - Serve the compiled frontend app when the app is in production
	if app.Production {
		workDir, _ := os.Getwd()
		staticDir := http.Dir(filepath.Join(workDir, "frontend/dist"))

		app.Router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			rctx := chi.RouteContext(r.Context())
			pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
			fs := http.StripPrefix(pathPrefix, http.FileServer(staticDir))
			fs.ServeHTTP(w, r)
		})

	}
}

func (app *App) setupCors() {
	app.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func (app *App) serve() error {
	app.InfoLog.Printf("Starting server at http://%v:%v", app.Host, app.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(app.Port), app.Router); err != nil {
		return err
	}
	return nil
}
