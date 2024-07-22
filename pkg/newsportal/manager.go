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

type TagMap map[int]Tag

func NewManager(db db.NewsRepo) *Manager {
	return &Manager{nr: db}
}

func CheckPagination(page, pageSize **int) {
	if *page == nil {
		*page = ptri(defaultPage)
	} else if **page <= 0 {
		*page = ptri(defaultPage)
	}

	if *pageSize == nil {
		*pageSize = ptri(defaultPageSize)
	} else if **pageSize <= 0 {
		*pageSize = ptri(defaultPageSize)
	}
}

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

func (m Manager) FillTags(newNewsList *[]News, news []db.News) error {
	// собираем все уникальные tagID
	tagIDMap := make(map[int]struct{})
	for _, summary := range news {
		for _, tagId := range summary.TagIDs {
			tagIDMap[tagId] = struct{}{}
		}
	}

	// заполняем карту уникальных tagId
	tagMap := make(TagMap)
	err := tagMap.Fill(tagIDMap, &m)

	for i, summary := range news {
		var newsTags []Tag
		for _, tagId := range summary.TagIDs {
			newsTags = append(newsTags, tagMap[tagId])
		}
		*newNewsList = append(*newNewsList, *NewsFromDB(&news[i], newsTags))
	}

	return err
}

func (m Manager) NewsByID(id int) (*News, error) {
	news, err := m.nr.NewsByID(id)
	if err != nil {
		return nil, err
	}

	if news == nil {
		return nil, nil
	}

	var newNews []News
	err = m.FillTags(&newNews, []db.News{*news})

	return &newNews[0], err
}

func (m Manager) News(categoryID, tagID, page, pageSize *int) ([]News, error) {
	CheckPagination(&page, &pageSize)
	news, err := m.nr.NewsWithPagination(*page, *pageSize, categoryID, tagID)
	if err != nil {
		return nil, err
	}

	if len(news) == 0 {
		return nil, nil
	}

	var newNews []News
	err = m.FillTags(&newNews, news)
	return newNews, err
}

func (m Manager) NewsCount(categoryID, tagID, page, pageSize *int) (*int, error) {
	CheckPagination(&page, &pageSize)
	count, err := m.nr.NewsCount(*page, *pageSize, categoryID, tagID)

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
