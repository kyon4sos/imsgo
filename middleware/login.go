package middleware

import (
	"im/engine"
	"log"
)

func Login() func(ctx engine.Context) {
	log.Println("login 中间件")
	return func(ctx engine.Context) {
			//ctx.Cancel()
	}
}
