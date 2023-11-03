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
	"testing"

	"github.com/seipan/csql/query"
	"github.com/stretchr/testify/assert"
)

func TestSQLiteInserter_Query(t *testing.T) {
	tests := []struct {
		name      string
		keys      []query.KeyValue
		tableName string
		expected  string
	}{
		{
			name: "single key-value pair",
			keys: []query.KeyValue{
				{Key: "name", Value: "John"},
			},
			tableName: "users",
			expected:  "INSERT INTO users (name) VALUES (?);",
		},
		{
			name: "multiple key-value pairs",
			keys: []query.KeyValue{
				{Key: "name", Value: "John"},
				{Key: "age", Value: "30"},
			},
			tableName: "users",
			expected:  "INSERT INTO users (name, age) VALUES (?, ?);",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inserter := &SQLiteInserter{
				keys:      tt.keys,
				tableName: tt.tableName,
			}

			result := inserter.Query()
			assert.Equal(t, tt.expected, result)
		})
	}
}
