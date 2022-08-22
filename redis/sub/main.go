package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	subscriber := redisClient.Subscribe(ctx, "stockfolio")

	user := User{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			panic(err)
		}

		fmt.Println("메시지가" + msg.Channel + " 채널에서 왔습니다.\n\n")
		fmt.Printf("%+v\n\n", user)
	}
}
