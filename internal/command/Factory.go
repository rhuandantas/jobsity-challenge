package command

import (
	"chat-jobsity/internal/client"
	"errors"
	"strings"
)

var (
	Stock = "/stock"
)

type Command interface {
	Run(param string) (string, error)
}

type StockCommand struct {
}

func GetCommand(name string) (Command, error) {
	switch strings.ToLower(name) {
	case Stock:
		return StockCommand{}, nil
	}

	return nil, errors.New("command is invalid")
}

func (sc StockCommand) Run(param string) (string, error) {
	stooqCli := client.StooqClient{}
	return stooqCli.GetStockDetails(param)
}
