package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/plapi/config"
	"github.com/plapi/controllers"
	"github.com/plapi/routes"
	"log"
)

func main()  {
	config.Connect()
	controllers.CreateSubjectTable()
	fmt.Println("Successfully setup database")
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":8080"))
}
