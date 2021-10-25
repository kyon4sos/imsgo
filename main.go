package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"im/engine"
	"im/handler"
	"im/middleware"
	"im/online"
	"im/router"
	"im/util"
	"log"
)



func  main()  {
	e:=gin.Default()

	//e.Use(middleware.Jwt())
	app := engine.NewApp()
 	app.Use(middleware.Login)
	app.AddHandler(&handler.Login{}).
		AddHandler(&handler.SingleChat{}).
		AddHandler(&handler.GroupChat{})
	go app.Run()
	e.GET("/ws", func(ctx *gin.Context) {
		token, b := ctx.GetQuery("token")
		if !b {
			return
		}
		log.Printf("token %v\n",token)
		jwtToken, err := util.ValidateJwtToken(token)
		if err!=nil {
			log.Printf("token err %v \n",err.Error())
			return
		}
		mapClaims := jwtToken.Claims.(jwt.MapClaims)
		log.Printf("token claims %v \n",mapClaims)
		w:=ctx.Writer
		r:=ctx.Request
		log.Println("ws")
		client := app.NewServer(w, r, nil)
		user:=&online.User{
			Id: mapClaims["sub"].(string),
			Client: client,
		}
		online.SetOnlineUser(user)
	})
	router.Register(e)
	err := e.Run(":9000")
	if err != nil {
		log.Panicf("newApp run err %v \n",err.Error())
		return
	}
}
