package pexels

import (
	"context"
	"fmt"
	"strconv"
)

type VideoService interface {
	Search(ctx context.Context, params *VideoParams) (*SearchVideoResponse, error)
	Popular(ctx context.Context, params *PopularParams) (*PopularVideoResponse, error)
	Find(ctx context.Context, videoId int) (*Video, error)
}

type videoService struct {
	cli *Client
}

func (s *videoService) Search(ctx context.Context, params *VideoParams) (*SearchVideoResponse, error) {
	var sv SearchVideoResponse
	if err := s.cli.get(ctx, "videos/search", StructToMap(params), &sv); err != nil {
		return nil, fmt.Errorf("GET search failed: %s", err)
	}

	return &sv, nil
}

func (s *videoService) Popular(ctx context.Context, params *PopularParams) (*PopularVideoResponse, error) {
	var pr PopularVideoResponse
	if err := s.cli.get(ctx, "videos/popular", StructToMap(params), &pr); err != nil {
		return nil, fmt.Errorf("GET popular failed: %s", err)
	}

	return &pr, nil
}

func (s *videoService) Find(ctx context.Context, videoId int) (*Video, error) {
	var v Video
	if err := s.cli.get(ctx, "videos/videos/"+strconv.Itoa(videoId), nil, &v); err != nil {
		return nil, fmt.Errorf("GET videos failed: %s", err)
	}

	return &v, nil
}
