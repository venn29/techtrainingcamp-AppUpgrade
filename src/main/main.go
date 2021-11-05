package main

import (
	"github.com/gin-gonic/gin"
)
/*
var (
	MaxWorker int = runtime.NumCPU()
	MaxQueue int = 1000
)
*/

func main() {
	r := gin.Default()
	// 配置Handler
	customizeouter(r)
	r.Run()
}
