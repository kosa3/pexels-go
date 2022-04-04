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
	cs, err := cli.CollectionService.Get(ctx, &pexels.PageParams{
		Page: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Collection:", cs.Collections)
}
