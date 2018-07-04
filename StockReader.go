package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Stock struct {
	Date     string  `json:"date"`
	Open     float64 `json:"open"`
	High     float64 `json:"high:`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	AdjClose float64 `json:"adj"`
	Volume   int     `json:"volume"`
}

func main() {
	csvFile, _ := os.Open("stock.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var stock []Stock
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)

		}

		Open, err := strconv.ParseFloat(line[1], 64)
		High, err := strconv.ParseFloat(line[2], 64)
		Low, err := strconv.ParseFloat(line[3], 64)
		Close, err := strconv.ParseFloat(line[4], 64)
		Adj, err := strconv.ParseFloat(line[5], 64)
		Volume, err := strconv.Atoi(line[6])

		if err == nil {

		}
		stock = append(stock, Stock{
			Date:     line[0],
			Open:     Open,
			High:     High,
			Low:      Low,
			Close:    Close,
			AdjClose: Adj,
			Volume:   Volume,
		})
	}

	stockJSON, _ := json.Marshal(stock)
	fmt.Println(string(stockJSON))
}
