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
	"database/sql"
	"fmt"

	"github.com/seipan/csql/mariadb"
	"github.com/seipan/csql/mysql"
	"github.com/seipan/csql/postgresql"
	"github.com/seipan/csql/query"
	"github.com/seipan/csql/sqlite"
)

func InsertExec(config Config) error {
	db, err := newSQLDB(config)
	if err != nil {
		return err
	}
	defer db.Close()

	csv, err := NewCsvFile(config.Filepath)
	if err != nil {
		return err
	}

	kvs, err := csv.GetKeyValues()
	if err != nil {
		return err
	}

	totals := csv.GetSize()
	fmt.Printf("insert %d rows\n", totals)

	for i, kv := range kvs {
		p := (i + 1) * 100 / totals
		fmt.Printf("\rInserting: %s %2d%%", mark(p), p)
		insert, err := newSQLInserter(config.Type, kv, csv.GetTableName(), db)
		if err != nil {
			return err
		}
		err = insert.Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

func QueryExec(config Config) (string, error) {
	db, err := newSQLDB(config)
	if err != nil {
		return "", err
	}
	defer db.Close()

	csv, err := NewCsvFile(config.Filepath)
	if err != nil {
		return "", err
	}

	kvs, err := csv.GetKeyValues()
	if err != nil {
		return "", err
	}
	insert, err := newSQLInserter(config.Type, kvs[0], csv.GetTableName(), db)
	if err != nil {
		return "", err
	}
	str := insert.Query()
	return str, nil
}

func newSQLInserter(dbtype string, kv query.KeyValues, tablename string, db *sql.DB) (query.Inserter, error) {
	switch dbtype {
	case "mysql":
		return mysql.NewMySQLInserter(kv, tablename, db), nil
	case "sqlite":
		return sqlite.NewSQLiteInserter(kv, tablename, db), nil
	case "postgresql":
		return postgresql.NewPostgresSQLInserter(kv, tablename, db), nil
	case "mariadb":
		return mariadb.NewMariaDBInserter(kv, tablename, db), nil
	default:
		return nil, fmt.Errorf("invalid dbtype: %s", dbtype)
	}
}

func newSQLDB(config Config) (*sql.DB, error) {
	switch config.Type {
	case "mysql":
		return mysql.NewDB(config.DSN)
	case "sqlite":
		return sqlite.NewDB(config.DSN)
	case "postgresql":
		return postgresql.NewDB(config.DSN)
	case "mariadb":
		return mariadb.NewDB(config.DSN)
	default:
		return nil, fmt.Errorf("invalid dbtype: %s", config.Type)
	}
}

var marks = []string{"|", "/", "-", "\\"}

func mark(i int) string {
	return marks[i%4]
}
