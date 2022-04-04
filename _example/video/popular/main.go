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
	vs, err := cli.VideoService.Popular(ctx, &pexels.PopularParams{
		MaxDuration: 4,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Video Link:", vs.Videos[0].VideoFiles[0].Link)
}
