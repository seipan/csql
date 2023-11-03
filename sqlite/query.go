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

package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/seipan/csql/query"
)

type SQLiteInserter struct {
	keys      query.KeyValues
	tableName string
	db        *sql.DB
}

func (i *SQLiteInserter) Query() string {
	placeholders := make([]string, 0, len(i.keys))
	keys := make([]string, 0, len(i.keys))

	for _, kv := range i.keys {
		keys = append(keys, kv.Key)
		placeholders = append(placeholders, "?")
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		i.tableName,
		strings.Join(keys, ", "),
		strings.Join(placeholders, ", "),
	)
	return query
}

func (i *SQLiteInserter) Insert() error {
	stmt, err := i.db.Prepare(i.Query())
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	values := make([]interface{}, 0, len(i.keys))
	for _, kv := range i.keys {
		values = append(values, kv.Value)
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func NewSQLiteInserter(kv query.KeyValues, tableName string, db *sql.DB) query.Inserter {
	return &SQLiteInserter{
		keys:      kv,
		tableName: tableName,
		db:        db,
	}
}
