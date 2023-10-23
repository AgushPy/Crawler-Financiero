package utils

import (
	"encoding/csv"
	"log"
	"os"
)

var (
	dirFile string = "utils/sites.csv"
	sites   []string
)

func ReadCSV() []string {
	file, err := os.Open(dirFile)
	defer file.Close()
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for index, row := range rows {
		if index == 0 {
			continue
		}
		sites = append(sites, row[0])
	}
	return sites
}
