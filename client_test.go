package pixels_test

import (
	"context"
	pixels "github.com/kosa3/pexels-go"
	"testing"
)

func TestNewClientInvalidToken(t *testing.T) {
	cli := pixels.NewClient("invalid")
	ctx := context.Background()
	_, err := cli.PhotoService.Curated(ctx, nil)
	if err == nil {
		t.Fatal("unexpected error", err)
	}
}
