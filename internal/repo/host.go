package repo

import (
	"context"

	"github.com/ei-sugimoto/datapuppy/internal/domain"
)

type HostRepo interface {
	Create(context.Context, *domain.Host) (*domain.Host, error)
}
