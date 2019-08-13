package deliveries

import (
	//"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	//"log"
	"github.com/crixo/woa-go-api/model"
	"github.com/crixo/woa-go-api/patients"
	"github.com/google/uuid"
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
	router.HandleFunc("/patients", handler.createPatient).Methods("POST")
	router.HandleFunc("/patients/{id}", handler.updatePatient).Methods("PUT")
	// e.GET("/articles", handler.FetchArticle)
	// e.POST("/articles", handler.Store)
	// e.GET("/articles/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
}

// @Summary Get paged patient profiles
// @Description Return a paged set of patient profiles w/ the total number of available patients
// @Tags patients
// @ID patients-profile-find
// @Accept  json
// @Produce  json
// @Success 200 {array} model.PatientProfile
// @Header 200 {int} X-Total-Count "Total number of records w/o pagination"
// @Router /patients [get]
func (ph *PatientsHandler) findPatients(w http.ResponseWriter, r *http.Request) {

	res1, pager, err := ph.PatientsRepository.Find(model.CreateRequestPaginatorFromRequest(r))
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	var res = make([]model.PatientProfile, len(res1))
	for index, element := range res1 {
		res[index] = element.PatientProfile
	}

	// pagedResult := model.PageFetchResult{}
	// for _, element := range res {
	// 	results := pagedResult.Results
	// 	pagedResult.Results = append(results, element)
	// }
	// pagedResult.Total = pager.Count

	// json.NewEncoder(w).Encode(pagedResult)

	w.Header().Set("X-Total-Count", strconv.Itoa(int(pager.Count)))
	json.NewEncoder(w).Encode(res)

	log.Println("FindPatients")
}

// @Summary Create Patient Profile
// @Description Create Patient Profile
// @Tags patients
// @ID patients-profile-create
// @Accept  json
// @Param patientProfile body model.PatientProfile true "Create new patient profile"
// @Produce  json
// @Success 200 {object} model.PatientProfile
// @Router /patients [post]
func (ph *PatientsHandler) createPatient(w http.ResponseWriter, r *http.Request) {
	var dto model.PatientProfile
	decodeHTTPBody(r, &dto)

	p := model.Patient{}
	p.PatientProfile = dto

	err := ph.PatientsRepository.Create(&p)
	if err != nil {
		panic("failed to connect database")
	}

	json.NewEncoder(w).Encode(&p.PatientProfile)

	log.Println("createPatients")
}

// @Summary Update Patient Profile
// @Description Update Patient Profile
// @Tags patients
// @ID patients-profile-update
// @Accept  json
// @Param patientProfile body model.PatientProfile true "Update an existing patient profile"
// @Param id path int true "Patient profile identifier that has to match w/ the identifier within the body contract"
// @Produce  json
// @Success 200 {object} model.PatientProfile
// @Failure 400 {object} model.ErrorResponse
// @Router /patients/{id} [put]
func (ph *PatientsHandler) updatePatient(w http.ResponseWriter, r *http.Request) {
	// urlSegments := strings.Split(r.URL.Path, "/")
	// idFromPath, _ := strconv.ParseUint(urlSegments[len(urlSegments)-1], 10, 64)
	vars := mux.Vars(r)
	idFromPath, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errorResponse := model.ErrorResponse{
			Code:    model.INALID_INPUT,
			Message: "Missing resource identifier in path",
			TraceID: uuid.New(),
		}
		writeError(w, &errorResponse)
		return
	}
	var dto model.PatientProfile
	decodeHTTPBody(r, &dto)
	if uint(idFromPath) != dto.ID {
		errorResponse := model.ErrorResponse{
			Code:    model.RESOURCE_IDENTIFIER_MISMATCH,
			Message: "Resource identifier in path should match with the body contract.ID",
			TraceID: uuid.New(),
		}
		writeError(w, &errorResponse)
		return
	}

	// p := model.Patient{}
	// p.PatientProfile = dto

	// err = ph.PatientsRepository.Create(&p)
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// json.NewEncoder(w).Encode(&p.PatientProfile)

	log.Println("updatePatient")
}

func decodeHTTPBody(r *http.Request, structure interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Print(err)
	}
}

func writeError(w http.ResponseWriter, errorResponse *model.ErrorResponse) {
	log.Println(errorResponse)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorResponse)
}
