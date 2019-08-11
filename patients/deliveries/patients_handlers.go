package deliveries

import (
	//"context"
	"encoding/json"
	"log"
	"net/http"

	//"log"
	"github.com/crixo/woa-go-api/model"
	"github.com/crixo/woa-go-api/patients"
	"github.com/gorilla/mux"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// PatientsHandler  represent the httphandler for article
type PatientsHandler struct {
	PatientsRepository patients.Repository
}

// NewPatientsHandler contains patients endpoint definitions
func NewPatientsHandler(router *mux.Router, patientsRepository patients.Repository) {
	handler := &PatientsHandler{
		PatientsRepository: patientsRepository,
	}
	router.HandleFunc("/patients", handler.findPatients).Methods("GET")
	// e.GET("/articles", handler.FetchArticle)
	// e.POST("/articles", handler.Store)
	// e.GET("/articles/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
}

// userProfile godoc
// @Summary Patient Profile
// @Description Patient Profile description
// @Tags patients
// @ID patients-profile
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Pazient
// @Router /patients [get]
// @Security ApiKeyAuth
func (ph *PatientsHandler) findPatients(w http.ResponseWriter, r *http.Request) {

	res, _, err := ph.PatientsRepository.Find(model.CreateRequestPaginatorFromRequest(r))
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	json.NewEncoder(w).Encode(res)

	log.Println("FindPatients")
}
