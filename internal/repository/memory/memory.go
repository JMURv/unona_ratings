package memory

import (
	"github.com/JMURv/unona/ratings/pkg/model"
	"github.com/google/uuid"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[uuid.UUID]*model.Rating
}

func New() *Repository {
	return &Repository{data: map[uuid.UUID]*model.Rating{}}
}

//TODO: Сделать
