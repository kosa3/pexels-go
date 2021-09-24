package main

import (
	"context"
	"fmt"
	"github.com/kosa3/pexels-go"
	"log"
	"os"
)

func main() {
	cli := pexels.NewClient(os.Args[1])
	ctx := context.Background()
	ps, err := cli.PhotoService.Search(ctx, &pexels.PhotoParams{
		Query: "people",
		Page:  2,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Source Medium URL:", ps.Photos[0].Src.Medium)
	fmt.Println("RateLimit:", cli.LastRateLimit.Limit)
}
