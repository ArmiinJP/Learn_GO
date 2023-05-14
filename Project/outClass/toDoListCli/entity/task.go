package entity

type Task struct {
	TaskID     int
	Title      string
	DueDate    string
	CategoryID int
	IsComplete bool
	UserID     int
}
