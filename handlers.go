package handlers

import (
	"fmt"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/kartheekvadde/policy/proto"
	"google.golang.org/grpc"
)

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

/*
//CreatePolicyHandler performs Addition Operation
func CreatePolicyVariableHandler(ctx *gin.Context) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewPolicyServiceClient(conn)

	req := &proto.RequestPolicyVariable{PolicyKey:"1", PolicyVal:"L3", PolicyName: "L3", Status: "Active"}
	if response, err := client.CreatePolicyVariable(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Status),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
*/
