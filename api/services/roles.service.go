package services

import (
	"context"
	"crm-glonass/api/dto"
	"crm-glonass/config"
	"crm-glonass/pkg/logging"
	"crm-glonass/pkg/service_errors"
	"crm-glonass/pkg/tools"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type RoleService struct {
	Mongo      *mongo.Database
	Collection *mongo.Collection
	ctx        context.Context
	logger     logging.Logger
}

func NewRoleService(db *mongo.Database, cfg *config.Config, ctx context.Context, collectionName string) RoleInterface {
	return &RoleService{
		Mongo:      db,
		Collection: db.Collection(collectionName),
		ctx:        ctx,
		logger:     logging.NewLogger(cfg),
	}
}

func (r *RoleService) CreateRole(role *dto.Role) error {
	role.ID = tools.GenerateUUID()
	role.Name = strings.ToLower(role.Name)
	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}

	if _, err := r.Collection.Indexes().CreateOne(r.ctx, index); err != nil {
		r.logger.Error(logging.MongoDB, logging.CreateIndex, err.Error(), nil)
		return err
	}

	_, err := r.Collection.InsertOne(r.ctx, role)
	if err != nil {
		var er mongo.WriteException
		if errors.As(err, &er) && er.WriteErrors[0].Code == 11000 {
			r.logger.Error(logging.MongoDB, logging.Insert, err.Error(), nil)
			return &service_errors.ServiceError{EndUserMessage: service_errors.RoleExists}
		}
		return err
	}

	return err
}

func (r *RoleService) ListRoles() ([]dto.RoleList, error) {
	cursor, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		r.logger.Error(logging.MongoDB, logging.Find, err.Error(), nil)
	}
	var roles []dto.RoleList
	if err = cursor.All(context.TODO(), &roles); err != nil {
		r.logger.Error(logging.MongoDB, logging.Find, err.Error(), nil)
	}
	fmt.Println(roles)
	return roles, nil
}
