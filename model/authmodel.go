package model

type AuthModel struct{
	UserID   string `json:"userid"`
	UserName string `json:"userame"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthInterface interface{
	Login()(string,error)
	Register()(string,error)
}