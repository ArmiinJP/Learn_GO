package categoryservice

import (
	"encoding/json"
	"fmt"
	"todolist/delivery/requestParam"
	"todolist/delivery/responseParam"
	"todolist/entity"
)

type CategoryRepo interface {
	Create(category entity.Category) error
	List(userID int) ([]entity.Category, error)
	Edit(task entity.Category) error
	DoesUserhaveCategory(userID, categoryID int) error
	NewCategoryIDGenerateForUser(userID int) (int, error) 	
}

type Service struct {
	repository CategoryRepo
}

func New(cr CategoryRepo) Service {
	return Service{repository: cr}
}

func (s Service) CreateCategoryRequest(request requestParam.ValuesCreateCategory) (responseParam.Response, error) {
	
	NewCategoryIDGenerate, nErr := s.repository.NewCategoryIDGenerateForUser(request.UserID)
	if nErr != nil{
		return responseParam.Response{StatusCode: 404, Message: "Faild to create category", Data: []byte{}}, fmt.Errorf("%s", nErr.Error())
	}
	
	cErr := s.repository.Create(entity.Category{
		CategoryID: NewCategoryIDGenerate,
		Title:      request.Title,
		Color:      request.Color,
		UserID:     request.UserID,
	})

	if cErr != nil {
		return responseParam.Response{StatusCode: 500, Message: "Failed to Create Category", Data: []byte{}}, fmt.Errorf("error Creating Category: %s", cErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Create Category Successfully", Data: []byte{}}, nil
}

func (s Service) ListCategoryRequest(request requestParam.ValuesListCategory) (responseParam.Response, error) {

	userCategory, lErr := s.repository.List(request.UserID)
	if lErr != nil {

		return responseParam.Response{StatusCode: 500, Message: "Failed to List Category", Data: []byte{}}, fmt.Errorf("error Listing Category: %s", lErr.Error())
	}

	//fmt.Println(userTask)
	data, mErr := json.Marshal(userCategory)
	if mErr != nil {

		return responseParam.Response{StatusCode: 500, Message: "Failed to List Category", Data: []byte{}}, fmt.Errorf("error Marshaling Response: %s", lErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Return List Category Successfully", Data: data}, nil
}

func (s Service) EditCategoryRequst(request requestParam.ValuesEditCategory) (responseParam.Response, error) {
	
	dErr := s.repository.DoesUserhaveCategory(request.UserID, request.CategoryID)
	if dErr != nil{
		return responseParam.Response{StatusCode: 404, Message: "Faild to Edit category", Data: []byte{}}, fmt.Errorf("user doesn't have category ID: %d error: %s",request.CategoryID ,dErr.Error())
	}

	eErr := s.repository.Edit(entity.Category{
		CategoryID: request.CategoryID,
		Title:      request.Title,
		Color:      request.Color,
		UserID:     request.UserID,
	})

	if eErr != nil {
		return responseParam.Response{StatusCode: 500, Message: "Failed to Edit Category", Data: []byte{}}, fmt.Errorf("error Editing Category: %s", eErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Edit Category Successfully", Data: []byte{}}, nil
}