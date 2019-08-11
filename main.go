package main

import (
	//"encoding/json"
	"fmt"
	//"net/http"
	"github.com/crixo/woa-go-api/api"
	//"github.com/crixo/woa-go-api/model"
	//
	"github.com/crixo/woa-go-api/sqlite"
	"github.com/crixo/woa-go-api/test/subtest"
)

// @title Swagger Example API
// @version 1.0
// @description This is a crixo cample
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email cris@swagger.io

// @host cri.swagger.io
// @BasePath /v2
func main() {
	fmt.Println("Go ORM Tutorial")

	sqlite.InitialMigration()
	// Handle Subsequent requests
	api.HandleRequests()
	subtest.MyTest()
}
