package main

import (
	"com.hyosakura/LovesAsuna/Go/server"
	"net/http"
)

func main() {
	r := server.New()
	r.GET("/", func(c *server.Context) {
		c.String(http.StatusOK, "Hello World\n")
	})

	// index out of range for testing Recovery()
	r.GET("/panic", func(c *server.Context) {
		names := []string{"LovesAsuna"}
		c.String(http.StatusOK, names[100])
	})
	r.Use(server.Recovery())
	r.Run(":8080")
}
