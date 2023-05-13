package param

type Request struct {
	Command string
	Values RequestValue
}

type RequestValue struct{
	CreatedTaskRequest
	ListedTaskRequest
}

type CreatedTaskRequest struct{
	Title              string
	DueDate            string
	CategoryID         int
}

type ListedTaskRequest struct{
	Test              string

}