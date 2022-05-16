package middleware

import (
	"blockchain-deal-hunter/api/utility"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func IpBlacklisthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipBlacklist := strings.Split(os.Getenv("ip_blacklist"), ",")
		for _, host := range  ipBlacklist{
			if c.ClientIP() == host {
				c.Set("responseErr", utility.NotAllowedError{errors.New(fmt.Sprintf("%v, not allowed ip", c.ClientIP()))})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
