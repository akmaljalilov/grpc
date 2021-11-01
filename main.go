package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"velox/gRPC/recommendation"
)

func main() {
	channel, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln( err)
	}
	client := recommendation.NewRecommendationsClient(channel)
	res, err := client.Recommend(context.Background(), &recommendation.RecommendationRequest{
		UserId:     1,
		Category:   recommendation.BookCategory_SCIENCE_FICTION,
		MaxResults: 3,
	})
	if err != nil {
		log.Fatalln( err)
	}
	log.Printf("outliers at: %v", res.Recommendations)
}
