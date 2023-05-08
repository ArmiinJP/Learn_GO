package requestParam

type RegisterUser struct{
	Command			string
	ValueCommand	ValuesRegisterUser
}
type ValuesRegisterUser struct{
	Email string
	Password string
}

type LoginUser struct{
	Command			string
	ValueCommand	ValuesLoginUser
}
type ValuesLoginUser struct{
	Email string
	Password string	
}

type Whoami struct{
	Command			string
	ValueCommand	ValuesWhoami
}
type ValuesWhoami struct{
	UserID int
}