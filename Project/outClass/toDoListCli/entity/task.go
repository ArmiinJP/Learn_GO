package entity

type Task struct {
	ID         int
	Title      string
	DueDate    string
	CategoryID   int
	IsComplete bool
	UserID     int
}
