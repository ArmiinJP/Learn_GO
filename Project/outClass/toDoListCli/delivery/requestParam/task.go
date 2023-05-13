package requestParam

import (
	"time"
)

type ValuesCreateTask struct {
	Title		string
	DueDate		string
	CategoryID	int	
	UserID		int
}

type ValuesListTask struct {
	UserID		int	
}

type ValueslistTodayTask struct{
	UserID		int	
	Date		time.Time
}


type ValuesListSpecificDayTask struct{
	UserID		int	
	Date		time.Time
}


type ValuesEditTask struct{
	ID         int
	Title      string
	DueDate    string
	CategoryID   int
	IsComplete bool
	UserID     int
}


type ValuesChangeStatusTask struct{
	ID         int
	IsComplete bool
	UserID     int
}
