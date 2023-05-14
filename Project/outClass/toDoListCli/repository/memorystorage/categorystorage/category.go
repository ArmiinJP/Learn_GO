package categorystorage

import (
	"fmt"
	"todolist/constant"
	"todolist/entity"
)

type CategoryStorage struct {
	categories map[int][]entity.Category
}

func New() CategoryStorage {
	return CategoryStorage{
		categories: make(map[int][]entity.Category),
	}
}

func (c *CategoryStorage) Create(category entity.Category) error {

	c.categories[category.UserID] = append(c.categories[category.UserID], category)
	return nil
}

func (c CategoryStorage) List(userID int) ([]entity.Category, error) {
	var tmpCategories = []entity.Category{}
	for user, userCategory := range c.categories {
		if user == userID {
			return userCategory, nil
		}
	}

	return tmpCategories, nil
}

func (c *CategoryStorage) Edit(category entity.Category) error {
	for user, userCategory := range c.categories {
		if user == category.UserID {
			for i, catValue := range userCategory {
				if catValue.CategoryID == category.CategoryID {
					if category.Color != "" {
						catValue.Color = category.Color
					}
					if category.Title != "" {
						catValue.Title = category.Title
					}
					c.categories[user][i] = catValue
				}
			}
		}
	}
	return nil
}

func (c CategoryStorage) DoesUserhaveCategory(userID, categoryID int) error {
	for user, userCategory := range c.categories {
		if user == userID {
			for _, category := range userCategory {
				if category.CategoryID == categoryID {
					return nil
				}
			}
		}
	}

	return fmt.Errorf("user: %d doesn't have category ID: %d", userID, categoryID)
}

func (c CategoryStorage) NewCategoryIDGenerateForUser(userID int) (int, error) {
	for user, userCategory := range c.categories {
		if user == userID {
			newID := constant.MinCategoryIDEachUser + len(userCategory) + 1
			if newID < constant.MaxCategoryIDEachUser {
				return newID, nil
			} else {
				return 0, fmt.Errorf("user dosen't allow add new category, category capacity is full")
			}
		}
	}
	return 0, nil
}
