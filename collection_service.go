package pixels

import (
	"context"
	"fmt"
)

type CollectionService interface {
	Featured(ctx context.Context, params *PageParams) (*CollectionResponse, error)
	Get(ctx context.Context, params *PageParams) (*CollectionResponse, error)
	Find(ctx context.Context, collectionId string, params *CollectionParams) (*CollectionMediaResponse, error)
}

type collectionService struct {
	cli *Client
}

func (s *collectionService) Featured(ctx context.Context, params *PageParams) (*CollectionResponse, error) {
	var sr CollectionResponse
	if err := s.cli.get(ctx, "collections/featured", StructToMap(params), &sr); err != nil {
		return nil, fmt.Errorf("GET featured failed: %w", err)
	}

	return &sr, nil
}

func (s *collectionService) Get(ctx context.Context, params *PageParams) (*CollectionResponse, error) {
	var cr CollectionResponse
	if err := s.cli.get(ctx, "collections", StructToMap(params), &cr); err != nil {
		return nil, fmt.Errorf("GET get failed: %w", err)
	}

	return &cr, nil
}

func (s *collectionService) Find(ctx context.Context, collectionId string, params *CollectionParams) (*CollectionMediaResponse, error) {
	var cmr CollectionMediaResponse
	if err := s.cli.get(ctx, "collections/"+collectionId, StructToMap(params), &cmr); err != nil {
		return nil, fmt.Errorf("GET find failed: %w", err)
	}

	return &cmr, nil
}
