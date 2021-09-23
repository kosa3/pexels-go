# Pexels API Client for Go

`kosa3/pexels-go` is [Pexels API](https://www.pexels.com/api/documentation/) Client for Go.
[Pexels](https://www.pexels.com/) is the best free stock photos & videos shared by talented creators.


## Install

```
$ go get -u github.com/kosa3/pexels-go
```

## Examples

See [example](_example) directory.

```go
func main() {
	cli := pixels.NewClient(os.Args[1])
	ctx := context.Background()
	ps, err := cli.PhotoService.Search(ctx, &pixels.PhotoParams{
		Query: "people",
		Page: 2,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Source Medium URL:", ps.Photos[0].Src.Medium)
	fmt.Println("RateLimit:", cli.LastRateLimit.Limit)
}
```

## Supported API

### Photo

|                 Endpoint                | HTTP Method |
|-----------------------------------------|:-----------:|
|/v1/search                              | GET         |
|/v1/curated                             | GET         |
|/v1/photos/{id}                         | GET         |



### Video

|                 Endpoint                | HTTP Method |
|-----------------------------------------|:-----------:|
|/videos/search                              | GET         |
|/videos/popular                             | GET         |
|/videos/videos/{id}                         | GET         |

### Collection

|                 Endpoint                | HTTP Method |
|-----------------------------------------|:-----------:|
|/v1/collections/featured                  | GET         |
|/v1/collections                           | GET         |
|/v1/collections/{id}                      | GET         |