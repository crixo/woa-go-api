package model

import (
	"net/http"

	"github.com/ulule/paging"

	"github.com/google/uuid"
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

// PageFetchResult is a generic paged result
type PageFetchResult struct {
	Total   int64
	Results []interface{}
}

// ErrorCode contians all the possible error code for API calls
type ErrorCode int

const (
	INALID_INPUT ErrorCode = 1 + iota
	RESOURCE_IDENTIFIER_MISMATCH
)

var errorCodes = [...]string{
	"INALID_INPUT",
	"RESOURCE_IDENTIFIER_MISMATCH",
}

// ErrorResponse contains the error information in case of an api call fails
type ErrorResponse struct {
	Code    ErrorCode
	Message string
	TraceID uuid.UUID
}
