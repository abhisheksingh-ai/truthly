package service

import (
	"context"
	"log/slog"
	"strconv"
	"truthly/internals/dto"
	"truthly/internals/repository"
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type FeedService interface {
	GetFeed(ctx context.Context, limit int, cursor string) (*dto.FeedResponseDto, error)
}

type feedService struct {
	feedRepo    repository.FeedRepository
	commentRepo repository.CommentRepository
	logger      *slog.Logger
}

func GetNewFeedService(fr repository.FeedRepository, commentRepo repository.CommentRepository, l *slog.Logger) FeedService {
	return &feedService{
		feedRepo:    fr,
		commentRepo: commentRepo,
		logger:      l,
	}
}

func (fs *feedService) GetFeed(ctx context.Context, limit int, cursor string) (*dto.FeedResponseDto, error) {

	//1. Get FeedRows
	rows, nextCursor, hasMore, err := fs.feedRepo.GetFeedItems(ctx, limit, cursor)
	if err != nil {
		fs.logger.Error(err.Error())
		return nil, err
	}

	//2. Image Ids collect
	imageIds := make([]string, 0, len(rows))
	for _, r := range rows {
		imageIds = append(imageIds, r.ImageId)
	}

	// 3. final response dto for Feed
	items := make([]dto.FeedItemDto, 0)

	for _, r := range rows {
		items = append(items, dto.FeedItemDto{
			ImageId:   r.ImageId,
			ImageUrl:  r.ImageUrl,
			Caption:   r.Caption,
			CreatedAt: r.CreatedAt,

			UserName: r.UserName,
			UserId:   r.UserId,

			Location: dto.LocationDto{
				City:    r.City,
				State:   r.State,
				Country: r.Country,
			},
			Analytics: dto.AnalyticsDto{
				Like:    mustAtoi(r.LikeCount),
				Comment: mustAtoi(r.CommentCount),
				Share:   mustAtoi(r.ShareCount),
			},
		})
	}

	fs.logger.Info("Feed responseded")

	return &dto.FeedResponseDto{
		Items: items,
		Pagination: dto.PaginationDto{
			NextCursor: nextCursor,
			HasMore:    hasMore,
		},
	}, nil
}
