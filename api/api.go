package api

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/crixo/woa-go-api/docs" // docs is generated by Swag CLI, you have to import it.
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// HandleRequests registers all routes
func HandleRequests() {
	myRouter := mux.NewRouter()

	myRouter.StrictSlash(true)

	// @description users
	myRouter.HandleFunc("/patients", patients).Methods("GET")
	// myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	// myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	// myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")

	myRouter.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
	//httpSwagger.URL("http://localhost:8081/swagger/doc.json"), //The url pointing to API definition"
	//httpSwagger.URL("swagger.json"),
	))
	//http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	//myRouter.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// userProfile godoc
// @Summary Patient Profile
// @Description Patient Profile description
// @Tags patients
// @ID patients-profile
// @Accept  json
// @Produce  json
// @Router /patients [get]
// @Security ApiKeyAuth
func patients(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// var users []User
	// db.Find(&users)

	// json.NewEncoder(w).Encode(users)

	fmt.Println("patients")
}
