package mysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNewMySQLDB(t *testing.T) {
	tests := []struct {
		name string
		want *mySqlDb
	}{
		{name: "転けなければ良い"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewMySQLDB()
		})
	}
}
