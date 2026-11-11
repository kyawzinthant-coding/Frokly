package handler

import (
	"context"
	"github.com/kyawzinthant-coding/froky/api"
)

func (h *Handler) GetStatus(ctx context.Context, request api.GetStatusRequestObject) (api.GetStatusResponseObject, error) {
	return api.GetStatus200JSONResponse{
		Status: "ok",
	}, nil

}
