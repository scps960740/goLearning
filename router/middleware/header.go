package middleware

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func NoCache(context *gin.Context) {
	context.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	context.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	context.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	context.Next()
}

func Options(context *gin.Context) {
	if context.Request.Method != "OPTIONS" {
		context.Next()
	} else {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		context.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		context.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		context.Header("Content-Type", "application/json")
		context.AbortWithStatus(200)
	}
}

func Secure(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("X-Frame-Options", "DENY")
	context.Header("X-Content-Type-Options", "nosniff")
	context.Header("X-XSS-Protection", "1; mode=block")
	if context.Request.TLS != nil {
		context.Header("Strict-Transport-Security", "max-age=31536000")
	}
}
