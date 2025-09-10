package transport

import (
	_ "backend/docs"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Listener() {
	r := gin.Default()
	log.Println("Listener started")

	r.GET("/", Root)
	r.POST("/echo", Echo)
	r.POST("/parsing", Parsing)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Listentig port 8080...")
	r.Run(":8080")
}
