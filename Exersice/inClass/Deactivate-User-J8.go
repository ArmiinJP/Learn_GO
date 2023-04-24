package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	var t1 = &Teacher{
		ID:       2,
		IsActive: true,
	}

	s := Student{
		ID:       3,
		Name:     "",
		Email:    "",
		IsActive: true,
		JoinDate: time.Time{},
	}



	err := DeactivateUser2(t1)
	err2 := DeactivateUser2(s)
	err3 := DeactivateUser(t1, &s, nil)

	fmt.Println(err, err2, err3)

}

type Teacher struct {
	ID       uint
	Name     string
	Email    string
	IsActive bool
	Salary   uint
	Grade    string
}

func (t *Teacher) Deactivate() error {
	// if t.IsActive == false
	if !t.IsActive {
		return errors.New("the teacher is deactivated already")
	}

	t.IsActive = false

	return nil
}

type Student struct {
	ID       uint
	Name     string
	Email    string
	IsActive bool
	JoinDate time.Time
}

func (s Student) Deactivate() error {
	// if s.IsActive == false
	if !s.IsActive {
		return errors.New("the student is deactivated already")
	}

	s.IsActive = false

	return nil
}

type Staff struct {
	ID       uint
	Name     string
	Position string
	Status   string
}

func (s *Staff) Deactivate() error {
	if s.Status == "active" {
		return errors.New("the student is deactivated already")
	}

	s.Status = "deactive"

	return nil
}

func DeactivateUser(t *Teacher, s *Student, staff *Staff) error {
	if t != nil {
		// directly return t.Deactivate() error
		//return t.Deactivate()

		// return new error
		//err := t.Deactivate()
		//if err != nil {
		//	errors.New("can't deactivate user")
		//}

		// wrap error
		err := t.Deactivate()
		if err != nil {
			return fmt.Errorf("can't deactivate user: %w", err)
		}
	}

	if s != nil {
		return s.Deactivate()
	}

	if staff != nil {
		return staff.Deactivate()
	}

	return nil
}

func DeactivateUser2(u deAc) error{
	if err := u.Deactivate(); err != nil{
		return fmt.Errorf("user not deactivate: %w",err)
	}

	return nil
}

type deAc interface{
	Deactivate() error
}