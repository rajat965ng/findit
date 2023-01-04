package model

import (
	"gorm.io/gorm"
)

// Repository model info
// @Description Git Repository information
// @Description with name and url
type Repository struct {
	gorm.Model
	Name        string       `json:"name,omitempty"` // Name this is name of repository
	Url         string       `json:"url,omitempty"`  // Url this is url of repository
	ScanDetails []ScanDetail `json:"scanDetails,omitempty" gorm:"foreignKey:RepositoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StatusType string

const (
	QUEUED     StatusType = "Queued"
	INPROGRESS StatusType = "In Progress"
	SUCCESS    StatusType = "Success"
	FAILURE    StatusType = "Failure"
)

var (
	StatusTypeMap = map[string]StatusType{
		"QUEUED":     QUEUED,
		"INPROGRESS": INPROGRESS,
		"SUCCESS":    SUCCESS,
		"FAILURE":    FAILURE,
	}
)

// Repository scan details
// @Description Git Repository information
// @Description with name and url
type ScanDetail struct {
	gorm.Model
	Status       StatusType `gorm:"type:status_type"`
	Description  string     `json:"description,omitempty"`
	RepositoryId uint
	Findings     []Finding `json:"findings,omitempty" gorm:"foreignKey:ScanDetailId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Finding struct {
	gorm.Model
	Type         string
	RuleId       string
	Path         string
	LineNumber   uint
	Description  string
	Severity     string
	ScanDetailId uint
}
