package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sercand/kuberesolver"
	"github.com/yangwenmai/examples/grpc-k8s-example/pb"
	"google.golang.org/grpc"
)

func main() {
	// Set up HTTP server
	r := gin.Default()
	r.GET("/gcd/:a/:b", func(c *gin.Context) {
		// Connect to GCD service
		kuberesolver.RegisterInCluster()

		conn, err := grpc.Dial("kubernetes:///gcd-service.default:3000", grpc.WithInsecure())
		defer conn.Close()
		log.Println(conn)
		log.Println(conn.GetState())
		// conn, err := grpc.Dial("gcd-service:3000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Dial failed: %v", err)
		}

		gcdClient := pb.NewGCDServiceClient(conn)
		// Parse parameters
		a, err := strconv.ParseUint(c.Param("a"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
			return
		}
		b, err := strconv.ParseUint(c.Param("b"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter B"})
			return
		}
		// Call GCD service
		req := &pb.GCDRequest{A: a, B: b}
		if res, err := gcdClient.Compute(c, req); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result":  fmt.Sprint(res.Result),
				"version": res.Version,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Run HTTP server
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
