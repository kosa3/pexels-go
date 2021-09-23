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
	vs, err := cli.VideoService.Find(ctx, 2499611)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Video Link:", vs.VideoFiles[0].Link)
	fmt.Println("RateLimit:", cli.LastRateLimit.Limit)
}