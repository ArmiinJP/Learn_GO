package requestParam

type CreateCategory struct{
	Command			string
	ValueCommand	ValuesCreateCategory
}
type ValuesCreateCategory struct{
	Title string
	Color string
	UserID int
}

type ListCategory struct{
	Command			string
	ValueCommand	ValuesListCategory
}
type ValuesListCategory struct{
	UserID int
}



type EditCategory struct{
	Command			string
	ValueCommand	ValuesEditCategory
}
type ValuesEditCategory struct{}


