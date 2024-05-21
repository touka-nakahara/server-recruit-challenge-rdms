package mysql

import (
	"context"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/pulse227/server-recruit-challenge-sample/domain/model"
)

func Test_singerDB_GetAll(t *testing.T) {
	db := NewMySQLDB()
	tc := []struct {
		input    string
		expected []*model.Singer
	}{
		{input: "", expected: []*model.Singer{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bella"},
			{ID: 3, Name: "Chris"},
			{ID: 4, Name: "Daisy"},
			{ID: 5, Name: "Ellen"},
		}},
	}
	for _, tt := range tc {
		t.Run(tt.input, func(t *testing.T) {
			singerDB := NewSingerDB(db)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			got, err := singerDB.GetAll(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("singerDB.GetAll() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_singerDB_Get(t *testing.T) {
	db := NewMySQLDB()
	tc := []struct {
		input    string
		expected *model.Singer
	}{
		{input: "", expected: &model.Singer{ID: 4, Name: "Daisy"}},
	}
	for _, tt := range tc {
		t.Run(tt.input, func(t *testing.T) {
			singerDB := NewSingerDB(db)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			got, err := singerDB.Get(ctx, 4)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("singerDB.GetAll() = %#v, want %#v", got, tt.expected)
			}
		})
	}
}

// TODO トランザクション入らんわ おわた...
func Test_singerDB_POST(t *testing.T) {
	connection := NewMySQLDB()
	tc := []struct {
		input    *model.Singer
		expected *model.Singer
	}{
		{
			input:    &model.Singer{ID: 6, Name: "Floria"},
			expected: &model.Singer{ID: 6, Name: "Floria"},
		},
	}
	for _, tt := range tc {
		t.Run(tt.input.Name, func(t *testing.T) {
			singerDB := NewSingerDB(connection)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err := singerDB.Add(ctx, tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got, err := singerDB.Get(ctx, tt.input.ID)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("singerDB.GetAll() = %#v, want %#v", got, tt.expected)
			}

			// ロールバック
		})
	}
}

func Test_singerDB_DELETE(t *testing.T) {
	connection := NewMySQLDB()
	tc := []struct {
		input    int
		expected *model.Singer
	}{
		{
			input:    1,
			expected: nil,
		},
	}
	for _, tt := range tc {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			singerDB := NewSingerDB(connection)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err := singerDB.Delete(ctx, model.SingerID(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			got, err := singerDB.Get(ctx, model.SingerID(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("singerDB.GetAll() = %#v, want %#v", got, tt.expected)
			}

			// ロールバック
		})
	}
}
