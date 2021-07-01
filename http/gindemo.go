package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const keyRequestId = "requestId"

func main() {
	engine := gin.Default()

	logger, e := zap.NewProduction()
	if e != nil {
		panic(e)
	}

	//file, e := os.OpenFile("./gin.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	//
	//if e != nil {
	//	panic(e)
	//}
	//
	//logger := log.New(file, "[GinDemo] ", log.Lshortfile|log.Ldate|log.Ltime)

	engine.Use(func(context *gin.Context) {
		start := time.Now()

		context.Next()

		logger.Info("Incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("spend", time.Since(start)))
		//logger.Printf("Incoming request, path: %s\n", context.Request.URL.Path)

	}, func(context *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		context.Set(keyRequestId, rand.Int())
		context.Next()
	})

	engine.Handle(http.MethodGet, "/ping", func(context *gin.Context) {
		h := gin.H{
			"message": "pong",
		}

		if rid, exists := context.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}

		context.JSON(200, h)
	})

	engine.GET("/hello", func(context *gin.Context) {
		s := "string"
		if rid, exists := context.Get(keyRequestId); exists {
			s += fmt.Sprintf(", requestId: %d", rid)
		}
		context.String(200, s)
	})

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
