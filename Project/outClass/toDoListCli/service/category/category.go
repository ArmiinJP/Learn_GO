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
	DoesUserhaveCategory(userID, categoryID int) bool
}

type Service struct {
	repository CategoryRepo
}

func New(cr CategoryRepo) Service {
	return Service{repository: cr}
}

func (s Service) CreateCategoryRequest(request requestParam.ValuesCreateCategory) (responseParam.Response, error) {
	cErr := s.repository.Create(entity.Category{
		ID:     2,
		Title:  request.Title,
		Color:  request.Color,
		UserID: request.UserID,
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
	eErr := s.repository.Edit(entity.Category{
		ID:         request.ID,
		Title:      request.Title,
		Color:    	request.Color,
		UserID:     request.UserID,
	})

	if eErr != nil {
		return responseParam.Response{StatusCode: 500, Message: "Failed to Edit Category", Data: []byte{}}, fmt.Errorf("error Editing Category: %s", eErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Edit Category Successfully", Data: []byte{}}, nil
}

func (s Service) DoesUserhaveCategory(UserID int, CategoryID int) error{
	dErr := s.DoesUserhaveCategory(UserID, CategoryID)
	return dErr
}

