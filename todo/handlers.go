package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func indexHandler(c *gin.Context) {
	userID, _ := getSession(c)
	rows, err := db.Query("SELECT id, title, deadline FROM todos WHERE user_id = ? ORDER BY deadline", userID)
	if err != nil {
		c.String(http.StatusInternalServerError, "DB error")
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		rows.Scan(&t.ID, &t.Title, &t.Deadline)
		todos = append(todos, t)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{"Todos": todos})
}

func loginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func loginPostHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var id int64
	var hash string
	err := db.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&id, &hash)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{"Error": "Неверный логин или пароль"})
		return
	}

	newSession(c, id)
	c.Redirect(http.StatusFound, "/")
}

func registerGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func registerPostHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Заполните все поля"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server error")
		return
	}

	result, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, string(hash))
	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Имя пользователя уже занято"})
		return
	}

	id, _ := result.LastInsertId()
	newSession(c, id)
	c.Redirect(http.StatusFound, "/")
}

func createTodoHandler(c *gin.Context) {
	userID, _ := getSession(c)
	title := c.PostForm("title")
	deadline := c.PostForm("deadline")

	if title != "" && deadline != "" {
		db.Exec("INSERT INTO todos (user_id, title, deadline) VALUES (?, ?, ?)", userID, title, deadline)
	}
	c.Redirect(http.StatusFound, "/")
}

func deleteTodoHandler(c *gin.Context) {
	userID, _ := getSession(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}
	db.Exec("DELETE FROM todos WHERE id = ? AND user_id = ?", id, userID)
	c.Redirect(http.StatusFound, "/")
}

func logoutHandler(c *gin.Context) {
	deleteSession(c)
	c.Redirect(http.StatusFound, "/login")
}
