package main

import (
	"github.com/ALCooper12/boggle/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/solveBoggleBoard", handlers.HandleBoggleBoardSubmission)
	router.Run(":8080")
}
