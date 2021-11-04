package main

import (
	"gRPC/grpc_test/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {

	server := gin.Default()
	server.GET("/", hello)
	server.GET("/testwork/:x/:y", GetFibonacci)
	server.Run(":8888")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "What's up, been working on it for while but it works")
}

func GetFibonacci(c *gin.Context) {

	x := c.Param("x")
	y := c.Param("y")
	numX, _ := strconv.Atoi(x)
	numY, _ := strconv.Atoi(y)

	result, err := database.RedisClientGet(numX, numY)
	if err == nil {
		c.JSON(http.StatusOK, result)
		return
	}

	countRes := GetFibonacciArr(numX, numY)
	c.JSON(http.StatusOK, countRes)

	database.RedisClientSet(numX, numY, countRes)

	c.JSON(http.StatusOK, result)
}

func GetFibonacciArr(s, n int) []int {

	t1 := 0
	t2 := 1
	nextTerm := 0

	arr := make([]int, 0)
	for i := 1; i <= n+1; i++ {
		if i == 1 {
			continue
		}
		if i == 2 {
			arr = append(arr, t2)
			continue
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		arr = append(arr, nextTerm)
	}
	fibArr := arr[(s - 1):]

	return fibArr

}
