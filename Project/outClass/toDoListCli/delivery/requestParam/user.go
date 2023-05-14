package requestParam

type ValuesRegisterUser struct{
	Email string
	Password string
}

type ValuesLoginUser struct{
	Email string
	Password string	
	RemoteAddr string
}

type ValuesWhoami struct{
	RemoteAddr string
}