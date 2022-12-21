package command

import "chat-jobsity/internal/client"

type StockCommand struct {
	stooqCli *client.StooqClient
}

func NewStooqCommand(stooqCli *client.StooqClient) *StockCommand {
	return &StockCommand{
		stooqCli: stooqCli,
	}
}

func (sc *StockCommand) Run(param string) (string, error) {
	return sc.stooqCli.GetStockDetails(param)
}
