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
	cs, err := cli.CollectionService.Featured(ctx, &pixels.PageParams{
		Page: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Collection Title:", cs.Collections[0].Title)
	fmt.Println("RateLimit:", cli.LastRateLimit.Limit)
}
