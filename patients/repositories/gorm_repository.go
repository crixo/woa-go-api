package repositories

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crixo/woa-go-api/model"
	"github.com/crixo/woa-go-api/patients"
	"github.com/jinzhu/gorm"
	"github.com/ulule/paging"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type gormPatientRepository struct {
	DB *gorm.DB
}

// NewGormPatientRepository will create an object that represent the patients.Repository interface
func NewGormPatientRepository(DB *gorm.DB) patients.Repository {
	return &gormPatientRepository{DB}
}

func (r *gormPatientRepository) Find(requestPaginator *model.RequestPaginator) ([]*model.Pazient, *paging.OffsetPaginator, error) {
	// query := `SELECT id,title,content, author_id, updated_at, created_at
	// 						FROM article WHERE created_at > ? ORDER BY created_at LIMIT ? `

	// decodedCursor, err := DecodeCursor(cursor)
	// if err != nil && cursor != "" {
	// 	return nil, "", models.ErrBadParamInput
	// }

	// res, err := m.fetch(ctx, query, decodedCursor, num)
	// if err != nil {
	// 	return nil, "", err
	// }

	// nextCursor := ""
	// if len(res) == int(num) {
	// 	nextCursor = EncodeCursor(res[len(res)-1].CreatedAt)
	// }

	// return res, nextCursor, err

	res := []*model.Pazient{}

	// Step 1: create the store. It takes your database connection pointer, a
	// pointer to models and the GORM "ORDER BY" string.
	store, err := paging.NewGORMStore(r.DB, &res)
	if err != nil {
		log.Fatal(err)
	}

	// Step 2: create options. Here, we use the default ones (see below).
	options := model.PaginatorOptions //paging.NewOptions()

	// Step 3: create a paginator instance and pass your store, your current HTTP
	// request and your options as arguments.
	request, _ := http.NewRequest("GET",
		fmt.Sprintf("http://example.com?%s=%d0&%s=%d",
			options.LimitKeyName, requestPaginator.Limit,
			options.OffsetKeyName, requestPaginator.Offset),
		nil)
	paginator, err := paging.NewOffsetPaginator(store, request, options)
	if err != nil {
		log.Fatal(err)
	}

	//Step 4: call the paginator.Page() method to get the page instance.
	err = paginator.Page()
	if err != nil {
		log.Fatal(err)
	}

	return res, paginator, err

}
