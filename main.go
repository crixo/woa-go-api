package main

import (
	//"encoding/json"
	"fmt"
	//"net/http"
	"github.com/crixo/woa-go-api/api"
	//"github.com/crixo/woa-go-api/model"
	//
	"github.com/crixo/woa-go-api/sqlite"
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
	fmt.Println("Go ORM Tutorial")

	db := sqlite.InitialMigration()
	// Handle Subsequent requests
	api.HandleRequests(db)
}
