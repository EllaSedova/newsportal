package newsportal

import (
	"newsportal/pkg/db"
)

type Manager struct {
	nr db.NewsRepo
}

func ptri(r int) *int { return &r }
func NewManager(db db.NewsRepo) *Manager {
	return &Manager{nr: db}
}
func (m Manager) NewsByID(id int) (*NewsSummary, error) {
	news, err := m.nr.NewsByID(id)
	return NewsSummaryFromDb(news), err
}

func (m Manager) News(id, categoryID, tagID, page, pageSize *int, sortTitle *bool) ([]NewsSummary, error) {
	qb := m.nr.QB
	if id != nil {
		qb.AddFilter(db.Columns.News.ID, *id)
	}
	if categoryID != nil {
		qb.AddFilter("\"categoryId\"", *categoryID)
	}
	if tagID != nil {
		qb.AddNewFilter("ANY (\"tagIds\")", *tagID)
	}
	if sortTitle != nil && *sortTitle {
		qb.AddSort("title", true)
	}

	if pageSize != nil { // если есть лимит
		if page == nil { // если нет пагинации
			page = ptri(1)
		}
		news, err := m.ByPage(*page, *pageSize, &qb)
		return news, err
	} else {
		news, err := m.Default(&qb) // если нет ни лимита, ни пагинации
		return news, err
	}
}

func (m Manager) Default(qb *db.QueryBuilder) ([]NewsSummary, error) {
	news, err := m.nr.NewsWithFilters(qb)
	return NewsFromDb(news), err
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
