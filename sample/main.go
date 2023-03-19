package main

import (
	"log"
	"sample/router"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	router.TestRouter(g)
	router.VideoRouter(g)

	log.Fatal(g.Run(":8000"))
}
