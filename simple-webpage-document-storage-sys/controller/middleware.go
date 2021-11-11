package controller

import (
	"github.com/gin-gonic/gin"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/token"
)

// SetHeader sets the http header and:
// 1. allows CORS
// 2. sets the allowed methods
// 3. sets the allowed headers
// 4. sets the access-control-expose-headers
// 5. sets the allowed credentials
func SetHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// DecodeToken decodes the token and put the extracted uid into the context
func DecodeToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr := c.Request.Header.Get("Authorization")

		uid, b := token.DecodeToken(tokenStr)

		if b {
			c.Set(common.TokenUid, uid)
		}
		c.Next()
	}
}
