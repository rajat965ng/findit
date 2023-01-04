package service

import (
	"findings/dao"
	"findings/model"
	"strconv"
	"strings"
)

type IRepositoryService interface {
	Create(repo model.Repository) error
	FindAll() ([]model.Repository, error)
	Update(repository model.Repository) error
	Delete(name string) error
	UpdateScanDetails(repository string, status model.StatusType) error
	ExecuteScanner()
	FindByStatus(status model.StatusType) ([]model.Repository, error)
}

type repository struct {
	repositoryDao dao.IRepositoryDao
}

func NewRepositoryService() *repository {
	return &repository{dao.NewRepositoryDao()}
}

func (svc *repository) Create(repo model.Repository) error {
	return svc.repositoryDao.Create(repo)
}

func (svc *repository) FindAll() ([]model.Repository, error) {
	return svc.repositoryDao.FindAll()
}

func (svc *repository) Update(repository model.Repository) error {
	return svc.repositoryDao.Update(repository)
}

func (svc *repository) Delete(name string) error {
	return svc.repositoryDao.Delete(name)
}

func (svc *repository) UpdateScanDetails(name string, status model.StatusType) error {
	return svc.repositoryDao.UpdateScanDetails(model.Repository{Name: name}, status)
}

func (svc *repository) ExecuteScanner() {
	repos, err := svc.FindByStatus(model.QUEUED)
	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		for _, scanDetail := range repo.ScanDetails {
			if scanDetail.Status == model.QUEUED {
				gitService := NewGitService(&repo, "./dir")
				_, err := gitService.GitClone()
				if err != nil {
					panic(err)
				}
				results, _ := gitService.GrepText([]string{"private-key", "public-key"})
				for _, result := range results {
					r := strings.Split(result, ":")
					u64, err := strconv.ParseUint(r[1], 10, 32)
					if err != nil {
						panic(err)
					}
					finding := model.Finding{Type: "SAST", RuleId: "G001", Path: r[0], LineNumber: uint(u64), Severity: "High", Description: r[2]}
					scanDetail.Status = model.SUCCESS
					svc.repositoryDao.SaveFindings(&finding, &scanDetail)
				}
				gitService.CleanUp()
			}
		}
	}

}

func (svc *repository) FindByStatus(status model.StatusType) ([]model.Repository, error) {
	return svc.repositoryDao.FindRepositoryByStatus(status)
}
