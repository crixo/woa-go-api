package sqlite

import (
	"os"
	"testing"

	"github.com/crixo/woa-go-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbFile = "woa-go-test.db"

func TestBootstrapComplexEntity(t *testing.T) {
	db := setup(t)

	historyKinds := []model.HistoryKind{
		model.HistoryKind{Description: "Kind 1"},
		model.HistoryKind{Description: "Kind 2"},
	}

	for i, element := range historyKinds {
		db.Create(&element)
		t.Logf("%+v\n", &element)
		t.Logf("DB sequence for element %d is %d", i, element.ID)
		// TODO: check how to have element as reference type avoiding next line
		historyKinds[i].ID = (&element).ID
	}

	for i, el := range historyKinds {
		t.Logf("%+v\n", &el)
		if historyKinds[i].ID == 0 {
			t.Errorf("DB sequence for element %d has not been genereated", i)
			t.FailNow()
		}
	}

	patient := model.Patient{
		PatientProfile: model.PatientProfile{
			FirstName: "Cristiano",
			LastName:  "Degiorgis"},
		RemoteHistories: []model.RemoteHistory{
			{Description: "prima anamnesi remota", HistoryKind: historyKinds[0]},
			{Description: "seconda anamnesi remota", HistoryKindID: historyKinds[1].ID},
		},
	}

	db.Create(&patient)

	var patientFromDb model.Patient
	db.Preload("RemoteHistories").
		Preload("RemoteHistories.HistoryKind").
		First(&patientFromDb)

	if len(patientFromDb.RemoteHistories) != 2 {
		t.Errorf("len of RemoteHistories should be 2")
		t.FailNow()
	}

	actual := patientFromDb.RemoteHistories[0].HistoryKindID
	expected := historyKinds[0].ID
	if actual != expected {
		t.Errorf("RemoteHistories[0].ID should be %d instead of %d", actual, expected)
		t.FailNow()
	}

	if len(patientFromDb.RemoteHistories[0].HistoryKind.Description) == 0 {
		t.Error("RemoteHistories[0].HistoryKind should be filled")
		t.FailNow()
	}

	teardown(t)
}

func setup(t *testing.T) (db *gorm.DB) {
	os.Remove(dbFile)

	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Patient{},
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

func teardown(t *testing.T) {
	os.Remove(dbFile)
}
