package main

import (
	"crypto/rand"
	"encoding/hex"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	sessions   = map[string]int64{}
	sessionsMu sync.Mutex
)

func newSession(c *gin.Context, userID int64) {
	b := make([]byte, 16)
	rand.Read(b)
	token := hex.EncodeToString(b)

	sessionsMu.Lock()
	sessions[token] = userID
	sessionsMu.Unlock()

	c.SetCookie("session", token, 86400, "/", "", false, true)
}

func getSession(c *gin.Context) (int64, bool) {
	token, err := c.Cookie("session")
	if err != nil {
		return 0, false
	}
	sessionsMu.Lock()
	userID, ok := sessions[token]
	sessionsMu.Unlock()
	return userID, ok
}

func deleteSession(c *gin.Context) {
	token, err := c.Cookie("session")
	if err == nil {
		sessionsMu.Lock()
		delete(sessions, token)
		sessionsMu.Unlock()
	}
	c.SetCookie("session", "", -1, "/", "", false, true)
}
