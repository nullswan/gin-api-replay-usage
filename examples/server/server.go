package main

import (
	"encoding/json"
	"fmt"

	apiUsage "github.com/c3b5aw/gin-api-replay-usage"
	"github.com/gin-gonic/gin"
)

func display(e *apiUsage.ReplayData) {
	prettyPrint, err := json.MarshalIndent(*e, "", "\t")
	if err != nil {
		fmt.Printf("unable to indent replaydata: %s\n", err)
	}
	_ = prettyPrint
	fmt.Print(string(prettyPrint) + "\n")
}

func main() {
	router := gin.Default()

	router.Use(
		apiUsage.Register(display),
	)

	router.Any("/*path", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	router.Run(":3000")
}
