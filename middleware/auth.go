package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naman-dave/chkdin/modals"
)

const (
	NoToken = "auth-token not provided"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("auth-token")
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"data":    "",
				"error":   NoToken,
				"message": NoToken,
			})
			c.Abort()
		}
		user, err := modals.NewUser()
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"data":    "",
				"error":   "something went wrong",
				"message": "something went wrong",
			})
			c.Abort()
		}
		if err := user.AuthUser(token); err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"data":    "",
				"error":   err.Error(),
				"message": err.Error(),
			})
			c.Abort()
		}

		// here we are adding the user to context so that we can use it for user specific func
		c.Set("user", user)
		c.Next()
	}
}
