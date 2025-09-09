package transport

import (
	_ "backend/docs"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Listener() {
	r := gin.Default()
	fmt.Println(">>> Listener started")

	r.GET("/", Root)
	r.GET("/echo", Echo)
	r.POST("/parsing", Parsing)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("Listentig port 8080...")
	r.Run(":8080")
}
