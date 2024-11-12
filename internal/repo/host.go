package repo

import "github.com/ei-sugimoto/datapuppy/internal/domain"

type HostRepo interface {
	Create(*domain.Host) (*domain.Host, error)
}
