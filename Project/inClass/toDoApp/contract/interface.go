package contract

import "todolistapp/entity"

type UserStoreWrite interface{
	Save(u entity.User)
}

type UserStoreRead interface{
	Read() []entity.User
}
