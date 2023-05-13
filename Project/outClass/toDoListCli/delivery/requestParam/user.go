package requestParam

type ValuesRegisterUser struct{
	Email string
	Password string
}

type ValuesLoginUser struct{
	Email string
	Password string	
}

type ValuesWhoami struct{
	UserID int
}