package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartheekvadde/policy/proto"
	"google.golang.org/grpc"
)

func main() {
	g := gin.Default()
	g.GET("/policy", CreatePolicyHandler)
	// g.GET("/policy-variable", handlers.CreatePolicyVariableHandler)

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

//CreatePolicyHandler performs Addition Operation
func CreatePolicyHandler(ctx *gin.Context) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewPolicyServiceClient(conn)

	req := &proto.RequestPolicy{Id: 1, OrgId: 2, UserId: 2, PolicyName: "L3", Status: "Active", Description: "L3-Policy"}
	if response, err := client.CreatePolicy(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Status),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
