package filestore2

import "todolistapp/entity"

type FileStorage2 struct {}

func (f FileStorage2) Save(entity.User) {}
func (f FileStorage2) Read() ([]entity.User) {
	return []entity.User{}
}
