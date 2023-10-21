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

package postgresql

import (
	"testing"

	"github.com/seipan/csql/lib"
	"github.com/stretchr/testify/assert"
)

func TestPostgresSQLInserter_Query(t *testing.T) {
	tests := []struct {
		name      string
		keys      []lib.KeyValue
		tableName string
		expected  string
	}{
		{
			name: "single key-value pair",
			keys: []lib.KeyValue{
				{Key: "name", Value: "John"},
			},
			tableName: "users",
			expected:  "INSERT INTO users (name) VALUES ($1)",
		},
		{
			name: "multiple key-value pairs",
			keys: []lib.KeyValue{
				{Key: "name", Value: "John"},
				{Key: "age", Value: "30"},
			},
			tableName: "users",
			expected:  "INSERT INTO users (name, age) VALUES ($1, $2)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inserter := &PostgresSQLInserter{
				keys:      tt.keys,
				tableName: tt.tableName,
			}

			assert.Equal(t, tt.expected, inserter.Query())
		})
	}
}
