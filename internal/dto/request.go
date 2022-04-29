package dto

import (
	"bytes"
	"fmt"
	"regexp"
)

type Pageable struct {
	Limit  int         `json:"limit,omitempty" form:"limit"`
	Page   int         `json:"page,omitempty" form:"page"`
	Rows   interface{} `json:"rows"`
	Search string      `json:"search" form:"search"`
	Total  int64       `json:"total"`
}

var regSQL = "('(''|[^'])*')|(;)|(\b(ALTER|CREATE|DELETE|DROP|EXEC(UTE){0,1}|INSERT( +INTO){0,1}|MERGE|SELECT|UPDATE|UNION( +ALL){0,1})\b)"

func (p *Pageable) SetDefaults() {
	p.Limit = 0
	p.Page = 1
	p.Search = ""

}

func (p *Pageable) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pageable) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pageable) GetPage() int {

	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pageable) GetSearchString() string {
	match, _ := regexp.MatchString(regSQL, p.Search)
	if match {
		return ""
	}
	return p.Search

}

type SearchOptions struct {
	Column       string
	Operator     string
	SearchString string
}

func (p *Pageable) GetSearchOptions() []SearchOptions {
	// This can be extended to be received from client and built a query engine. Currently hardcoding it with 2 columns needed to filter on
	// [["name","like","john"],["OR"],["description","like", "something"]]

	searchOptions := []SearchOptions{
		{
			Column:       "title",
			Operator:     "LIKE",
			SearchString: p.GetSearchString(),
		},
		{
			Column:       "description",
			Operator:     "LIKE",
			SearchString: p.GetSearchString(),
		},
	}
	return searchOptions

}

func (p *Pageable) BuildSearchString() (string, []interface{}) {
	var searchByteBuffer bytes.Buffer
	sOptions := p.GetSearchOptions()
	sValues := []interface{}{}

	for i, s := range sOptions {
		searchByteBuffer.WriteString(fmt.Sprintf(" %s %s ? ", s.Column, s.Operator))
		sValues = append(sValues, fmt.Sprintf("%%%v%%", s.SearchString))
		if i < len(sOptions)-1 {
			searchByteBuffer.WriteString("OR")
		}
	}

	return searchByteBuffer.String(), sValues
}
