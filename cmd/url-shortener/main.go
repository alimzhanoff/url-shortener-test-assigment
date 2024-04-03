package main

import (
	"github.com/alimzhanoff/url-shortener-test-assigment/internal/config"
	"github.com/alimzhanoff/url-shortener-test-assigment/internal/http-server/handlers/url/getByAlias"
	"github.com/alimzhanoff/url-shortener-test-assigment/internal/http-server/handlers/url/save"
	mwLogger "github.com/alimzhanoff/url-shortener-test-assigment/internal/http-server/middleware/logger"
	"github.com/alimzhanoff/url-shortener-test-assigment/internal/lib/logger/handlers/slogpretty"
	"github.com/alimzhanoff/url-shortener-test-assigment/internal/lib/logger/sl"
	"github.com/alimzhanoff/url-shortener-test-assigment/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//TODO: init config: cleanenv
	cfg := config.MustLoad()

	//TODO: init logger:slog
	log := setupLogger(cfg.Env)

	//TODO: init storage: sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
	//TODO: init router: chi, chi-render

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(mwLogger.New(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/save", save.New(log, storage))
			r.Get("/{alias}", getByAlias.New(log, storage))
		})
	})

	//TODO: init server:
	//srv := &http.Server{
	//	Addr:         cfg.Address,
	//	Handler:      r,
	//	ReadTimeout:  cfg.HTTPServer.Timeout,
	//	WriteTimeout: cfg.HTTPServer.Timeout,
	//	IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	//}

	http.ListenAndServe(":3000", r)
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
