package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naman-dave/chkdin/contollers"
	"github.com/naman-dave/chkdin/middleware"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")
	r.addAuthRoute(v1)
	v1.Use(middleware.Auth())
	r.addPersonsRoute(v1)

	return r
}

func (r routes) Run() error {
	return r.router.Run()
}

func (r routes) addPersonsRoute(rg *gin.RouterGroup) {
	person := rg.Group("/person")
	{
		person.GET("/:id", contollers.APIGetPersonDetails)
		person.POST("", contollers.APICreatePerson)
		person.PUT("", contollers.APIEditPerson)
		person.DELETE("/:id", contollers.APIDeletePerson)
	}
}

func (r routes) addAuthRoute(rg *gin.RouterGroup) {
	rg.POST("/register", contollers.APIRegister)
	rg.POST("/login", contollers.APILogin)
}
