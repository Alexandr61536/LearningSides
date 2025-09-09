package transport

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Go is listening")
}
