package repositories_test

import (
	"net/http"
	"testing"

	"github.com/crixo/woa-go-api/model"
	"github.com/crixo/woa-go-api/patients/repositories"
	"github.com/crixo/woa-go-api/test"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	db := test.SetUp(t)

	patient := model.Pazient{
		PazientProfile: model.PazientProfile{
			FirstName: "Cristiano",
			LastName:  "Degiorgis"},
		// RemoteHistories: []model.RemoteHistory{
		// 	{Description: "prima anamnesi remota", HistoryKind: historyKinds[0]},
		// 	{Description: "seconda anamnesi remota", HistoryKindID: historyKinds[1].ID},
		// },
	}

	db.Create(&patient)

	r := repositories.NewGormPatientRepository(db)

	request, _ := http.NewRequest("GET", "http://example.com?limit=20&offset=0", nil)
	res, paginator, _ := r.Find(request)

	n := int64(1)
	assert.Len(t, res, 1)
	assert.True(t, paginator.Count == n)
	assert.False(t, paginator.NextURI.Valid)

	// mockArticles := []models.Article{
	// 	models.Article{
	// 		ID: 1, Title: "title 1", Content: "content 1",
	// 		Author: models.Author{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
	// 	},
	// 	models.Article{
	// 		ID: 2, Title: "title 2", Content: "content 2",
	// 		Author: models.Author{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
	// 	},
	// }

	// rows := sqlmock.NewRows([]string{"id", "title", "content", "author_id", "updated_at", "created_at"}).
	// 	AddRow(mockArticles[0].ID, mockArticles[0].Title, mockArticles[0].Content,
	// 		mockArticles[0].Author.ID, mockArticles[0].UpdatedAt, mockArticles[0].CreatedAt).
	// 	AddRow(mockArticles[1].ID, mockArticles[1].Title, mockArticles[1].Content,
	// 		mockArticles[1].Author.ID, mockArticles[1].UpdatedAt, mockArticles[1].CreatedAt)

	// query := "SELECT id,title,content, author_id, updated_at, created_at FROM article WHERE created_at > \\? ORDER BY created_at LIMIT \\?"

	// mock.ExpectQuery(query).WillReturnRows(rows)
	// a := articleRepo.NewMysqlArticleRepository(db)
	// cursor := articleRepo.EncodeCursor(mockArticles[1].CreatedAt)
	// num := int64(2)
	// list, nextCursor, err := a.Fetch(context.TODO(), cursor, num)
	// assert.NotEmpty(t, nextCursor)
	// assert.NoError(t, err)
	// assert.Len(t, list, 2)
}
