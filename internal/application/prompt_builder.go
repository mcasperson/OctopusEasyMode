package application

import (
	"github.com/gin-gonic/gin"
)

func PromptBuilderHandler(c *gin.Context) {
	c.File("html/prompt.html")
}
