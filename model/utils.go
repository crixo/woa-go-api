package model

import (
	"net/http"

	"github.com/ulule/paging"
)

// PaginatorOptions contains paginator options for the current application
var PaginatorOptions = paging.NewOptions()

// RequestPaginator contains the pagination input paramters
type RequestPaginator struct {
	Offset int64
	Limit  int64
	Filter string
}

// CreateRequestPaginatorFromRequest returns an intance filled using http request
func CreateRequestPaginatorFromRequest(r *http.Request) *RequestPaginator {
	return &RequestPaginator{
		Limit:  paging.GetLimitFromRequest(r, PaginatorOptions),
		Offset: paging.GetOffsetFromRequest(r, PaginatorOptions),
		Filter: r.URL.Query().Get("filter"),
	}
}
