package rpc

import (
	"github.com/vmkteam/zenrpc/v2"
	"newsportal/pkg/newsportal"
)

type NewsService struct {
	zenrpc.Service
	m *newsportal.Manager
}

type FilterParams struct {
	CategoryID *int
	TagID      *int
	Page       *int
	PageSize   *int
}

func NewNewsService(m *newsportal.Manager) *NewsService {
	return &NewsService{
		m: m,
	}
}

// NewsByID получение новости по id
func (rs NewsService) NewsByID(id int) (*News, error) {
	news, err := rs.m.NewsByID(id)
	return NewsFromManager(news), err
}

// Categories получение всех категорий
func (rs NewsService) Categories() ([]Category, error) {
	categories, err := rs.m.Categories()
	return CategoriesFromManager(categories), err
}

// Tags получение всех тегов
func (rs NewsService) Tags() ([]Tag, error) {
	tags, err := rs.m.Tags()
	return TagsFromManager(tags), err
}

// NewsWithFilters получение новости с фильтрами
func (rs NewsService) NewsWithFilters(categoryID, tagID, page, pageSize *int) ([]NewsSummary, error) {
	var params FilterParams
	params.Fill(categoryID, tagID, page, pageSize)

	newsResponse, err := rs.m.News(params.CategoryID, params.TagID, params.Page, params.PageSize)

	var newNewsList []NewsSummary
	for _, summary := range newsResponse {
		newNews := NewsSummaryFromManager(&summary)
		newNewsList = append(newNewsList, *newNews)
	}
	return newNewsList, err
}

// NewsCountWithFilters получение количества новостей с фильтрами
func (rs NewsService) NewsCountWithFilters(categoryID, tagID, page, pageSize *int) (*int, error) {
	var params FilterParams
	params.Fill(categoryID, tagID, page, pageSize)

	count, err := rs.m.NewsCount(params.CategoryID, params.TagID, params.Page, params.PageSize)

	return count, err
}
func (p *FilterParams) Fill(categoryID, tagID, page, pageSize *int) {
	p.CategoryID = categoryID
	p.TagID = tagID
	p.Page = page
	p.PageSize = pageSize
}
