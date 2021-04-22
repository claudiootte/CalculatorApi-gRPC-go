package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/claudiootte/CalculatorApi-gRPC-go/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	g := gin.Default()
	g.GET("/sum/:num01/:num02", func(ctx *gin.Context) {
		num01, err := strconv.ParseInt(ctx.Param("num01"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": http.StatusBadRequest})
			return
		}

		num02, err := strconv.ParseInt(ctx.Param("num02"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": http.StatusBadRequest})
			return
		}

		req := &proto.Request{Num01: int64(num01), Num02: int64(num02)}
		if response, err := client.Sum(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/subtract/:num01/:num02", func(ctx *gin.Context) {
		num01, err := strconv.ParseInt(ctx.Param("num01"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": http.StatusBadRequest})
			return
		}
		num02, err := strconv.ParseInt(ctx.Param("num02"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": http.StatusBadRequest})
			return
		}
		req := &proto.Request{Num01: int64(num01), Num02: int64(num02)}

		if response, err := client.Subtract(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/multiply/:num01/:num02", func(ctx *gin.Context) {
		num01, err := strconv.ParseInt(ctx.Param("num01"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": http.StatusBadRequest})
			return
		}
		num02, err := strconv.ParseInt(ctx.Param("num02"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": http.StatusBadRequest})
			return
		}
		req := &proto.Request{Num01: int64(num01), Num02: int64(num02)}

		if response, err := client.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/divide/:num01/:num02", func(ctx *gin.Context) {
		num01, err := strconv.ParseInt(ctx.Param("num01"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": http.StatusBadRequest})
			return
		}
		num02, err := strconv.ParseInt(ctx.Param("num02"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": http.StatusBadRequest})
			return
		}

		if num02 == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"error": "Can't divide by zero",
			})
		} else if num01/num02 == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"error": "This API only work with integer results",
			})

		} else {

			req := &proto.Request{Num01: int64(num01), Num02: int64(num02)}

			if response, err := client.Divide(ctx, req); err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"result": fmt.Sprint(response.Result),
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
