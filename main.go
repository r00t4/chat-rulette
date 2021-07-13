package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/r00t4/chat-rulette/internal"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]

	m, err := internal.NewManager(args[0])
	if err != nil {
		log.Fatal(err)
		return
	}

	r := gin.Default()
	r.POST("/updates/:token", func(context *gin.Context) {
		token := context.Param("token")

		upd := &tgbotapi.Update{}
		if err := context.BindJSON(upd); err != nil {
			context.JSON(400, err)
			return
		}
		if err := m.OnUpdate(token, upd); err != nil {
			context.JSON(400, err)
			return
		}
		context.JSON(200, "success")
		return
	})

	if err := http.ListenAndServe(":8443", r); err != nil {
		fmt.Println("ListenAndServe error:", err)
		return
	}
}
