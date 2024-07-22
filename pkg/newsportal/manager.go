package newsportal

import (
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

func (m Manager) FillTags(newsList []News) error {
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
	tags, err := m.TagsByIDs(uniqueTagIDs)

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

func (m Manager) NewsByID(id int) (*News, error) {
	news, err := m.nr.NewsByID(id)
	if err != nil {
		return nil, err
	} else if news == nil {
		return nil, nil
	}

	n := NewNewsListFromDB([]db.News{*news})

	err = m.FillTags(n)

	return &n[0], err
}

func (m Manager) News(categoryID, tagID, page, pageSize *int) ([]News, error) {
	newPage, newPageSize := checkPagination(page, pageSize)
	news, err := m.nr.NewsWithPagination(newPage, newPageSize, categoryID, tagID)
	if err != nil {
		return nil, err
	} else if len(news) == 0 {
		return nil, nil
	}

	newsList := NewNewsListFromDB(news)
	err = m.FillTags(newsList)

	return newsList, err
}

func (m Manager) NewsCount(categoryID, tagID, page, pageSize *int) (*int, error) {
	newPage, newPageSize := checkPagination(page, pageSize)
	count, err := m.nr.NewsCount(newPage, newPageSize, categoryID, tagID)

	return &count, err
}

// Categories возвращает все категории
func (m Manager) Categories() ([]Category, error) {
	categories, err := m.nr.CategoriesWithSort()

	return CategoriesFromDb(categories), err
}

func (m Manager) TagsByIDs(ids []int) ([]Tag, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	tags, err := m.nr.TagsByIDs(ids)

	return TagsFromDb(tags), err
}

// Tags возвращает все теги
func (m Manager) Tags() ([]Tag, error) {
	tags, err := m.nr.TagsWithSort()

	return TagsFromDb(tags), err
}
