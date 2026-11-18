package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/config"
	"github.com/kyawzinthant-coding/froky/internal/handler"
	authmw "github.com/kyawzinthant-coding/froky/internal/middleware"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"net/http"
	"strings"
)

type HttpServer struct {
	cfg    *config.AppConfig
	router http.Handler
}

func NewHttpServer(
	h *handler.Handler,
	cfg *config.AppConfig,
) *HttpServer {

	log.Debug().Msg("Starting HTTP server")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// CORS middlewares
	allowedOrigins := make([]string, 0)
	for _, o := range strings.Split(cfg.CorsConfig.AllowOrigins, ",") {
		if trimmed := strings.TrimSpace(o); trimmed != "" {
			allowedOrigins = append(allowedOrigins, trimmed)
		}
	}

	allowMethods := make([]string, 0)
	for _, m := range strings.Split(cfg.CorsConfig.AllowMethods, ",") {
		if trimmed := strings.TrimSpace(m); trimmed != "" {
			allowMethods = append(allowMethods, trimmed)
		}
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   allowMethods,
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:           300,
	}))

	publicRoutes := map[string]bool{
		"GET /v1/status":        true,
		"POST /v1/vendor/login": true,
	}

	r.Use(requireVendorAuth(publicRoutes, []byte(cfg.JWT.Secret)))

	strict := api.NewStrictHandler(h, nil)

	mux := api.HandlerFromMux(strict, r)

	return &HttpServer{
		cfg:    cfg,
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
	return nil
}

func requireVendorAuth(public map[string]bool, secret []byte) func(http.Handler) http.Handler {
	vendorAuth := authmw.VendorJwtAuth(secret)
	return func(next http.Handler) http.Handler {
		protected := vendorAuth(next)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if public[r.Method+" "+r.URL.Path] {
				next.ServeHTTP(w, r)
				return
			}
			protected.ServeHTTP(w, r)
		})
	}
}
