package main

import "github.com/gin-gonic/gin"

func main() {
	container := gin.Default()
	_ = container.Run(":8080")
}
