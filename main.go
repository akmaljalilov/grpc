package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	"velox/gRPC/recommendation"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	go func() {
		sec := 0
		for _ = range ticker.C {
			sec++
			fmt.Println("Sec: ", sec)

		}
	}()
	channel, err := grpc.DialContext(ctx, "localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := recommendation.NewRecommendationsClient(channel)
	res, err := client.Recommend(ctx, &recommendation.RecommendationRequest{
		UserId:     1,
		Category:   recommendation.BookCategory_SCIENCE_FICTION,
		MaxResults: 3,
	})
	ticker.Stop()
	cancel()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("outliers at: %v", res.Recommendations)
}
