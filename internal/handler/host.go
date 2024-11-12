package handler

import (
	"context"

	"github.com/ei-sugimoto/datapuppy/internal/handler/views"
	"github.com/ei-sugimoto/datapuppy/internal/usecase"
)

type HostHandler struct {
	HUsecase usecase.HostUsecase
}

func NewHostHandler(hu usecase.HostUsecase) *HostHandler {
	return &HostHandler{
		HUsecase: hu,
	}
}

func (h *HostHandler) CreateHost(ctx context.Context, name string, port int) (*views.ViewHost, *views.ViewError) {
	host, err := h.HUsecase.Create(ctx, name, port)
	if err != nil {
		viewErr := &views.ViewError{
			Message: err.Error.Error(),
			Code:    err.Code,
		}
		return nil, viewErr
	}
	viewHost := &views.ViewHost{
		Name: host.Name,
		Port: host.Port,
	}
	return viewHost, nil
}
