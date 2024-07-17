package db

import (
	"github.com/go-pg/pg/v10"
)

type NewsRepo struct {
	*pg.DB
}

func NewNewsRepo(db *pg.DB) NewsRepo {
	return NewsRepo{DB: db}
}

// GetNewsByID возвращает News по id из бд
func (nr *NewsRepo) GetNewsByID(id int) (*News, error) {
	news := &News{ID: id}
	err := nr.Model(news).WherePK().Select()
	return news, err
}
