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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryExec(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantStr string
	}{
		{
			name: "mysql query",
			config: Config{
				Type:     "mysql",
				DSN:      "",
				Filepath: "../testdata/csv/test01.csv",
			},
			wantStr: "INSERT INTO user (name, id, email) VALUES (?, ?, ?)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, err := QueryExec(tt.config)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStr, str)
		})
	}
}

func TestNewSQLInserter(t *testing.T) {
	db := &sql.DB{}
	tests := []struct {
		name      string
		dbtype    string
		tablename string
		wantErr   bool
	}{
		{
			name:      "mysql inserter",
			dbtype:    "mysql",
			tablename: "testTable",
			wantErr:   false,
		},
		{
			name:      "sqlite inserter",
			dbtype:    "sqlite",
			tablename: "testTable",
			wantErr:   false,
		},
		{
			name:      "postgresql inserter",
			dbtype:    "postgresql",
			tablename: "testTable",
			wantErr:   false,
		},
		{
			name:      "mariadb inserter",
			dbtype:    "mariadb",
			tablename: "testTable",
			wantErr:   false,
		},
		{
			name:      "invalid dbtype",
			dbtype:    "invalid",
			tablename: "testTable",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inserter, err := newSQLInserter(tt.dbtype, nil, tt.tablename, db)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, inserter)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, inserter)
			}
		})
	}
}

func TestNewSQLDB(t *testing.T) {
	tests := []struct {
		name     string
		config   Config
		wantErr  bool
		expected string
	}{
		{
			name:     "mysql db",
			config:   Config{Type: "mysql"},
			wantErr:  false,
			expected: "*sql.DB",
		},
		{
			name:     "sqlite db",
			config:   Config{Type: "sqlite"},
			wantErr:  false,
			expected: "*sql.DB",
		},
		{
			name:     "postgresql db",
			config:   Config{Type: "postgresql"},
			wantErr:  false,
			expected: "*sql.DB",
		},
		{
			name:     "mariadb db",
			config:   Config{Type: "mariadb"},
			wantErr:  false,
			expected: "*sql.DB",
		},
		{
			name:    "invalid db",
			config:  Config{Type: "invalid"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := newSQLDB(tt.config)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, db)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, db)
				assert.Equal(t, tt.expected, fmt.Sprintf("%T", db))
			}
		})
	}
}
