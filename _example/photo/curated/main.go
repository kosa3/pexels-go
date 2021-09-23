package main

import (
	"context"
	"fmt"
	pixels "github.com/kosa3/pexels-go"
	"log"
	"os"
)

func main() {
	cli := pixels.NewClient(os.Args[1])
	ctx := context.Background()
	ps, err := cli.PhotoService.Curated(ctx, &pixels.PageParams{
		Page: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Source Medium URL:", ps.Photos[0].Src.Medium)
	fmt.Println("RateLimit:", cli.LastRateLimit.Limit)
}
