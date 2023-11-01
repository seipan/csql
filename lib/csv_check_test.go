package lib

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckCsvFormatSuccess(t *testing.T) {
	tests := []struct {
		name    string
		csvfile CsvFile
	}{
		{
			name: "test01(success)",
			csvfile: CsvFile{
				path: "../testdata/csv/test01.csv",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.csvfile.SetContent()
			assert.NoError(t, err)
			err = tt.csvfile.CheckCsvFormat()
			assert.NoError(t, err)
		})
	}
}

func TestCheckCsvFormatFailed(t *testing.T) {
	tests := []struct {
		name    string
		csvfile CsvFile
		wantErr error
	}{
		{
			name: "test02(table empty)",
			csvfile: CsvFile{
				path: "../testdata/csv/test02.csv",
			},
			wantErr: fmt.Errorf("table name is empty"),
		},
		{
			name: "test03(tabel schema empty))",
			csvfile: CsvFile{
				path: "../testdata/csv/test03.csv",
			},
			wantErr: fmt.Errorf("table schema is empty at "),
		},
		{
			name: "test04",
			csvfile: CsvFile{
				path: "../testdata/csv/test04.csv",
			},
			wantErr: fmt.Errorf("table value is empty at "),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.csvfile.SetContent()
			assert.NoError(t, err)
			err = tt.csvfile.CheckCsvFormat()
			if !strings.Contains(err.Error(), tt.wantErr.Error()) {
				t.Errorf("CheckCsvFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCsvFormatExec(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr error
	}{
		{
			name: "test01(success)",
			path: "../testdata/csv/test01.csv",
		},
		{
			name:    "test02(table empty)",
			path:    "../testdata/csv/test02.csv",
			wantErr: fmt.Errorf("table name is empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{
				Filepath: tt.path,
			}
			err := CsvFormatExec(cfg)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else if !strings.Contains(err.Error(), tt.wantErr.Error()) {
				t.Errorf("CheckCsvFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
