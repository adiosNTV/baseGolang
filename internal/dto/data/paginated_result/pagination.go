package data

type PaginatedResult struct {
	Items       interface{} `json:"items"`
	TotalCount  int         `json:"totalCount"`
	HasNextPage bool        `json:"hasNextPage"`
	HasPrevPage bool        `json:"hasPrevPage"`
	PageNumber  int         `json:"pageNumber"`
	PageSize    int         `json:"pageSize"`
}
