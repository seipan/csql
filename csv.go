package csql

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvFile struct {
	path        string
	content     [][]string
	tableSchema map[string]string
}

func (c *CsvFile) ReadCsvFile() ([][]string, error) {
	file, err := os.Open(c.path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv file: %w", err)
	}
	return rows, nil
}

func (c *CsvFile) SetContent() error {
	str, err := c.ReadCsvFile()
	if err != nil {
		return err
	}
	c.content = str
	return nil
}

func (c *CsvFile) GetTableSchema() {

}
