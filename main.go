package main

import (
	"github.com/naman-dave/chkdin/router"
)

func main() {
	r := router.NewRoutes()
	r.Run()
}
