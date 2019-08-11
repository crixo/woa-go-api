package test

import (
	"os"
	"testing"

	"github.com/crixo/woa-go-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbFile = "woa-go-test.db"

func SetUp(t *testing.T) (db *gorm.DB) {
	os.Remove(dbFile)

	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Pazient{},
		&model.RemoteHistory{},
		&model.HistoryKind{},
		&model.ExaminationKind{},
		&model.Province{},
		&model.Consultation{},
		&model.NextHistory{},
		&model.Examination{},
		&model.Assessment{},
		&model.Treatment{},
		// all other tables of you app
	).Error
	if err != nil {
		t.Fatalf("Could not migrate: %v", err)
	}

	return db
}

func TearDown(t *testing.T) {
	os.Remove(dbFile)
}
