package client

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
)

type StooqClient struct {
}

type Stooq struct {
	Symbol string
	Date   string
	Time   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

func (cli StooqClient) GetStockDetails(stockCode string) (string, error) {
	url := fmt.Sprintf("https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv", stockCode)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	reader := csv.NewReader(res.Body)
	data, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		return "", errors.New("no records for this stockCode")
	}

	msg := fmt.Sprintf("%s quote is $%s per share", stockCode, data[1][3])
	return msg, nil
}
