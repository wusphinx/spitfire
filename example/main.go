package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	pprof.Register(r)

	r.GET("/test", func(c *gin.Context) {
		var a []int
		for i := 0; i < 100000; i++ {
			a = append(a, i)
		}
	})

	r.Run()
}
