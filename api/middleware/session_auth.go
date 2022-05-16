package middleware

import (
	"blockchain-deal-hunter/api/utility"
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if name,ok:=session.Get("user").(string);!ok||name==""{
			c.Set("responseErr", utility.NoLoginError{errors.New("Not logged in")})
			c.Abort()
			return
		}
		c.Next()
	}
}
