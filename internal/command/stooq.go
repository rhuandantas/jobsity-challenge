package command

import "chat-jobsity/internal/client"

//go:generate mockgen -source=$GOFILE -package=mock_command -destination=../../test/mock/command/$GOFILE

type StockCommand struct {
	stooqCli client.StooqClient
}

func NewStooqCommand(stooqCli client.StooqClient) *StockCommand {
	return &StockCommand{
		stooqCli: stooqCli,
	}
}

func (sc *StockCommand) Run(param string) (string, error) {
	return sc.stooqCli.GetStockDetails(param)
}
