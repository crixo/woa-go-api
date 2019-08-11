package sqlite

import (
	"log"
	"os"

	//"github.com/crixo/woa-go-api/model"
	"github.com/crixo/woa-go-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/gormigrate.v1"
)

// InitialMigration handles gorms migrations for development
func InitialMigration() *gorm.DB {
	dbFile := "woa-go.db"

	var err = os.Remove(dbFile)

	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

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
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")

	historyKinds := []model.HistoryKind{
		model.HistoryKind{Description: "Kind 1"},
		model.HistoryKind{Description: "Kind 2"},
	}

	for i, element := range historyKinds {
		db.Create(&element)
		log.Printf("%+v\n", &element)
		log.Printf("%+v\n", element.ID)
		historyKinds[i].ID = (&element).ID
	}
	log.Printf("%+v\n", historyKinds)
	patient := model.Pazient{
		PazientProfile: model.PazientProfile{
			FirstName: "Cristiano",
			LastName:  "Degiorgis"},
		RemoteHistories: []model.RemoteHistory{
			{Description: "prima anamnesi remota", HistoryKind: historyKinds[0]},
			{Description: "seconda anamnesi remota", HistoryKindID: historyKinds[1].ID},
		},
	}

	db.Create(&patient)

	return db
}

// GoMigrate handles struturected migrations
func GoMigrate() {
	db, err := gorm.Open("sqlite3", "mydb.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// you migrations here
	})

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&model.User{},
			&model.Pazient{},
			&model.RemoteHistory{},
			&model.HistoryKind{},
			&model.ExaminationKind{},
			&model.Province{},
			// all other tables of you app
		).Error
		if err != nil {
			return err
		}

		// if err := tx.Model(Pet{}).AddForeignKey("person_id", "people (id)", "RESTRICT", "RESTRICT").Error; err != nil {
		// 	return err
		// }
		// all other foreign keys...
		return nil
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
