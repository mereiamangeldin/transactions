package handler

import (
	"github.com/mereiamangeldin/transaction/config"
	"github.com/mereiamangeldin/transaction/service"
)

type Manager struct {
	srv *service.Manager
}

func NewManager(conf *config.Config, srv *service.Manager) *Manager {
	return &Manager{srv: srv}
}
