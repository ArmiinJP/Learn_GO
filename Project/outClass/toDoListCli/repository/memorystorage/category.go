package memorystorage

import (
	"fmt"
	"todolist/entity"
)



type CategoryStorage struct {
	categories map[int][]entity.Category
}

func (c *CategoryStorage) Create(category entity.Category) error{
	
	c.categories[category.UserID] = append(c.categories[category.UserID], category)
	return nil
}

func (c CategoryStorage) List(userID int) ([]entity.Category, error){
	var tmpCategories = []entity.Category{}
	for user, userCategory := range c.categories{
		if user == userID{
			return userCategory, nil
		}
	}

	return tmpCategories, nil
}

func (c *CategoryStorage) Edit(category entity.Category) error{
	for user, userCategory := range c.categories{
		if user == category.UserID{
			for i, catValue := range(userCategory){
				if catValue.ID == category.ID{
					if category.Color != ""{
						catValue.Color = category.Color
					}
					if category.Title != ""{
						catValue.Title = category.Title
					}
					c.categories[user][i] = catValue
				}
			}
		}
	}
	return nil
}

func (c CategoryStorage) DoesUserhaveCategory(userID, categoryID int) bool

func (c CategoryStorage) isExist(CategoryID int) bool {
	for _, v := range c.categories{
		if v.ID == CategoryID{
			return true
		}
	}
	return false
}