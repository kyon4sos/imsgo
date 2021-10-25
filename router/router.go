package router

import (
	"github.com/gin-gonic/gin"
	"im/helper"
	"im/online"
	"im/util"
)

func Register(e *gin.Engine)  {
	e.GET("/online-user", GetOnlineUser)
	e.GET("/api/token",GetToken)
	e.POST("/api/register",register)
	e.POST("/api/login",login)
}

func GetOnlineUser(ctx *gin.Context) {
	c:= online.GetClientById("abc123")
	helper.Ok(ctx,c.Conn.RemoteAddr())
}

func GetToken(ctx *gin.Context)  {
	token, _ := util.CreateJwtToken("abc123")
	helper.Ok(ctx,token)
}