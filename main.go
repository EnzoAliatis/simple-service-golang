package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	fmt.Println("Si ves esto desde un contenedor docker, eres cahingo")
	c.JSON(http.StatusOK, "Si ves esto eres chingon")
}

func panicGo(c *gin.Context) {
	fmt.Println("Vamos en entrar en panico en 5 seg")
	c.JSON(http.StatusOK, "por cuestiones de seguridad inabilitamos este endpoint")
	// time.AfterFunc(5*time.Second, func() {
	// 	panic("Entrando en panico")
	// })
}

func main() {
	r := gin.Default()

	r.GET("/hello", hello)
	r.GET("/panic", panicGo)
	r.Run()
}
