// MIT License

// Copyright (c) 2023 Yamasaki Shotaro

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package lib

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/seipan/csql/query"
)

type CsvFile struct {
	path        string
	content     [][]string
	tableSchema map[string]string
}

func CsvFormatExec(cfg Config) error {
	return nil
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
	for i, v := range c.content[0] {
		c.tableSchema[v] = c.content[1][i]
	}
}

func (c *CsvFile) GetTableName() string {
	return c.content[0][0]
}

func (c *CsvFile) GetTableSchemaMap() map[string]string {
	return c.tableSchema
}

func (c *CsvFile) GetKeyValues() ([]query.KeyValues, error) {
	if c.content == nil {
		return nil, fmt.Errorf("csv file content is nil")
	}
	var kvs []query.KeyValues
	for i, v := range c.content {
		if i == 0 {
			continue
		}
		var kv query.KeyValues
		for j, vv := range v {
			if j == 0 {
				continue
			}
			k := query.KeyValue{
				Key:   c.content[0][j],
				Value: vv,
			}
			kv = append(kv, k)
		}
		kvs = append(kvs, kv)
	}

	return kvs, nil
}

func (c *CsvFile) CheckCsvFormat() error {
	if err := c.checkTableNames(); err != nil {
		return err
	}
	if err := c.checkLength(); err != nil {
		return err
	}
	if err := c.checkTableSchema(); err != nil {
		return err
	}
	if err := c.checkTableValue(); err != nil {
		return err
	}
	return nil
}

func (c *CsvFile) checkTableValue() error {
	if c.content == nil {
		return fmt.Errorf("csv file content is nil")
	}
	for i, v := range c.content {
		if i == 0 {
			continue
		}
		for j, vv := range v {
			if j == 0 {
				continue
			}
			if vv == "" && c.content[0][j] != "" {
				return fmt.Errorf("table value is empty at %d %d", i, j)
			}
		}
	}
	return nil
}

func (c *CsvFile) checkTableSchema() error {
	if c.content == nil {
		return fmt.Errorf("csv file content is nil")
	}
	for i, v := range c.content[0] {
		if i == 0 {
			continue
		}
		if v == "" && c.content[1][i] != "" {
			return fmt.Errorf("table schema is empty at %d", i)
		}
	}
	return nil
}

func (c *CsvFile) checkTableNames() error {
	if c.content == nil {
		return fmt.Errorf("csv file content is nil")
	}
	if c.content[0][0] == "" {
		return fmt.Errorf("table name is empty")
	}
	return nil
}

func (c *CsvFile) checkLength() error {
	if c.content == nil {
		return fmt.Errorf("csv file content is nil")
	}
	for i, v := range c.content {
		if i == 0 {
			continue
		}
		if len(v) != len(c.content[0]) {
			return fmt.Errorf("table length is not equal at %d", i)
		}
	}
	return nil
}

func NewCsvFile(path string) (*CsvFile, error) {
	c := &CsvFile{
		path:        path,
		tableSchema: make(map[string]string),
	}
	if err := c.SetContent(); err != nil {
		return nil, err
	}
	c.GetTableSchema()
	return c, nil
}
