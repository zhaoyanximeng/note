package main

import (
	"github.com/gin-gonic/gin"
	"note/redislock/lib"
	"time"
)

var a = 1

func main() {
	r := gin.New()
	r.Use(Middleware)
	r.GET("/", func(c *gin.Context) {
		l := lib.NewLockerWithTTL("lock1",3 * time.Second).Lock()
		defer l.UnLock()
		if c.Query("t") != "" {
			time.Sleep(10*time.Second)
		}
		a++

		c.JSON(200,gin.H{"message":a})
	})

	r.Run(":8080")
}

func Middleware(c *gin.Context) {
	defer func() {
		if e := recover() ; e != nil {
			c.AbortWithStatusJSON(400,gin.H{"message":e})
		}
	}()

	c.Next()
}
