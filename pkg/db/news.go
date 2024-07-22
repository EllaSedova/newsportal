package db

import (
	"errors"
	"github.com/go-pg/pg/v10"
)

type NewsRepo struct {
	*pg.DB
	qb QueryBuilder
}

const (
	StatusEnabled = iota + 1
	StatusDisabled
	StatusDeleted
)

func NewNewsRepo(db *pg.DB) NewsRepo {
	return NewsRepo{DB: db, qb: *NewQueryBuilder()}
}

func AddFilter(categoryID, tagID *int, qb *QueryBuilder) {
	if categoryID != nil {
		qb.AddFilterEqual(Columns.News.CategoryID, *categoryID)
	}
	if tagID != nil {
		qb.AddFilterAny(Columns.News.TagIDs, *tagID)
	}
}

// NewsByID возвращает News по id из бд
func (nr *NewsRepo) NewsByID(id int) (*News, error) {
	news := &News{ID: id}
	err := nr.Model(news).Relation(`Category`).Where(`t."statusId" != ?`, StatusDeleted).WherePK().Select()
	if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}
	return news, err
}

// NewsWithPagination возвращает новости с пагинацией и фильтрами
func (nr *NewsRepo) NewsWithPagination(page, pageSize int, categoryID, tagID *int) ([]News, error) {
	var news []News
	qb := nr.qb
	AddFilter(categoryID, tagID, &qb)
	offset := (page - 1) * pageSize
	query := nr.Model(&news)
	query = qb.Apply(query)
	err := query.Relation(`Category`).Where(`t."statusId" != ?`, StatusDeleted).
		Offset(offset).
		Limit(pageSize).
		Select()
	if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}
	return news, err
}

// NewsCount возвращает количество новостей с пагинацией и фильтрами
func (nr *NewsRepo) NewsCount(page, pageSize int, categoryID, tagID *int) (int, error) {
	qb := nr.qb
	AddFilter(categoryID, tagID, &qb)
	offset := (page - 1) * pageSize
	query := nr.Model((*News)(nil))
	query = qb.Apply(query)
	count, err := query.Relation(`Category`).Where(`t."statusId" != ?`, StatusDeleted).
		Offset(offset).
		Limit(pageSize).
		Count()
	if errors.Is(err, pg.ErrNoRows) {
		return 0, nil
	}
	return count, err
}

// CategoryByID возвращает Category по id из бд
func (nr *NewsRepo) CategoryByID(id int) (*Category, error) {
	category := &Category{ID: id}
	err := nr.Model(category).Where(`"statusId" != ?`, StatusDeleted).WherePK().Select()
	if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}
	return category, err
}

// TagByID возвращает Tag по id из бд
func (nr *NewsRepo) TagByID(id int) (*Tag, error) {
	tag := &Tag{ID: id}
	err := nr.Model(tag).Where(`"statusId" != ?`, StatusDeleted).WherePK().Select()
	if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}
	return tag, err
}

// TagsByIDs возвращает Tags по ids из бд
func (nr *NewsRepo) TagsByIDs(ids []int) ([]Tag, error) {
	var tags []Tag
	err := nr.Model(&tags).Where(`"tagId" IN (?)`, pg.In(ids)).Select()
	if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}
	return tags, err
}

// CategoriesWithSort возвращает все категории со статусом не равным 3
// и отсортированными по полю orderNumber, а затем по полю title
func (nr *NewsRepo) CategoriesWithSort() ([]Category, error) {
	var categories []Category
	err := nr.Model(&categories).
		Where(` "statusId" != ?`, StatusDeleted).
		OrderExpr(`"orderNumber" IS NULL, "orderNumber" ASC, title ASC`).
		Select()
	if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}
	return categories, err
}

// TagsWithSort возвращает все теги со статусом не равным 3, отсортированные по полю title
func (nr *NewsRepo) TagsWithSort() ([]Tag, error) {
	var tags []Tag
	err := nr.Model(&tags).
		Where(` "statusId" != ?`, StatusDeleted).
		OrderExpr("title ASC").
		Select()
	if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}
	return tags, err
}
