package pexels

import (
	"context"
	"fmt"
	"strconv"
)

type PhotoService interface {
	Search(ctx context.Context, params *PhotoParams) (*SearchPhotoResponse, error)
	Curated(ctx context.Context, params *PageParams) (*CuratedResponse, error)
	Find(ctx context.Context, photoId int) (*Photo, error)
}

type photoService struct {
	cli *Client
}

func (s *photoService) Search(ctx context.Context, params *PhotoParams) (*SearchPhotoResponse, error) {
	var sr SearchPhotoResponse
	if err := s.cli.get(ctx, "search", StructToMap(params), &sr); err != nil {
		return nil, fmt.Errorf("GET search failed: %s", err)
	}

	return &sr, nil
}

func (s *photoService) Curated(ctx context.Context, params *PageParams) (*CuratedResponse, error) {
	var cr CuratedResponse
	if err := s.cli.get(ctx, "curated", StructToMap(params), &cr); err != nil {
		return nil, fmt.Errorf("GET curated failed: %s", err)
	}

	return &cr, nil
}

func (s *photoService) Find(ctx context.Context, photoId int) (*Photo, error) {
	var p Photo
	if err := s.cli.get(ctx, "photos/"+strconv.Itoa(photoId), nil, &p); err != nil {
		return nil, fmt.Errorf("GET curated failed: %s", err)
	}

	return &p, nil
}
