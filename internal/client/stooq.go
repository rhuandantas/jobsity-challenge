package client

import (
	"chat-jobsity/internal/client/kafka"
	"chat-jobsity/internal/config"
	"encoding/csv"
	"fmt"
	"net/http"
)

//go:generate mockgen -source=$GOFILE -package=mock_client -destination=../../test/mock/client/$GOFILE

type StooqClient interface {
	GetStockDetails(stockCode string) (string, error)
}

type StooqClientImpl struct {
	kafkaProducer kafka.Producer
	cfg           config.ConfigStore
}

func NewStooqClient(cfg config.ConfigStore, kafkaProducer kafka.Producer) StooqClient {
	return &StooqClientImpl{
		cfg:           cfg,
		kafkaProducer: kafkaProducer,
	}
}

func (cli StooqClientImpl) GetStockDetails(stockCode string) (string, error) {
	urlTemplate := fmt.Sprintf("%s", cli.cfg.Get("stooq.url"))
	url := fmt.Sprintf(urlTemplate, stockCode)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	reader := csv.NewReader(res.Body)
	data, err := reader.ReadAll()
	if err != nil {
		return "", err
	}
	msg := ""
	if len(data) < 2 || data[1][3] == "N/D" {
		msg = fmt.Sprintf("{\"message\":\"ChatBot no records for this stockCode %s\", \"error\":null}", stockCode)
	} else {
		msg = fmt.Sprintf("{\"message\":\"ChatBot %s quote is $%s per share\", \"error\":null}", stockCode, data[1][3])
	}

	go cli.kafkaProducer.SendMessage(msg)

	return "", nil
}
