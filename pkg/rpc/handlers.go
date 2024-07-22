package rpc

import (
	"github.com/vmkteam/zenrpc/v2"
	"newsportal/pkg/newsportal"
)

type RPCService struct {
	zenrpc.Service
	m *newsportal.Manager
}
type FilterParams struct {
	CategoryID *int
	TagID      *int
	Page       *int
	PageSize   *int
	SortTitle  *bool
}

func ptrb(b bool) *bool { return &b }
func NewRPCService(m *newsportal.Manager) *RPCService {
	return &RPCService{
		m: m,
	}
}

// NewsByID получение новости по id
func (rs RPCService) NewsByID(id int) (*News, error) {
	news, err := rs.m.NewsByID(id)
	return NewsFromManager(news), err
}

// Categories получение всех категорий
func (rs RPCService) Categories() ([]Category, error) {
	categories, err := rs.m.Categories()
	return CategoriesFromManager(categories), err
}

// Tags получение всех тегов
func (rs RPCService) Tags() ([]Tag, error) {
	tags, err := rs.m.Tags()
	return TagsFromManager(tags), err
}

// NewsWithFilters получение новости с фильтрами
func (rs RPCService) NewsWithFilters(categoryID, tagID, page, pageSize *int, sortTitle *bool) ([]NewsSummary, error) {
	var params FilterParams
	params.Fill(categoryID, tagID, page, pageSize, sortTitle)
	news, _, err := rs.m.News(params.CategoryID, params.TagID, params.Page, params.PageSize, params.SortTitle, false)
	var newNewsList []NewsSummary
	for _, summary := range news {
		newNews := NewsSummaryFromManager(&summary)
		newNewsList = append(newNewsList, *newNews)
	}
	return newNewsList, err
}

// NewsCountWithFilters получение количества новостей с фильтрами
func (rs RPCService) NewsCountWithFilters(categoryID, tagID, page, pageSize *int) (int, error) {
	var params FilterParams
	params.Fill(categoryID, tagID, page, pageSize, ptrb(false))
	_, count, err := rs.m.News(params.CategoryID, params.TagID, params.Page, params.PageSize, params.SortTitle, true)
	return count, err
}

func (p *FilterParams) Fill(categoryID, tagID, page, pageSize *int, sortTitle *bool) {
	p.CategoryID = categoryID
	p.TagID = tagID
	p.Page = page
	p.PageSize = pageSize
	p.SortTitle = sortTitle
}
