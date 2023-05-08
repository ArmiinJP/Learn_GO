package requestParam

type CreateTask struct{
	Command			string
	ValueCommand	ValuesCreateTask
}

type ValuesCreateTask struct {
	Title		string
	DueDate		string
	CategoryID	int	
	UserID		int
}

type ListTask struct{
	Command			string
	ValueCommand	ValuesListTask
}

type ValuesListTask struct {
	UserID		int	
}



type ListTodayTask struct{
	Command			string
	ValueCommand	ValueslistTodayTask
}
type ValueslistTodayTask struct{}



type ListSpecificDayTask struct{
	Command			string
	ValueCommand	ValuesListSpecificDayTask
}
type ValuesListSpecificDayTask struct{}



type EditTask struct{
	Command			string
	ValueCommand	ValuesEditTask
}
type ValuesEditTask struct{}



type ChangeStatusTask struct{
	Command			string
	ValueCommand	ValuesChangeStatusTask
}
type ValuesChangeStatusTask struct{}
