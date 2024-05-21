package mysql

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/domain/model"
	"github.com/pulse227/server-recruit-challenge-sample/domain/repository"
)

type singerDB struct {
	connection *mySqlDb
}

// 単純に挿入しようしとしているだけか
var _ repository.SingerRepository = (*singerDB)(nil)

func NewSingerDB(db *mySqlDb) *singerDB {
	return &singerDB{
		connection: db,
	}
}

func (r *singerDB) GetAll(ctx context.Context) ([]*model.Singer, error) {
	rows, err := r.connection.db.QueryContext(ctx, "SELECT * FROM music.singer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var singers []*model.Singer
	for rows.Next() {
		var singerID model.SingerID
		var name string
		err := rows.Scan(&singerID, &name)
		singers = append(singers, &model.Singer{ID: singerID, Name: name})
		if err != nil {
			return nil, err
		}
	}

	return singers, nil
}

func (r *singerDB) Get(ctx context.Context, id model.SingerID) (*model.Singer, error) {
	var singerID model.SingerID
	var name string
	row := r.connection.db.QueryRowContext(ctx, "SELECT * FROM music.singer WHERE ID = ?", id)
	err := row.Scan(&singerID, &name)
	if err != nil {
		return nil, err
	}
	return &model.Singer{ID: singerID, Name: name}, nil
}

func (r *singerDB) Add(ctx context.Context, singer *model.Singer) error {
	_, err := r.connection.db.ExecContext(ctx, "INSERT INTO music.singer (ID, Name) VALUES (?, ?)", singer.ID, singer.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *singerDB) Delete(ctx context.Context, id model.SingerID) error {
	_, err := r.connection.db.ExecContext(ctx, "DELETE FROM music.singer WHERE ID = ?", id)
	if err != nil {
		return err
	}
	return nil
}
