package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
	if err := initDB(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/login", loginGetHandler)
	r.POST("/login", loginPostHandler)
	r.GET("/register", registerGetHandler)
	r.POST("/register", registerPostHandler)

	auth := r.Group("/")
	auth.Use(authMiddleware())
	{
		auth.GET("/", indexHandler)
		auth.POST("/todos", createTodoHandler)
		auth.POST("/todos/:id/delete", deleteTodoHandler)
		auth.GET("/logout", logoutHandler)
	}

	r.Run(":8080")
}
