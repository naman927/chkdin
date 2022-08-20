package contollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naman-dave/chkdin/modals"
)

// register user
func APIRegister(c *gin.Context) {
	resp := map[string]interface{}{}

	user, err := modals.NewUser()
	if err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if err := c.ShouldBindJSON(user); err != nil {
		resp["data"] = nil
		resp["error"] = "unprocessable entity"
		resp["message"] = "unmarshal error"
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	if err := user.ValidateUserForAuth(); err != nil {
		resp["data"] = nil
		resp["error"] = err.Error()
		resp["message"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	if err = user.CreateUser(); err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp["data"] = map[string]string{"token": user.Token}
	resp["error"] = nil
	resp["message"] = "Successfully registered, Use that token for auth purpose"
	c.JSON(http.StatusOK, resp)
}

func APILogin(c *gin.Context) {
	resp := map[string]interface{}{}

	user, err := modals.NewUser()
	if err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if err := c.ShouldBindJSON(user); err != nil {
		resp["data"] = nil
		resp["error"] = "unprocessable entity"
		resp["message"] = "unmarshal error"
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	if err := user.ValidateUserForAuth(); err != nil {
		resp["data"] = nil
		resp["error"] = err.Error()
		resp["message"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	if err = user.Login(); err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp["data"] = map[string]string{"token": user.Token}
	resp["error"] = nil
	resp["message"] = "Successfully Logged in, Use that token for auth purpose"
	c.JSON(http.StatusOK, resp)
}
