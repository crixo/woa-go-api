package patients

import (
	"github.com/crixo/woa-go-api/model"

	"github.com/ulule/paging"
)

// Repository represent the article's repository contract
type Repository interface {
	// Fetch(ctx context.Context, cursor string, num int64) (res []*model.PatientProfile, nextCursor string, err error)
	// GetByID(ctx context.Context, id int64) (*model.PatientProfile, error)
	// Update(ctx context.Context, ar *model.PatientProfile) error
	Create(p *model.Patient) error
	// Delete(ctx context.Context, id int64) error
	Find(requestPaginator *model.RequestPaginator) (res []*model.Patient, paginator *paging.OffsetPaginator, err error)
}
