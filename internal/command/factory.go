package command

import (
	"errors"
	"fmt"
	"strings"
)

//go:generate mockgen -source=$GOFILE -package=mock_command -destination=../../test/mock/command/$GOFILE

var (
	Stock = "/stock"
)

type Manager interface {
	GetCommand(name string) (Command, error)
}

type ManagerCommand struct {
	stockCommand Command
}

func NewCommandManager(stockCommand Command) Manager {
	return &ManagerCommand{
		stockCommand: stockCommand,
	}
}

func (m *ManagerCommand) GetCommand(name string) (Command, error) {
	switch strings.ToLower(name) {
	case Stock:
		return m.stockCommand, nil
	}

	return nil, errors.New(fmt.Sprintf("ChatBot command %s is invalid", name))
}
