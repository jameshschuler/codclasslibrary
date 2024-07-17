package common

type PaginationResponse struct {
	HasNextPage bool `json:"hasNextPage"`
	TotalCount  int  `json:"totalCount"`
}
