package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	g := gin.Default()
	g.POST("/login", func(ctx *gin.Context) {
		// var body Body
		bodyStr, _ := ioutil.ReadAll(ctx.Request.Body)
		log.Println("[login] body = ", string(bodyStr))
		b64String, _ := base64.StdEncoding.DecodeString(string(bodyStr))
		rawIn := json.RawMessage(b64String)
		bodyBytes, err := rawIn.MarshalJSON()
		if err != nil {
			log.Println("[login] erorr = ", err)
		}
		log.Println("[login] bodyBytes = ", string(bodyBytes))
	})

	g.Run()
}
