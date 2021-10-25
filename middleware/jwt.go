package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"im/util"
	"log"
	"strings"
	"time"
)

const tokenPrefix = "bearer%20"

func Jwt() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		log.Printf("query %v\n", time.Now())
		token, b := ctx.GetQuery("token")
		log.Printf("token %v\n",token)
		if !b {
			ctx.AbortWithStatusJSON(200, gin.H{
				"code": 40001,
				"msg":  "token不存在，请登录",
			})
			return
		}
		prefix := strings.HasPrefix(token,tokenPrefix)
		if!prefix {
			ctx.AbortWithStatusJSON(200, gin.H{
				"code": 40001,
				"msg":  "token不合法，请登录",
			})
			return
		}
		token=strings.TrimLeft(token, tokenPrefix)
		jwtToken, err := util.ValidateJwtToken(token)
		if err!=nil {
			ctx.AbortWithStatusJSON(200, gin.H{
				"code": 40001,
				"msg":  "token无效，请登录",
			})
			return
		}
		claims := jwtToken.Claims.(jwt.MapClaims)
		mobile,ok := claims["mobile"]
		if !ok {
			ctx.AbortWithStatusJSON(200, gin.H{
				"code": 40001,
				"msg":  "token无效，请登录",
			})
			return
		}
		log.Println("mobile %v \n",mobile)
	}
}
