package controller

import (
	"findings/model"
	"findings/service"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

var (
	lock = sync.Mutex{}
)

type repository struct {
	repositoryService service.IRepositoryService
}

func NewRepositoryController() *repository {
	return &repository{service.NewRepositoryService()}
}

// Create Repository
//
//	@Summary	Add a new repository to the store
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		message	body		model.Repository true	"Repository Data"
//	@Success	200		{string}	json			 "ok"
//	@Failure	400		{object}	string	 		 "We need name and url !!"
//	@Router		/repository [post]
func (controller *repository) Create(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	v := &model.Repository{}
	c.Bind(v)
	controller.repositoryService.Create(*v)
	return c.JSON(http.StatusCreated, v)
}

// FindAll Repository
//
//	@Summary	Find all repositories from the store
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Success	200		{string}	json			 "ok"
//	@Failure	400		{object}	string	 		 "We need name and url !!"
//	@Router		/repository [get]
func (controller *repository) FindAll(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	repos, _ := controller.repositoryService.FindAll()
	return c.JSON(http.StatusOK, repos)
}

// FindByStatus
//
//	@Summary	Find repositories by status
//	@Tags		admin
//	@Accept		json
//  @Param   	status  path     string     false  "string enums"       Enums(QUEUED,INPROGRESS,SUCCESS,FAILURE)
//	@Produce	json
//	@Success	200		{string}	json			 "ok"
//	@Failure	400		{object}	string	 		 "We need name and url !!"
//	@Router		/repository/{status} [get]
func (controller *repository) FindByStatus(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	status := model.StatusTypeMap[c.Param("status")]
	fmt.Println(status)
	repos, _ := controller.repositoryService.FindByStatus(status)
	return c.JSON(http.StatusOK, repos)
}

// Update Repository
//
//	@Summary	Update details of exiting repository
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		message	body		model.Repository true	"Repository Data"
//	@Success	200		{string}	json			 "ok"
//	@Failure	400		{object}	string	 		 "We need name and url !!"
//	@Router		/repository [put]
func (controller *repository) Update(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	v := &model.Repository{}
	c.Bind(v)
	controller.repositoryService.Update(*v)
	return c.JSON(http.StatusOK, v)
}

//  Queue Scan
//
//	@Summary	Queue exiting repository for scanning
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		name	path		string			 true	"Repository name"
//	@Success	200		{string}	json			 "ok"
//	@Failure	400		{object}	string	 		 "We need name !!"
//	@Router		/repository/scan/{name} [put]
func (controller *repository) InitScan(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	name := c.Param("name")
	controller.repositoryService.UpdateScanDetails(name, model.QUEUED)
	return c.JSON(http.StatusOK, name)
}

// Delete Repository
//
//	@Summary	Delete exiting repository
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		name	path		string			 true	"Repository name"
//	@Success	200		{string}	json			 "ok"
//	@Failure	400		{object}	string	 		 "We need name and url !!"
//	@Router		/repository/{name} [delete]
func (controller *repository) Delete(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	name := c.Param("name")
	controller.repositoryService.Delete(name)
	return c.NoContent(http.StatusNoContent)
}

//  Force Execute Scan
//
//	@Summary	Execute exiting repository for scanning
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Success	200		{string}	json			 "ok"
//	@Failure	400		{object}	string	 		 "We need name !!"
//	@Router		/repository/scan [put]
func (controller *repository) ExecuteScanner(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	controller.repositoryService.ExecuteScanner()
	return c.JSON(http.StatusOK, nil)
}
