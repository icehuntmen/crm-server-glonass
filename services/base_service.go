package services

import (
	"context"
	"crm-glonass/config"
	"crm-glonass/pkg/logging"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostServiceImpl struct {
	postCollection *mongo.Collection
	ctx            context.Context
}

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	baseCollection *mongo.Collection
	Logger         logging.Logger
	ctx            context.Context
}

func NewBaseService[T any, Tc any, Tu any, Tr any](baseCollection *mongo.Collection, cfg *config.Config, ctx context.Context) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		baseCollection: baseCollection,
		Logger:         logging.NewLogger(cfg),
		ctx:            ctx,
	}
}
