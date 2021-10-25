package router

import (
	"github.com/gin-gonic/gin"
	"im/db"
	"im/dto"
	"im/helper"
	"im/model"
	"im/util"
	"log"
)

func login(ctx *gin.Context) {
	var login dto.LoginDto
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		helper.Err(ctx, 4000, "username password参数不正确")
		return
	}
	var user model.ChatUser
	first := db.GetDb().Where("username", login.Username).First(&user)
	if first.Error != nil {
		helper.Err(ctx, 4001, "用户名或密码错误")
		return
	}
	if user.Password != login.Password {
		helper.Err(ctx, 4001, "用户名或密码错误")
		return
	}
	token, err := util.CreateJwtToken(user.Username)
	if err != nil {
		helper.Err(ctx, 5000, "系统错误")
		return
	}
	res := make(map[string]interface{})
	res["username"] = user.Username
	res["token"] = token
	helper.Ok(ctx, res)
}

func register(ctx *gin.Context) {
	var register dto.RegisterDto
	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		helper.Err(ctx, 4000, "username ,password参数不正确")
		return
	}
	var user model.ChatUser
	first := db.GetDb().Where("username", register.Username).First(&user)
	if first.Error == nil {
		helper.Err(ctx, 4000, "用户名已存在")
		return
	}
	user.Username = register.Username
	user.Password = register.Password
	create := db.GetDb().Create(&user)
	if create.Error != nil {
		log.Printf("err %v \n", create.Error)
		helper.Err(ctx, 5000, "系统错误")
		return
	}

	token, err := util.CreateJwtToken(user.Username)
	if err != nil {
		helper.Err(ctx, 5000, "系统错误")
		return
	}
	res := make(map[string]interface{})
	res["username"] = user.Username
	res["token"] = token
	helper.Ok(ctx, res)
}
