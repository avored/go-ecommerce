package controllers

import (
	// "github.com/avored/go-ecommerce/dto"
	// "github.com/avored/go-ecommerce/services"

	// "github.com/gin-gonic/gin"
)


// func LoginHandler(loginService services.LoginService, jWtService services.JWTService) LoginController {
// 	return &loginController{
// 		loginService: loginService,
// 		jWtService:   jWtService,
// 	}
// }

// func (controller *loginController) Login(ctx *gin.Context) string {
// 	var credential dto.LoginCredentials
// 	err := ctx.ShouldBind(&credential)
// 	if err != nil {
// 		return "no data found"
// 	}
// 	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
// 	if isUserAuthenticated {
// 		return controller.jWtService.GenerateToken(credential.Email, true)

// 	}
// 	return ""
// }
