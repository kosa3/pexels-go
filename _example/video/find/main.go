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
	vs, err := cli.VideoService.Find(ctx, 2499611)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Video Link:", vs.VideoFiles[0].Link)
}
