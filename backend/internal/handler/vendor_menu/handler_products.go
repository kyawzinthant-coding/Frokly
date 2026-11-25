package vendor_menu

import (
	"context"
	"errors"

	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/service/menu"
	"github.com/rs/zerolog/log"
)

func (h *VendorMenuHandler) ListVendorProducts(ctx context.Context, _ api.ListVendorProductsRequestObject) (api.ListVendorProductsResponseObject, error) {
	vendorID, ok := vendorIDFromContext(ctx)
	if !ok {
		return api.ListVendorProducts401JSONResponse{UnauthorizedJSONResponse: unauthorized()}, nil
	}

	products, err := h.menuService.ListProducts(ctx, vendorID)
	if err != nil {
		log.Error().Err(err).Int64("vendor_id", vendorID).Msg("list products failed")
		return nil, err
	}

	resp := make(api.ListVendorProducts200JSONResponse, 0, len(products))
	for _, p := range products {
		resp = append(resp, toAPIProduct(p))
	}
	return resp, nil
}

func (h *VendorMenuHandler) GetVendorProduct(ctx context.Context, req api.GetVendorProductRequestObject) (api.GetVendorProductResponseObject, error) {
	vendorID, ok := vendorIDFromContext(ctx)
	if !ok {
		return api.GetVendorProduct401JSONResponse{UnauthorizedJSONResponse: unauthorized()}, nil
	}

	product, err := h.menuService.GetProduct(ctx, vendorID, req.ProductId)
	switch {
	case errors.Is(err, menu.ErrProductNotFound):
		return api.GetVendorProduct404JSONResponse{ProductNotFoundJSONResponse: api.ProductNotFoundJSONResponse{Error: msgProductNotFound}}, nil
	case err != nil:
		log.Error().Err(err).Int64("vendor_id", vendorID).Int64("product_id", req.ProductId).Msg("get product failed")
		return nil, err
	}

	return api.GetVendorProduct200JSONResponse(toAPIProduct(*product)), nil
}

func (h *VendorMenuHandler) CreateVendorProduct(ctx context.Context, req api.CreateVendorProductRequestObject) (api.CreateVendorProductResponseObject, error) {
	vendorID, ok := vendorIDFromContext(ctx)
	if !ok {
		return api.CreateVendorProduct401JSONResponse{UnauthorizedJSONResponse: unauthorized()}, nil
	}

	product, err := h.menuService.CreateProduct(ctx, vendorID, toProductInput(req.Body))
	switch {
	case errors.Is(err, menu.ErrInvalidInput):
		return api.CreateVendorProduct400JSONResponse{BadRequestJSONResponse: api.BadRequestJSONResponse{Error: err.Error()}}, nil
	case errors.Is(err, menu.ErrCategoryNotFound):
		return api.CreateVendorProduct404JSONResponse{Error: msgCategoryNotFound}, nil
	case errors.Is(err, menu.ErrNameTaken):
		return api.CreateVendorProduct409JSONResponse{ProductNameTakenJSONResponse: api.ProductNameTakenJSONResponse{Error: msgProductNameTaken}}, nil
	case err != nil:
		log.Error().Err(err).Int64("vendor_id", vendorID).Msg("create product failed")
		return nil, err
	}

	return api.CreateVendorProduct201JSONResponse(toAPIProduct(*product)), nil
}

func (h *VendorMenuHandler) UpdateVendorProduct(ctx context.Context, req api.UpdateVendorProductRequestObject) (api.UpdateVendorProductResponseObject, error) {
	vendorID, ok := vendorIDFromContext(ctx)
	if !ok {
		return api.UpdateVendorProduct401JSONResponse{UnauthorizedJSONResponse: unauthorized()}, nil
	}

	product, err := h.menuService.UpdateProduct(ctx, vendorID, req.ProductId, toProductInput(req.Body))
	switch {
	case errors.Is(err, menu.ErrInvalidInput):
		return api.UpdateVendorProduct400JSONResponse{BadRequestJSONResponse: api.BadRequestJSONResponse{Error: err.Error()}}, nil
	case errors.Is(err, menu.ErrProductNotFound):
		return api.UpdateVendorProduct404JSONResponse{Error: msgProductNotFound}, nil
	case errors.Is(err, menu.ErrCategoryNotFound):
		return api.UpdateVendorProduct404JSONResponse{Error: msgCategoryNotFound}, nil
	case errors.Is(err, menu.ErrNameTaken):
		return api.UpdateVendorProduct409JSONResponse{ProductNameTakenJSONResponse: api.ProductNameTakenJSONResponse{Error: msgProductNameTaken}}, nil
	case err != nil:
		log.Error().Err(err).Int64("vendor_id", vendorID).Int64("product_id", req.ProductId).Msg("update product failed")
		return nil, err
	}

	return api.UpdateVendorProduct200JSONResponse(toAPIProduct(*product)), nil
}

func (h *VendorMenuHandler) DeleteVendorProduct(ctx context.Context, req api.DeleteVendorProductRequestObject) (api.DeleteVendorProductResponseObject, error) {
	vendorID, ok := vendorIDFromContext(ctx)
	if !ok {
		return api.DeleteVendorProduct401JSONResponse{UnauthorizedJSONResponse: unauthorized()}, nil
	}

	err := h.menuService.ArchiveProduct(ctx, vendorID, req.ProductId)
	switch {
	case errors.Is(err, menu.ErrProductNotFound):
		return api.DeleteVendorProduct404JSONResponse{ProductNotFoundJSONResponse: api.ProductNotFoundJSONResponse{Error: msgProductNotFound}}, nil
	case err != nil:
		log.Error().Err(err).Int64("vendor_id", vendorID).Int64("product_id", req.ProductId).Msg("delete product failed")
		return nil, err
	}

	return api.DeleteVendorProduct204Response{}, nil
}
