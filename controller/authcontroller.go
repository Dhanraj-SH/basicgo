package controller

import (
"net/http"


)

type AuthController struct{
authRepository model.AuthInterface
}

func NewAuthController(ai model.AuthInterface) *AuthController {
return &AuthController{
authRepository: ai,
}
}

func(ac *AuthController) RegisterController(w http.ResponseWriter , r *http.Request){
mes , err := ac.authRepository.Register(r)
if err != nil {
w.WriteHeader(http.StatusBadRequest)
w.Write([]byte(err.Error()))
return
}
w.WriteHeader(http.StatusOK)
w.Write([]byte(mes))
}

