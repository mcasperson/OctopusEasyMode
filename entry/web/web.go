package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mcasperson/OctopusEasyMode/internal/application"
)

func main() {
	r := gin.Default()

	r.GET("/", application.PromptBuilderHandler)

	port := ":8080"
	fmt.Printf("Starting Gin web server on port %s\n", port)
	r.Run(port)
}
