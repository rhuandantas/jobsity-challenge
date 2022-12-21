package command

import (
	"errors"
	"fmt"
	"strings"
)

var (
	Stock = "/stock"
)

type Command interface {
	Run(param string) (string, error)
}

type Manager struct {
	stockCommand *StockCommand
}

func NewCommandManager(stockCommand *StockCommand) *Manager {
	return &Manager{
		stockCommand: stockCommand,
	}
}

func (m *Manager) GetCommand(name string) (Command, error) {
	switch strings.ToLower(name) {
	case Stock:
		return m.stockCommand, nil
	}

	return nil, errors.New(fmt.Sprintf("ChatBot command %s is invalid", name))
}
