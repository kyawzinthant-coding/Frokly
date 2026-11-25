package vendor_menu

import (
	"context"
	"errors"

	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/service/menu"
	"github.com/rs/zerolog/log"
)

func (h *VendorMenuHandler) ListVendorCategories(ctx context.Context, _ api.ListVendorCategoriesRequestObject) (api.ListVendorCategoriesResponseObject, error) {
	vendorID, ok := vendorIDFromContext(ctx)
	if !ok {
		return api.ListVendorCategories401JSONResponse{UnauthorizedJSONResponse: unauthorized()}, nil
	}

	categories, err := h.menuService.ListCategories(ctx, vendorID)
	if err != nil {
		log.Error().Err(err).Int64("vendor_id", vendorID).Msg("list categories failed")
		return nil, err
	}

	resp := make(api.ListVendorCategories200JSONResponse, 0, len(categories))
	for _, c := range categories {
		resp = append(resp, toAPICategory(c))
	}
	return resp, nil
}

func (h *VendorMenuHandler) GetVendorCategory(ctx context.Context, req api.GetVendorCategoryRequestObject) (api.GetVendorCategoryResponseObject, error) {
	vendorID, ok := vendorIDFromContext(ctx)
	if !ok {
		return api.GetVendorCategory401JSONResponse{UnauthorizedJSONResponse: unauthorized()}, nil
	}

	category, err := h.menuService.GetCategory(ctx, vendorID, req.CategoryId)
	switch {
	case errors.Is(err, menu.ErrCategoryNotFound):
		return api.GetVendorCategory404JSONResponse{CategoryNotFoundJSONResponse: api.CategoryNotFoundJSONResponse{Error: msgCategoryNotFound}}, nil
	case err != nil:
		log.Error().Err(err).Int64("vendor_id", vendorID).Int64("category_id", req.CategoryId).Msg("get category failed")
		return nil, err
	}

	return api.GetVendorCategory200JSONResponse(toAPICategory(*category)), nil
}

func (h *VendorMenuHandler) CreateVendorCategory(ctx context.Context, req api.CreateVendorCategoryRequestObject) (api.CreateVendorCategoryResponseObject, error) {
	vendorID, ok := vendorIDFromContext(ctx)
	if !ok {
		return api.CreateVendorCategory401JSONResponse{UnauthorizedJSONResponse: unauthorized()}, nil
	}

	category, err := h.menuService.CreateCategory(ctx, vendorID, req.Body.Name)
	switch {
	case errors.Is(err, menu.ErrInvalidInput):
		return api.CreateVendorCategory400JSONResponse{BadRequestJSONResponse: api.BadRequestJSONResponse{Error: err.Error()}}, nil
	case errors.Is(err, menu.ErrNameTaken):
		return api.CreateVendorCategory409JSONResponse{CategoryNameTakenJSONResponse: api.CategoryNameTakenJSONResponse{Error: msgCategoryNameTaken}}, nil
	case err != nil:
		log.Error().Err(err).Int64("vendor_id", vendorID).Msg("create category failed")
		return nil, err
	}

	return api.CreateVendorCategory201JSONResponse(toAPICategory(*category)), nil
}

func (h *VendorMenuHandler) UpdateVendorCategory(ctx context.Context, req api.UpdateVendorCategoryRequestObject) (api.UpdateVendorCategoryResponseObject, error) {
	vendorID, ok := vendorIDFromContext(ctx)
	if !ok {
		return api.UpdateVendorCategory401JSONResponse{UnauthorizedJSONResponse: unauthorized()}, nil
	}

	category, err := h.menuService.UpdateCategory(ctx, vendorID, req.CategoryId, req.Body.Name)
	switch {
	case errors.Is(err, menu.ErrInvalidInput):
		return api.UpdateVendorCategory400JSONResponse{BadRequestJSONResponse: api.BadRequestJSONResponse{Error: err.Error()}}, nil
	case errors.Is(err, menu.ErrCategoryNotFound):
		return api.UpdateVendorCategory404JSONResponse{CategoryNotFoundJSONResponse: api.CategoryNotFoundJSONResponse{Error: msgCategoryNotFound}}, nil
	case errors.Is(err, menu.ErrNameTaken):
		return api.UpdateVendorCategory409JSONResponse{CategoryNameTakenJSONResponse: api.CategoryNameTakenJSONResponse{Error: msgCategoryNameTaken}}, nil
	case err != nil:
		log.Error().Err(err).Int64("vendor_id", vendorID).Int64("category_id", req.CategoryId).Msg("update category failed")
		return nil, err
	}

	return api.UpdateVendorCategory200JSONResponse(toAPICategory(*category)), nil
}

func (h *VendorMenuHandler) DeleteVendorCategory(ctx context.Context, req api.DeleteVendorCategoryRequestObject) (api.DeleteVendorCategoryResponseObject, error) {
	vendorID, ok := vendorIDFromContext(ctx)
	if !ok {
		return api.DeleteVendorCategory401JSONResponse{UnauthorizedJSONResponse: unauthorized()}, nil
	}

	err := h.menuService.ArchiveCategory(ctx, vendorID, req.CategoryId)
	switch {
	case errors.Is(err, menu.ErrCategoryNotFound):
		return api.DeleteVendorCategory404JSONResponse{CategoryNotFoundJSONResponse: api.CategoryNotFoundJSONResponse{Error: msgCategoryNotFound}}, nil
	case errors.Is(err, menu.ErrCategoryNotEmpty):
		return api.DeleteVendorCategory409JSONResponse{CategoryNotEmptyJSONResponse: api.CategoryNotEmptyJSONResponse{Error: msgCategoryNotEmpty}}, nil
	case err != nil:
		log.Error().Err(err).Int64("vendor_id", vendorID).Int64("category_id", req.CategoryId).Msg("delete category failed")
		return nil, err
	}

	return api.DeleteVendorCategory204Response{}, nil
}
