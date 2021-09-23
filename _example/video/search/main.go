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
	vs, err := cli.VideoService.Search(ctx, &pixels.VideoParams{
		Query: "nature",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Video Link:", vs.Videos[0].VideoFiles[0].Link)
	fmt.Println("RateLimit:", cli.LastRateLimit.Limit)
}