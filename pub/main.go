package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/unsafe-risk/golam"
	"log"
	"net/http"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

var ctx = context.Background()

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	g := golam.New()
	g.POST("/", PostString)
	log.Print("Start server")
	err := g.StartWithLocalAddr(":3000")
	if err != nil {
		log.Print(err)
	}
}

func PostString(c golam.Context) error {
	var user User
	body := c.Request().Body
	data := []byte(body)
	json.Unmarshal(data, &user)
	payload, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	if err := redisClient.Publish(ctx, "stockfolio", payload).Err(); err != nil {
		log.Print(err)
	}
	return c.String(http.StatusOK, "good")
}
