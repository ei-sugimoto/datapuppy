package usecase

import (
	"context"
	"fmt"

	"github.com/ei-sugimoto/datapuppy/internal"
	"github.com/ei-sugimoto/datapuppy/internal/domain"
	"github.com/ei-sugimoto/datapuppy/internal/repo"
)

type HostUsecase interface {
	Create(ctx context.Context, name string, port int) (*domain.Host, *internal.CustomError)
}

type hostUsecase struct {
	hrepo repo.HostRepo
}

func NewHostUsecase(hrepo repo.HostRepo) HostUsecase {
	return &hostUsecase{
		hrepo: hrepo,
	}
}

func (u *hostUsecase) Create(ctx context.Context, name string, port int) (*domain.Host, *internal.CustomError) {
	host, err := domain.NewHost(name, port)
	if err != nil {
		CustomError := &internal.CustomError{
			Error: err,
			Code:  400,
		}
		return nil, CustomError
	}

	host, err = u.hrepo.Create(ctx, host)
	if err != nil {
		CustomError := &internal.CustomError{
			Error: fmt.Errorf("Failed to create a host: %w", err),
			Code:  500,
		}
		return nil, CustomError
	}
	return host, nil

}
