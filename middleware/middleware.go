// Package middleware
// @author tabuyos
// @since 2023/6/30
// @description middleware
package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1: in go
		// parentContext := context.WithValue(ctx.Request.Context(), "user", "gtp-user")
		// ctx.Request = ctx.Request.WithContext(parentContext)
		// ctx.Next()

		// 2: in gin
		ctx.Set("user", "gtp-user")
	}
}

func CheckRBAC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
