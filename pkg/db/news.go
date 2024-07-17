package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
)

type NewsRepo struct {
	*pg.DB
}

func NewNewsRepo(db *pg.DB) NewsRepo {
	return NewsRepo{DB: db}
}

// NewsByID возвращает News по id из бд
func (nr *NewsRepo) NewsByID(id int) (*News, error) {
	news := &News{ID: id}
	err := nr.Model(news).WherePK().Select()
	if err != nil {
		return nil, fmt.Errorf("[NewsByID]: %s\n", err)
	}
	return news, nil
}

// NewsByTagID возвращает все новости с заданным тегом
func (nr *NewsRepo) NewsByTagID(tagID int) ([]News, error) {
	var news []News
	err := nr.Model(&news).Where("? = ANY (\"tagIds\")", tagID).Select()
	if err != nil {
		return nil, fmt.Errorf("[NewsByTagID]: %s\n", err)
	}
	return news, nil
}

// NewsByCategoryID возвращает все новости в заданной категории
func (nr *NewsRepo) NewsByCategoryID(categoryID int) ([]News, error) {
	var news []News
	err := nr.Model(&news).Where("\"categoryId\" = ?", categoryID).Select()
	if err != nil {
		return nil, fmt.Errorf("[NewsByCategoryID]: %s\n", err)
	}
	return news, nil
}

// NewsByTagIDWithLimit возвращает ограниченное количество новостей с заданным тегом
func (nr *NewsRepo) NewsByTagIDWithLimit(tagID, limit int) ([]News, error) {
	var news []News
	err := nr.Model(&news).Where("? = ANY (\"tagIds\")", tagID).Limit(limit).Select()
	if err != nil {
		return nil, fmt.Errorf("[NewsByTagIDWithLimit]: %s\n", err)
	}
	return news, nil
}

// NewsByCategoryIDWithLimit возвращает ограниченное количество новостей в заданной категории
func (nr *NewsRepo) NewsByCategoryIDWithLimit(categoryID, limit int) ([]News, error) {
	var news []News
	err := nr.Model(&news).Where("\"categoryId\" = ?", categoryID).Limit(limit).Select()
	if err != nil {
		return nil, fmt.Errorf("[NewsByCategoryIDWithLimit]: %s\n", err)
	}
	return news, nil
}

// NewsByTagIDWithPagination возвращает новости с заданным тегом с пагинацией
func (nr *NewsRepo) NewsByTagIDWithPagination(tagID, page, pageSize int) ([]News, error) {
	var news []News
	offset := (page - 1) * pageSize
	err := nr.Model(&news).
		Where("? = ANY (\"tagIds\")", tagID).
		Offset(offset).
		Limit(pageSize).
		Select()
	if err != nil {
		return nil, fmt.Errorf("[NewsByTagIDWithPagination]: %s\n", err)
	}
	return news, nil
}

// NewsByCategoryIDWithPagination возвращает новости в заданной категории с пагинацией
func (nr *NewsRepo) NewsByCategoryIDWithPagination(categoryID, page, pageSize int) ([]News, error) {
	var news []News
	offset := (page - 1) * pageSize
	err := nr.Model(&news).
		Where("\"categoryId\" = ?", categoryID).
		Offset(offset).
		Limit(pageSize).
		Select()
	if err != nil {
		return nil, fmt.Errorf("[NewsByCategoryIDWithPagination]: %s\n", err)
	}
	return news, nil
}

// CategoryByID возвращает Category по id из бд
func (nr *NewsRepo) CategoryByID(id int) (*Category, error) {
	category := &Category{ID: id}
	err := nr.Model(category).WherePK().Select()
	if err != nil {
		return nil, fmt.Errorf("[CategoryByID]: %s\n", err)
	}
	return category, nil
}

// TagByID возвращает Tag по id из бд
func (nr *NewsRepo) TagByID(id int) (*Tag, error) {
	tag := &Tag{ID: id}
	err := nr.Model(tag).WherePK().Select()
	if err != nil {
		return nil, fmt.Errorf("[TagByID]: %s\n", err)
	}
	return tag, nil
}
