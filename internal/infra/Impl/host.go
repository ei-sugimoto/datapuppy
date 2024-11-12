package impl

import (
	"database/sql"
	"fmt"

	"github.com/ei-sugimoto/datapuppy/internal/domain"
	"github.com/ei-sugimoto/datapuppy/internal/repo"
)

type HostImpl struct {
	db *sql.DB
}

func NewHostImpl(db *sql.DB) repo.HostRepo {
	return &HostImpl{db: db}

}

func (h *HostImpl) Create(host *domain.Host) (*domain.Host, error) {
	res, err := h.db.Exec("INSERT INTO hosts (name, port) VALUES (?, ?)", host.Name, host.Port)
	if err != nil {
		return nil, fmt.Errorf("Failed to insert a host: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("Failed to get the last inserted ID: %w", err)
	}
	host.ID = int(id)

	return host, nil
}
