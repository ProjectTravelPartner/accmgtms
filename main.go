package main

import (
	"fmt"

	"github.com/ProjectTravelPartner/accmgtms/api"
	"github.com/ProjectTravelPartner/commonclient"
	"github.com/ProjectTravelPartner/dbclient"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hey, there!")
	port := commonclient.CliPort(3000)
	dbclient.Init()
	defer dbclient.Close()
	router := gin.New()
	group := router.Group("accmgts")

	group.GET("/account/:id", api.AccountGet)
	group.POST("/account", api.AccountCreate)
	group.GET("/login", api.AccountLogin)
	router.Run(fmt.Sprintf(":%v", port))
	//group.DELETE("/account/:id", api.AccountDelete)
}
