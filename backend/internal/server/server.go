package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/config"
	"github.com/kyawzinthant-coding/froky/internal/handler"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"net/http"
)

type HttpServer struct {
	cfg    *config.AppConfig
	db     *sql.DB
	router http.Handler
}

func NewHttpServer(
	cfg *config.AppConfig,
	db *sql.DB,
) *HttpServer {

	log.Debug().Msg("Starting HTTP server")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	h := handler.NewHandler()
	strict := api.NewStrictHandler(h, nil)

	mux := api.HandlerFromMux(strict, r)

	return &HttpServer{
		cfg:    cfg,
		db:     db,
		router: mux,
	}
}

func (h *HttpServer) Start(lc fx.Lifecycle) error {

	addr := fmt.Sprintf(":%d", h.cfg.Http.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: h.router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info().Str("addr", addr).Msg("Starting HTTP server")
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatal().Err(err).Msg("Failed to start HTTP server")
				}
			}()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			log.Info().Str("addr", addr).Msg("Stopping HTTP server")
			return srv.Shutdown(ctx)
		},
	})

	log.Info().Msg("Starting HTTP server")
	return srv.ListenAndServe()

}
