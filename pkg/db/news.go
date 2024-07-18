package db

import (
	"errors"
	"github.com/go-pg/pg/v10"
)

type NewsRepo struct {
	*pg.DB
	QB QueryBuilder
}

const (
	StatusEnabled = iota + 1
	StatusDisabled
	StatusDeleted
)

func NewNewsRepo(db *pg.DB) NewsRepo {
	return NewsRepo{DB: db, QB: *NewQueryBuilder()}
}

// NewsByID возвращает News по id из бд
func (nr *NewsRepo) NewsByID(id int) (*News, error) {
	news := &News{ID: id}
	err := nr.Model(news).Where("\"statusId\" != ?", StatusDeleted).WherePK().Select()
	if errors.Is(err, pg.ErrNoRows) {
		err = errors.New("no news were found with this newsID")
	}
	return news, err
}

// NewsWithFilters возвращает все новости с необходимыми фильтрами
func (nr *NewsRepo) NewsWithFilters(qb *QueryBuilder) ([]News, error) {
	var news []News
	query := nr.Model(&news)
	query = qb.Apply(query)
	err := query.Where("\"statusId\" != ?", StatusDeleted).Select()
	if errors.Is(err, pg.ErrNoRows) {
		err = errors.New("no news were found with this filter")
	}
	return news, err
}

// NewsWithPagination возвращает новости с пагинацией и фильтрами
func (nr *NewsRepo) NewsWithPagination(page, pageSize int, qb *QueryBuilder) ([]News, error) {
	var news []News
	offset := (page - 1) * pageSize
	query := nr.Model(&news)
	query = qb.Apply(query)

	err := query.Where("\"statusId\" != ?", StatusDeleted).
		Offset(offset).
		Limit(pageSize).
		Select()
	if errors.Is(err, pg.ErrNoRows) {
		err = errors.New("no news were found on this page")
	}
	return news, err
}

// CategoryByID возвращает Category по id из бд
func (nr *NewsRepo) CategoryByID(id int) (*Category, error) {
	category := &Category{ID: id}
	err := nr.Model(category).Where("\"statusId\" != ?", StatusDeleted).WherePK().Select()
	if errors.Is(err, pg.ErrNoRows) {
		err = errors.New("no categories were found with this categoryID")
	}
	return category, err
}

// TagByID возвращает Tag по id из бд
func (nr *NewsRepo) TagByID(id int) (*Tag, error) {
	tag := &Tag{ID: id}
	err := nr.Model(tag).Where("\"statusId\" != ?", StatusDeleted).WherePK().Select()
	if errors.Is(err, pg.ErrNoRows) {
		err = errors.New("no tags were found with this tagID")
	}
	return tag, err
}

// CategoriesWithSort возвращает все категории со статусом не равным 3
// и отсортированными по полю orderNumber, а затем по полю title
func (nr *NewsRepo) CategoriesWithSort() ([]Category, error) {
	var categories []Category
	err := nr.Model(&categories).
		Where(" \"statusId\" != ?", StatusDeleted).
		OrderExpr("\"orderNumber\" IS NULL, \"orderNumber\" ASC, title ASC").
		Select()
	if errors.Is(err, pg.ErrNoRows) {
		err = errors.New("no categories were found")
	}
	return categories, err
}

// TagsWithSort возвращает все теги со статусом не равным 3, отсортированные по полю title
func (nr *NewsRepo) TagsWithSort() ([]Tag, error) {
	var tags []Tag
	err := nr.Model(&tags).
		Where(" \"statusId\" != ?", StatusDeleted).
		OrderExpr("title ASC").
		Select()
	if errors.Is(err, pg.ErrNoRows) {
		err = errors.New("no tags were found")
	}
	return tags, err
}
