package models

import (
	"log"
	"strconv"
	"strings"
)

type TypeActifs string

type Information struct {
	TypeActifs  TypeActifs
	Price       uint64
	Amount      uint64
	Transmitter string
}

type FuturesActifs struct {
	Symbol        string
	Name          string
	LastPrice     float64
	MarketTime    string
	Change        string
	PercentChange string
	Volume        string
	// Volume        float64
	TotalValue string
}

func ParserStringToFloat64(value string) float64 {
	str := strings.Replace(value, ",", "", -1)
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func NewFuture() *FuturesActifs {
	return &FuturesActifs{}
}
