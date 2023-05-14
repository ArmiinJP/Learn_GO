package requestParam

type ValuesCreateCategory struct {
	Title  string
	Color  string
	UserID int
}

type ValuesListCategory struct {
	UserID int
}

type ValuesEditCategory struct {
	CategoryID int
	Title      string
	Color      string
	UserID     int
}
