package newsportal

import (
	"log"
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

func (m Manager) NewsByID(id int) (*NewsSummary, error) {
	news, err := m.nr.NewsByID(id)
	return NewsSummaryFromDb(news), err
}

func (m Manager) TagsByIDs(ids []int) ([]Tag, error) {
	tags, err := m.nr.TagsByIDs(ids)
	return TagsFromDb(tags), err
}

func (m Manager) News(categoryID, tagID, page, pageSize *int, sortTitle *bool) ([]NewsSummary, error) {
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
	log.Println(page, pageSize)
	news, err := m.ByPage(*page, *pageSize, &qb)
	return news, err
}

// ByPage возвращает все новости с определённой страницы
func (m Manager) ByPage(page, pageSize int, qb *db.QueryBuilder) ([]NewsSummary, error) {
	// get news with pagination
	news, err := m.nr.NewsWithPagination(page, pageSize, qb)
	return NewsFromDb(news), err
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
