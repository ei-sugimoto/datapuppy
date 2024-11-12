package usecase

import (
	"fmt"

	"github.com/ei-sugimoto/datapuppy/internal/domain"
	"github.com/ei-sugimoto/datapuppy/internal/repo"
)

type HostUsecase interface {
	Create(name string, port int) (*domain.Host, error)
}

type hostUsecase struct {
	hrepo repo.HostRepo
}

func NewHostUsecase(hrepo repo.HostRepo) HostUsecase {
	return &hostUsecase{
		hrepo: hrepo,
	}
}

func (u *hostUsecase) Create(name string, port int) (*domain.Host, error) {
	host, err := domain.NewHost(name, port)
	if err != nil {
		return nil, fmt.Errorf("failed to create host. when usecase: %w", err)
	}
	host, err = u.hrepo.Create(host)
	if err != nil {
		return nil, fmt.Errorf("failed to create host. when usecase: %w", err)
	}
	return host, nil

}
