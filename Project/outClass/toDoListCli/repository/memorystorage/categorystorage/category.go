package categorystorage

import (
	"fmt"
	"todolist/constant"
	"todolist/entity"
)

type storage struct {
	categories map[int][]entity.Category
}

func New() storage {
	return storage{
		categories: make(map[int][]entity.Category),
	}
}

func (c *storage) Create(category entity.Category) error {

	c.categories[category.UserID] = append(c.categories[category.UserID], category)
	return nil
}

func (c storage) List(userID int) ([]entity.Category, error) {
	
	for user, userCategory := range c.categories {
		if user == userID {
			return userCategory, nil
		}
	}

	return []entity.Category{}, fmt.Errorf("user doesn't have any Cateogry")
}

func (c *storage) Edit(category entity.Category) error {
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

func (c storage) DoesUserhaveCategory(userID, categoryID int) error {
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

func (c storage) NewCategoryIDGenerateForUser(userID int) (int, error) {
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
	return 0, fmt.Errorf("error generating category id")
}

func (c storage) Print() {
	fmt.Println("All Category is: -------------------------")
	for user, userCategory := range  c.categories{
		for _, category := range userCategory {
			fmt.Printf("User ID: %d\nCategory ID: %d\nCategory Title: %s\nCategory Color: %s\n",
				user, category.CategoryID, category.Color, category.Title)
		}
	}
}
