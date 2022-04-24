package dto

type Pageable struct {
	Limit     int         `json:"limit,omitempty" form:"limit"`
	Page      int         `json:"page,omitempty" form:"page"`
	TotalRows int64       `json:"total_rows"`
	Rows      interface{} `json:"rows"`
	Search    string      `json:"search" form:"search"`
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
			SearchString: p.Search,
		},
		{
			Column:       "description",
			Operator:     "LIKE",
			SearchString: p.Search,
		},
	}
	return searchOptions

}
