package contract

import "session_2_2/entity"

type UserStoreWrite interface{
	Save(u entity.User)
}

type UserStoreRead interface{
	Read() []entity.User
}
