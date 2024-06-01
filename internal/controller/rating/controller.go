package user

import (
	"context"
	"fmt"
	"github.com/JMURv/unona/ratings/pkg/model"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"time"
)

const cacheKey = "ratings:%v"

type BrokerRepository interface{}

type CacheRepository interface {
	Get(ctx context.Context, key string) (*model.Rating, error)
	GetRatingValue(ctx context.Context, key string) (float32, error)
	Set(ctx context.Context, t time.Duration, key string, r *model.Rating) error
	SetRatingValue(ctx context.Context, t time.Duration, key string, r float32) error
	Delete(ctx context.Context, key string) error
}

type ratingRepository interface {
	GetUserRating(ctx context.Context, userUUID uuid.UUID) (float32, error)
	CreateReport(ctx context.Context, rating *model.Rating) (*model.Rating, error)
	UpdateReport(ctx context.Context, ratingID uint, newData *model.Rating) (*model.Rating, error)
	DeleteReport(ctx context.Context, ratingID uint) error
}

type Controller struct {
	repo   ratingRepository
	cache  CacheRepository
	broker BrokerRepository
}

func New(repo ratingRepository, cache CacheRepository, broker BrokerRepository) *Controller {
	return &Controller{
		repo:   repo,
		cache:  cache,
		broker: broker,
	}
}

func (c *Controller) GetUserRating(ctx context.Context, userUUID uuid.UUID) (float32, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ratings.GetUserRating.ctrl")
	ctx = opentracing.ContextWithSpan(ctx, span)
	defer span.Finish()

	cached, err := c.cache.GetRatingValue(ctx, fmt.Sprintf(cacheKey, userUUID))
	if err == nil {
		return cached, nil
	}

	res, err := c.repo.GetUserRating(ctx, userUUID)
	if err != nil {
		return 0, err
	}

	err = c.cache.SetRatingValue(ctx, time.Hour, fmt.Sprintf(cacheKey, userUUID), res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *Controller) CreateReport(ctx context.Context, ratingData *model.Rating) (*model.Rating, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ratings.CreateReport.ctrl")
	ctx = opentracing.ContextWithSpan(ctx, span)
	defer span.Finish()

	res, err := c.repo.CreateReport(ctx, ratingData)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set(ctx, time.Hour, fmt.Sprintf(cacheKey, res.ID), res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *Controller) UpdateReport(ctx context.Context, ratingID uint, newData *model.Rating) (*model.Rating, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ratings.UpdateReport.ctrl")
	ctx = opentracing.ContextWithSpan(ctx, span)
	defer span.Finish()

	res, err := c.repo.UpdateReport(ctx, ratingID, newData)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set(ctx, time.Hour, fmt.Sprintf(cacheKey, res.ID), res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *Controller) DeleteReport(ctx context.Context, ratingID uint) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "ratings.DeleteReport.ctrl")
	ctx = opentracing.ContextWithSpan(ctx, span)
	defer span.Finish()

	if err := c.repo.DeleteReport(ctx, ratingID); err != nil {
		return err
	}

	if err := c.cache.Delete(ctx, fmt.Sprintf(cacheKey, ratingID)); err != nil {
		return err
	}
	return nil
}
