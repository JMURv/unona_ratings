package db

import (
	"context"
	"fmt"
	repo "github.com/JMURv/unona/ratings/internal/repository"
	conf "github.com/JMURv/unona/ratings/pkg/config"
	"github.com/JMURv/unona/ratings/pkg/model"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type Repository struct {
	conn *gorm.DB
}

func New(conf *conf.DBConfig) *Repository {
	DSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)

	conn, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = conn.AutoMigrate(
		&model.Rating{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{conn: conn}
}

func (r *Repository) GetRatingByID(ctx context.Context, ratingID uint) (*model.Rating, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ratings.GetRatingByID.repo")
	defer span.Finish()

	var rating model.Rating
	if err := r.conn.Where("ID=?", ratingID).First(&rating).Error; err != nil {
		return nil, err
	}
	return &rating, nil
}

func (r *Repository) GetUserRating(ctx context.Context, userUUID uuid.UUID) (*model.Rating, error) {
	//TODO: Привести к 5.0
	span, _ := opentracing.StartSpanFromContext(ctx, "ratings.GetUserRating.repo")
	defer span.Finish()

	var ratings []*model.Rating
	if err := r.conn.Where("UserUUID=?", userUUID).Find(&ratings).Error; err != nil {
		return nil, err
	}

	var sum uint8
	var digits []uint8
	for _, rating := range ratings {
		sum += rating.Rating
		digits = append(digits, rating.Rating)
	}

	return nil, nil
}

func (r *Repository) CreateReport(ctx context.Context, rating *model.Rating) (*model.Rating, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ratings.CreateReport.repo")
	defer span.Finish()

	if rating.UserUUID == uuid.Nil {
		return nil, repo.ErrUserUUIDIsRequired
	}

	if rating.Rating == 0 {
		return nil, repo.ErrRatingIsRequired
	}

	if rating.Text == "" {
		return nil, repo.ErrTextIsRequired
	}

	rating.CreatedAt = time.Now()
	rating.UpdatedAt = time.Now()
	if err := r.conn.Create(&rating).Error; err != nil {
		return nil, err
	}

	return rating, nil
}

func (r *Repository) UpdateReport(ctx context.Context, ratingID uint, newData *model.Rating) (*model.Rating, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ratings.UpdateReport.repo")
	defer span.Finish()

	rating, err := r.GetRatingByID(ctx, ratingID)
	if err != nil {
		return nil, err
	}

	if rating.Rating == 0 {
		rating.Rating = newData.Rating
	}

	if newData.Text != "" {
		rating.Text = newData.Text
	}

	rating.UpdatedAt = time.Now()
	r.conn.Save(&rating)
	return rating, nil
}

func (r *Repository) DeleteReport(ctx context.Context, ratingID uint) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "ratings.DeleteReport.repo")
	defer span.Finish()

	var rating model.Rating
	if err := r.conn.Delete(&rating, ratingID).Error; err != nil {
		return err
	}
	return nil
}
