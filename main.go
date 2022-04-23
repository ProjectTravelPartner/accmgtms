package main

import (
	"fmt"

	"github.com/ProjectTravelPartner/commonclient"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hey, there!")
	port := commonclient.CliPort(3000)
	router := gin.New()
	router.Run(fmt.Sprintf(":%v", port))
}
