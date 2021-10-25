package helper

import "github.com/gin-gonic/gin"


func Ok(ctx *gin.Context,data interface{})  {
	ctx.JSON(200,gin.H{
		"code":2000,
		"msg":"ok",
		"data":data,
	})
}
func Err(ctx *gin.Context,code int ,msg string )  {
	ctx.JSON(200,gin.H{
		"code":code,
		"msg":msg,
	})
}
