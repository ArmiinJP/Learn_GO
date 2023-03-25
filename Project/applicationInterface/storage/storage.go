package storage

import (
 	"applicationInterface/user"
)


// slice storage
type InMemorySlice struct{
	users []user.User
}

func (m *InMemorySlice) CreateUser(u user.User){
	m.users = append(m.users, u)
}

func (m *InMemorySlice) ListUser() []user.User{
	return m.users
}

func (m *InMemorySlice) GetUserByID(id uint) user.User{
	for _, v := range m.users{
		if v.ID == id{
			return v
		}
	}
	return user.User{}
}

// map storage
type InMemoryMap struct{
	users  map[uint]user.User
}

func (m *InMemoryMap) CreateUser(u user.User){
	if m.users == nil{
		m.users = make(map[uint]user.User)
	}
	m.users[u.ID] = u
}

func (m *InMemoryMap) ListUser() []user.User{
	
	//tmpSlice := make([]user.User, len(m.users))
	//fmt.Println(tmpSlice, cap(tmpSlice), len(tmpSlice))
	
	var tmpSlice []user.User
	for _, v := range m.users{
		tmpSlice = append(tmpSlice, v)
	}

	return tmpSlice
}

func (m *InMemoryMap) GetUserByID(id uint) user.User{
	return m.users[id]
}
