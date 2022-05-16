package controller

import (
	"blockchain-deal-hunter/api/dto"
	"github.com/gin-gonic/gin"
)

type User struct {
	Controller
}

func (ctrl *User) Login(c *gin.Context) {
	input := &dto.LoginInput{}
	if err := ctrl.DefaultGetValidParams(c, input); err != nil {
		c.Set("responseErr", err)
		return
	}
	//if err := api.BindingValidParams(c); err != nil {
	//	middleware.ResponseError(c, 2001, err)
	//	return
	//}
	//if api.Username == "admin" && api.Password == "123456" {
	//	session := sessions.Default(c)
	//	session.Set("user", api.Username)
	//	session.Save()
	//	middleware.ResponseSuccess(c, "")
	//	return
	//}
	//middleware.ResponseError(c, 2002, errors.New("账号或密码错误"))
	//return
}

func (ctrl *User) Login2(c *gin.Context) {
	input := &dto.LoginInput{}
	if err := ctrl.DefaultGetValidParams(c, input); err != nil {
		c.Set("responseMsg", "呵呵2")
		c.Set("responseErr", err)
		return
	}
	//if err := api.BindingValidParams(c); err != nil {
	//	middleware.ResponseError(c, 2001, err)
	//	return
	//}
	//if api.Username == "admin" && api.Password == "123456" {
	//	session := sessions.Default(c)
	//	session.Set("user", api.Username)
	//	session.Save()
	//	middleware.ResponseSuccess(c, "")
	//	return
	//}
	//middleware.ResponseError(c, 2002, errors.New("账号或密码错误"))
	//return
}