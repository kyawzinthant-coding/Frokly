package handler

import (
	"context"
	"time"

	"github.com/kyawzinthant-coding/froky/api"
)

func (h *Handler) GetStatus(ctx context.Context, request api.GetStatusRequestObject) (api.GetStatusResponseObject, error) {
	pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := h.db.PingContext(pingCtx); err != nil {
		return api.GetStatus503JSONResponse{
			Status:   "unhealthy",
			Database: "down",
		}, nil
	}

	return api.GetStatus200JSONResponse{
		Status:   "ok",
		Database: "ok",
	}, nil
}
