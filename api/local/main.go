package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"itmx/routers"
)

func main() {
	//utilities.SetEnv()
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(ginsession.New())

	r.POST("/customers", routers.CreateCustomer)
	r.GET("/customers/:id", routers.GetCustomer)
	r.GET("/customers", routers.GetCustomers)
	r.PUT("/customers/:id", routers.UpdateCustomer)
	r.DELETE("/customers/:id", routers.DeleteCustomer)
	_ = r.Run(":8080")
}
