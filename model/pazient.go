package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// PazientProfile contains anagrafic data
type PazientProfile struct {
	gorm.Model
	FirstName   string
	LastName    string
	Profession  string
	Address     string
	City        string
	ZipCode     string
	Province    string
	Email       string
	Mobile      string
	Phone       string
	DateOfBirth *time.Time
}

// Pazient contains the full patient data including medical history
type Pazient struct {
	PazientProfile
	RemoteHistories []RemoteHistory
}

// RemoteHistory holds the overall history of pazient
type RemoteHistory struct {
	gorm.Model
	PazientID     uint
	HistoryKindID uint
	HistoryKind   HistoryKind `gorm:"association_save_reference:true;save_associations:false"`
	Date          *time.Time
	Description   string
}

// Consultation contains the general information for the medical consultation
type Consultation struct {
	gorm.Model
	PazientID      uint
	Date           *time.Time
	InitialProblem string
}

// NextHistory contains the overall information collected during each consultation
type NextHistory struct {
	gorm.Model
	ConsultationID uint
	FirstTime      string
	Typology       string
	Localization   string
	Irradiation    string
	OnsetPeriod    string
	Duration       string
	Familiarity    string
	OtherTherapies string
	Various        string
}

// Examination contains the overall information collected during each consultation
type Examination struct {
	gorm.Model
	ConsultationID    uint
	ExaminationKindID uint
	Date              *time.Time
	Description       string
}

// Treatment contains the information related the medical treatment
type Treatment struct {
	gorm.Model
	ConsultationID uint
	Date           *time.Time
	Description    string
}

// Assessment contains the cousulation result
type Assessment struct {
	gorm.Model
	ConsultationID uint
	Structural     string
	CranioSacral   string
	AkOrthodontic  string
}

// HistoryKind contains the list of possible RemoteHistory kinds
type HistoryKind struct {
	ID          uint `gorm:"primary_key"`
	Description string
}

// ExaminationKind contains the list of possible medical examination kinds
type ExaminationKind struct {
	ID          uint `gorm:"primary_key"`
	Description string
}

// Province contains the list of possible RemoteHistory kinds
type Province struct {
	ID          uint   `gorm:"primary_key"`
	Code        string `gorm:"size:2"`
	Description string
}
