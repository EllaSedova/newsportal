package newsportal

import (
	"newsportal/pkg/db"
)

type Manager struct {
	nr db.NewsRepo
}

const defaultPage = 1
const defaultPageSize = 10

func ptri(r int) *int { return &r }

func NewManager(db db.NewsRepo) *Manager {
	return &Manager{nr: db}
}

func (m Manager) NewsByID(id int) (*News, error) {
	news, err := m.nr.NewsByID(id)
	if err != nil || news == nil {
		return nil, err
	}
	// собираем все уникальные tagID
	tagIDMap := make(map[int]struct{})

	for _, tagID := range news.TagIDs {
		tagIDMap[tagID] = struct{}{}
	}

	var uniqueTagIDs []int
	for tagID := range tagIDMap {
		uniqueTagIDs = append(uniqueTagIDs, tagID)
	}

	// возвращаем теги из бд
	tags, err := m.TagsByIDs(uniqueTagIDs)

	// создаём карту тегов
	tagMap := make(map[int]Tag)
	for _, tag := range tags {
		tagMap[tag.ID] = tag
	}

	var newsTags []Tag

	for _, tagID := range news.TagIDs {
		newsTags = append(newsTags, tagMap[tagID])
	}

	return NewsFromDb(news, newsTags), err
}

func (m Manager) TagsByIDs(ids []int) ([]Tag, error) {
	tags, err := m.nr.TagsByIDs(ids)
	return TagsFromDb(tags), err
}

func (m Manager) News(categoryID, tagID, page, pageSize *int, sortTitle *bool, countRequest bool) ([]News, int, error) {
	qb := m.nr.QB

	if categoryID != nil {
		qb.AddFilter(`t."categoryId"`, *categoryID)
	}
	if tagID != nil {
		qb.AddNewFilter(`ANY (t."tagIds")`, *tagID)
	}
	if sortTitle != nil && *sortTitle {
		qb.AddSort(`t.title`, true)
	}
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
	// если нужо только количество новостей
	if countRequest {
		count, err := m.Count(*page, *pageSize, &qb)
		return nil, count, err
	}
	news, err := m.ByPage(*page, *pageSize, &qb)
	if news != nil {
		// собираем все уникальные tagID
		tagIDMap := make(map[int]struct{})
		for _, summary := range news {
			for _, tagId := range summary.TagIDs {
				tagIDMap[tagId] = struct{}{}
			}
		}

		var uniqueTagIDs []int
		for tagId := range tagIDMap {
			uniqueTagIDs = append(uniqueTagIDs, tagId)
		}
		// возвращаем теги из бд
		tags, err := m.TagsByIDs(uniqueTagIDs)
		// создаём карту тегов
		tagMap := make(map[int]Tag)
		for _, tag := range tags {
			tagMap[tag.ID] = tag
		}

		var newNewsList []News
		for _, summary := range news {
			var newsTags []Tag
			for _, tagId := range summary.TagIDs {
				newsTags = append(newsTags, tagMap[tagId])
			}
			newNews := NewsFromDb(&summary, newsTags)
			newNewsList = append(newNewsList, *newNews)
		}
		return newNewsList, 0, err
	}

	return nil, 0, err
}

// ByPage возвращает все новости с определённой страницы
func (m Manager) ByPage(page, pageSize int, qb *db.QueryBuilder) ([]db.News, error) {
	// get news with pagination
	news, err := m.nr.NewsWithPagination(page, pageSize, qb)
	return news, err
}

// Count возвращает количество новостей с определённой страницы
func (m Manager) Count(page, pageSize int, qb *db.QueryBuilder) (int, error) {
	// get news with pagination
	count, err := m.nr.NewsCount(page, pageSize, qb)
	return count, err
}

// Categories возвращает все категории
func (m Manager) Categories() ([]Category, error) {
	categories, err := m.nr.CategoriesWithSort()
	return CategoriesFromDb(categories), err
}

// Tags возвращает все теги
func (m Manager) Tags() ([]Tag, error) {
	tags, err := m.nr.TagsWithSort()
	return TagsFromDb(tags), err
}
