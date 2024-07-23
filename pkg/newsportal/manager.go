package newsportal

import (
	"context"
	"newsportal/pkg/db"
)

type Manager struct {
	nr db.NewsRepo
}

const defaultPage = 1
const defaultPageSize = 2

func ptri(r int) *int { return &r }

func NewManager(db db.NewsRepo) *Manager {
	return &Manager{nr: db}
}

func checkPagination(page, pageSize *int) (int, int) {
	if page == nil {
		page = ptri(defaultPage)
	} else if *page <= 0 {
		page = ptri(defaultPage)
	}

	if pageSize == nil {
		pageSize = ptri(defaultPageSize)
	} else if *pageSize <= 0 {
		pageSize = ptri(defaultPageSize)
	}
	return *page, *pageSize
}

func (m Manager) FillTags(ctx context.Context, newsList []News) error {
	// собираем все уникальные tagID
	tagIDMap := make(map[int]struct{})
	for _, summary := range newsList {
		for _, tagId := range summary.TagIDs {
			tagIDMap[tagId] = struct{}{}
		}
	}

	// заполняем карту уникальных tagId
	tagMap := make(map[int]Tag)
	var uniqueTagIDs []int
	for tagID := range tagIDMap {
		uniqueTagIDs = append(uniqueTagIDs, tagID)
	}

	// возвращаем теги из бд
	tags, err := m.TagsByIDs(ctx, uniqueTagIDs)

	// создаём карту тегов
	for _, tag := range tags {
		tagMap[tag.ID] = tag
	}
	for i, summary := range newsList {
		for _, tagId := range summary.TagIDs {
			_, exist := tagMap[tagId]
			if exist {
				newsList[i].Tags = append(newsList[i].Tags, tagMap[tagId])
			}
		}
	}

	return err
}

func (m Manager) NewsByID(ctx context.Context, id int) (*News, error) {
	news, err := m.nr.NewsByID(ctx, id, db.WithRelations(db.Columns.News.Category))
	if err != nil {
		return nil, err
	} else if news == nil {
		return nil, nil
	}

	n := NewNewsListFromDB([]db.News{*news})

	err = m.FillTags(ctx, n)

	return &n[0], err
}

func (m Manager) News(ctx context.Context, categoryID, tagID, page, pageSize *int) ([]News, error) {
	newPage, newPageSize := checkPagination(page, pageSize)
	searchParams := &db.NewsSearch{CategoryID: categoryID}
	if tagID != nil {
		searchParams = searchParams.WithTagID(tagID)
	}

	news, err := m.nr.NewsByFilters(ctx, searchParams, db.Pager{Page: newPage, PageSize: newPageSize}, db.WithRelations(db.Columns.News.Category))
	if err != nil {
		return nil, err
	} else if len(news) == 0 {
		return nil, nil
	}

	newsList := NewNewsListFromDB(news)
	err = m.FillTags(ctx, newsList)

	return newsList, err
}

func (m Manager) NewsCount(ctx context.Context, categoryID, tagID *int) (*int, error) {
	searchParams := &db.NewsSearch{CategoryID: categoryID}
	if tagID != nil {
		searchParams = searchParams.WithTagID(tagID)
	}
	count, err := m.nr.CountNews(ctx, searchParams)

	return &count, err
}

// Categories возвращает все категории
func (m Manager) Categories(ctx context.Context) ([]Category, error) {
	categories, err := m.nr.CategoriesByFilters(ctx, &db.CategorySearch{}, db.PagerNoLimit)

	return CategoriesFromDb(categories), err
}

func (m Manager) TagsByIDs(ctx context.Context, tagIDs []int) ([]Tag, error) {
	if len(tagIDs) == 0 {
		return nil, nil
	}

	tags, err := m.nr.TagsByFilters(ctx, &db.TagSearch{IDs: tagIDs}, db.PagerNoLimit)

	return TagsFromDb(tags), err
}

// Tags возвращает все теги
func (m Manager) Tags(ctx context.Context) ([]Tag, error) {
	tags, err := m.nr.TagsByFilters(ctx, &db.TagSearch{}, db.PagerNoLimit)

	return TagsFromDb(tags), err
}
