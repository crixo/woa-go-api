package patients

import (
	"net/http"

	"github.com/crixo/woa-go-api/model"

	"github.com/ulule/paging"
)

// Repository represent the article's repository contract
type Repository interface {
	// Fetch(ctx context.Context, cursor string, num int64) (res []*model.PazientProfile, nextCursor string, err error)
	// GetByID(ctx context.Context, id int64) (*model.PazientProfile, error)
	// Update(ctx context.Context, ar *model.PazientProfile) error
	// Create(ctx context.Context, a *model.PazientProfile) error
	// Delete(ctx context.Context, id int64) error
	Find(Request *http.Request) (res []*model.Pazient, paginator *paging.OffsetPaginator, err error)
}
