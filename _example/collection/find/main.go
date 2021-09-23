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
	cs, err := cli.CollectionService.Find(ctx, "ji3n6vt", &pixels.CollectionParams{
		Type: "photos",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Collection:", cs.ID)
	fmt.Println("RateLimit:", cli.LastRateLimit.Limit)
}
