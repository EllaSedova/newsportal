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

type NewsResponse struct {
	News  []News
	Count *int
}

type TagMap map[int]Tag

func (t *TagMap) Fill(tagIDMap map[int]struct{}, m *Manager) error {

	var uniqueTagIDs []int
	for tagID := range tagIDMap {
		uniqueTagIDs = append(uniqueTagIDs, tagID)
	}

	// возвращаем теги из бд
	tags, err := m.TagsByIDs(uniqueTagIDs)

	// создаём карту тегов
	for _, tag := range tags {
		(*t)[tag.ID] = tag
	}

	return err
}

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
	// заполняем карту уникальных tagId
	tagMap := make(TagMap)
	err = tagMap.Fill(tagIDMap, &m)

	var newsTags []Tag
	for _, tagID := range news.TagIDs {
		newsTags = append(newsTags, tagMap[tagID])
	}
	return NewsFromDb(news, newsTags), err
}

func (m Manager) TagsByIDs(ids []int) ([]Tag, error) {
	if ids != nil {
		tags, err := m.nr.TagsByIDs(ids)
		return TagsFromDb(tags), err
	}
	return nil, nil
}

func (m Manager) News(categoryID, tagID, page, pageSize *int, sortTitle, countRequest *bool) (NewsResponse, error) {
	qb := m.nr.QB

	if categoryID != nil {
		qb.AddFilterEqual(db.Columns.News.CategoryID, *categoryID)
	}
	if tagID != nil {
		qb.AddFilterAny(db.Columns.News.TagIDs, *tagID)
	}
	if sortTitle != nil && *sortTitle {
		qb.AddSort(db.Columns.News.Title, true)
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
	if sortTitle != nil && *sortTitle {
		qb.AddSort(db.Columns.News.Title, true)
	}
	if countRequest != nil && *countRequest {
		count, err := m.Count(*page, *pageSize, &qb)
		return NewsResponse{News: nil, Count: &count}, err
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
		// заполняем карту уникальных tagId
		tagMap := make(TagMap)
		err = tagMap.Fill(tagIDMap, &m)
		var newNewsList []News
		for i, summary := range news {
			var newsTags []Tag
			for _, tagId := range summary.TagIDs {
				newsTags = append(newsTags, tagMap[tagId])
			}
			newNews := NewsFromDb(&news[i], newsTags)
			newNewsList = append(newNewsList, *newNews)
		}
		return NewsResponse{News: newNewsList, Count: nil}, err
	}
	return NewsResponse{News: nil, Count: nil}, err
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
