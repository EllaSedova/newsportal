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
	SortTitle  *bool
	Count      *bool
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
func (rs NewsService) NewsWithFilters(categoryID, tagID, page, pageSize *int, sortTitle, count *bool) (NewsResponse, error) {
	var params FilterParams
	params.Fill(categoryID, tagID, page, pageSize, sortTitle, count)
	newsResponse, err := rs.m.News(params.CategoryID, params.TagID, params.Page, params.PageSize, params.SortTitle, params.Count)
	var newNewsList []NewsSummary
	for _, summary := range newsResponse.News {
		newNews := NewsSummaryFromManager(&summary)
		newNewsList = append(newNewsList, *newNews)
	}
	return NewsResponse{News: newNewsList, Count: newsResponse.Count}, err
}

func (p *FilterParams) Fill(categoryID, tagID, page, pageSize *int, sortTitle, count *bool) {
	p.CategoryID = categoryID
	p.TagID = tagID
	p.Page = page
	p.PageSize = pageSize
	p.SortTitle = sortTitle
	p.Count = count
}
