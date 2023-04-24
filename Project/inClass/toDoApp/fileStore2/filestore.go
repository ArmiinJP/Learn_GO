package filestore2

import "session_2_2/entity"

type FileStorage2 struct {}

func (f FileStorage2) Save(entity.User) {}
func (f FileStorage2) Read() ([]entity.User) {
	return []entity.User{}
}
