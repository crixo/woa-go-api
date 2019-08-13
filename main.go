package main

import (
	"log"
	"os"
	"time"

	"github.com/crixo/woa-go-api/api"
	"github.com/crixo/woa-go-api/sqlite"
	"github.com/jinzhu/gorm"
)

// @title Swagger Example API
// @version 1.0
// @description This is a crixo cample
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email cris@swagger.io

// @host localhost:8081
// @BasePath
func main() {
	log.Println("Go ORM Tutorial")

	// gorm.NowFunc = func() time.Time {
	// 	return time.Now().UTC()
	// }

	dbFile := "woa-go.db"
	var err = os.Remove(dbFile)
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetNowFuncOverride(myNowFunc)
	db.LogMode(true)

	sqlite.InitialMigration(db)
	// Handle Subsequent requests
	api.HandleRequests(db)
}

var myNowFunc = func() time.Time {
	return time.Now().UTC()
}
