package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	proto "github.com/utevo/gRPC-API/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseInt(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseInt(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		req := &proto.Request{A: a, B: b}
		response, err := client.Add(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error: " + err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
	})

	g.GET("/mul/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseInt(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseInt(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		req := &proto.Request{A: a, B: b}
		response, err := client.Multiply(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error: " + err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Not able run the server: %v", err)
	}
}
