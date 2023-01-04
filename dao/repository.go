package dao

import (
	"findings/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IRepositoryDao interface {
	Create(repository model.Repository) error
	FindAll() ([]model.Repository, error)
	Update(repository model.Repository) error
	Delete(name string) error
	UpdateScanDetails(repository model.Repository, status model.StatusType) error
	FindRepositoryByStatus(status model.StatusType) ([]model.Repository, error)
	SaveFindings(finding *model.Finding, scanDetail *model.ScanDetail) error
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDao() *repository {
	return &repository{NewDatabaseInstance().GetConnection()}
}

func (dao *repository) Create(repo model.Repository) error {
	dao.db.Debug().Create(&repo)
	return nil
}

func (dao *repository) FindAll() ([]model.Repository, error) {
	var repos []model.Repository
	dao.db.Find(&repos)
	return repos, nil
}

func (dao *repository) Update(repository model.Repository) error {
	dao.db.Debug().Where("name = ?", repository.Name).Updates(repository)
	return nil
}

func (dao *repository) Delete(name string) error {
	var repository model.Repository
	dao.db.Debug().Preload("ScanDetails").Where("name = ?", name).Find(&repository).Select(clause.Associations).Delete(&repository)
	return nil
}

func (dao *repository) UpdateScanDetails(repository model.Repository, status model.StatusType) error {
	scanDetails := &model.ScanDetail{
		Status: status,
	}
	repository.ScanDetails = append(repository.ScanDetails, *scanDetails)
	dao.db.Debug().Where("name = ?", repository.Name).Find(&repository).Updates(scanDetails)
	return nil
}

func (dao *repository) FindRepositoryByStatus(status model.StatusType) ([]model.Repository, error) {
	var repos []model.Repository
	dao.db.Debug().Preload("ScanDetails", "status = ?", status).Preload("ScanDetails.Findings").Find(&repos)
	return repos, nil
}

func (dao *repository) SaveFindings(finding *model.Finding, scanDetail *model.ScanDetail) error {
	scanDetail.Findings = append(scanDetail.Findings, *finding)
	dao.db.Debug().Save(scanDetail)
	return nil
}
