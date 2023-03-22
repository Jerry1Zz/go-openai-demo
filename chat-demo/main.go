package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"os"
)

func main(){
	key := os.Getenv("OPEN_AI_KEY")
	r := gin.Default()
	r.GET("gerAns",func(c *gin.Context){
		comment := c.Query("comment")
		ans := GetAns(key,comment)
		c.JSON(200,gin.H){
			"answer" : ans,
		}
	})
	err := r.Run()
	if err != nil{
		return
	}

}

func GetAns(key, comment string) string{
	client := openai.NewClient(key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: comment,
				},
			},
		},
	)

	if err !=nil{
		return ""
	}
	return resp.Choices[0].Message.Content
}

