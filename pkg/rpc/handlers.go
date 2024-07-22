package rpc

import (
	"context"
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
	ctx := context.Background()
	news, err := rs.m.NewsByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return NewsFromManager(news), nil
}

// Categories получение всех категорий
func (rs NewsService) Categories() ([]Category, error) {
	ctx := context.Background()
	categories, err := rs.m.Categories(ctx)
	if err != nil {
		return nil, err
	}

	return CategoriesFromManager(categories), err
}

// Tags получение всех тегов
func (rs NewsService) Tags() ([]Tag, error) {
	ctx := context.Background()
	tags, err := rs.m.Tags(ctx)
	if err != nil {
		return nil, err
	}

	return TagsFromManager(tags), err
}

// NewsWithFilters получение новости с фильтрами
func (rs NewsService) NewsWithFilters(categoryID, tagID, page, pageSize *int) ([]NewsSummary, error) {
	ctx := context.Background()
	var params FilterParams
	params.Fill(categoryID, tagID, page, pageSize)

	newsResponse, err := rs.m.News(ctx, params.CategoryID, params.TagID, params.Page, params.PageSize)
	if err != nil {
		return nil, err
	}

	var newNewsList []NewsSummary
	for _, summary := range newsResponse {
		newNews := NewsSummaryFromManager(&summary)
		newNewsList = append(newNewsList, *newNews)
	}
	return newNewsList, nil
}

// NewsCountWithFilters получение количества новостей с фильтрами
func (rs NewsService) NewsCountWithFilters(categoryID, tagID *int) (*int, error) {
	ctx := context.Background()
	var params FilterParams
	params.Fill(categoryID, tagID, nil, nil)

	count, err := rs.m.NewsCount(ctx, params.CategoryID, params.TagID)

	return count, err
}

func (p *FilterParams) Fill(categoryID, tagID, page, pageSize *int) {
	p.CategoryID = categoryID
	p.TagID = tagID
	p.Page = page
	p.PageSize = pageSize
}
