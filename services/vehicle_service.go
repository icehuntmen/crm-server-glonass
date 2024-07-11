package services

import (
	"context"
	"crm-glonass/config"
	"crm-glonass/data/models"
	"crm-glonass/pkg/logging"
	"crm-glonass/pkg/tools"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type VehicleService struct {
	Mongo      *mongo.Database
	Collection *mongo.Collection
	ctx        context.Context
	logger     logging.Logger
}

type VehicleServiceApi interface {
	// Create creates a new vehicle post using the provided request.
	//
	// Parameters:
	// - post: The request containing the data for the new post.
	//
	// Returns:
	// - *models.DBVehicle: The newly created post.
	// - error: An error if the creation operation fails.
	Create(*models.CreateVehicleRequest) (*models.DBVehicle, error)
	// Update updates a vehicle post with the given ID using the provided data.
	//
	// Parameters:
	// - id: The ID of the vehicle post to update.
	// - data: The updated data for the vehicle post.
	//
	// Returns:
	// - *models.DBVehicle: The updated vehicle post.
	// - error: An error if the update operation fails.
	Update(string, *models.UpdateVehicleRequest) (*models.DBVehicle, error)
	// FindById retrieves a DBVehicle object based on the provided ID.
	//
	// Parameters:
	// - id: The ID of the post to retrieve.
	// Return type:
	// - *models.DBVehicle: A pointer to the retrieved post.
	// - error: An error if the retrieval operation fails.
	FindById(string) (*models.DBVehicle, error)
	// Find retrieves a list of DBVehicle objects based on the provided page and limit.
	//
	// Parameters:
	// - page: The page number of the results to retrieve. Defaults to 1 if not provided.
	// - limit: The maximum number of results to retrieve per page. Defaults to 10 if not provided.
	//
	// Returns:
	// - []*models.DBVehicle: A slice of pointers to DBVehicle objects representing the retrieved posts.
	// - error: An error if the retrieval operation fails.
	Find(page int, limit int) ([]*models.DBVehicle, error)
	// Delete deletes a post with the given ID.
	//
	// id: The ID of the post to be deleted.
	// Returns an error if the deletion fails.
	Delete(string) error
}

func NewVehicleService(db *mongo.Database, cfg *config.Config, ctx context.Context, collectionName string) VehicleServiceApi {
	return &VehicleService{
		Mongo:      db,
		Collection: db.Collection(collectionName),
		ctx:        ctx,
		logger:     logging.NewLogger(cfg),
	}
}

// CreatePost creates a new vehicle post using the provided request.
//
// Parameters:
// - post: The request containing the data for the new post.
// Returns:
// - *models.DBVehicle: The newly created post.
// - error: An error if the creation operation fails.
func (p *VehicleService) Create(vehicle *models.CreateVehicleRequest) (*models.DBVehicle, error) {
	vehicle.CreatedAt = time.Now()
	vehicle.UpdatedAt = vehicle.CreatedAt
	fmt.Println(vehicle)
	res, err := p.Collection.InsertOne(p.ctx, vehicle)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("vehicle with that name already exists")
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}

	if _, err := p.Collection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New("could not create index for title")
	}

	var newPost *models.DBVehicle
	query := bson.M{"_id": res.InsertedID}
	if err = p.Collection.FindOne(p.ctx, query).Decode(&newPost); err != nil {
		return nil, err
	}

	return newPost, nil
}

// Update updates a vehicle post with the given ID using the provided data.
//
// Parameters:
// - id: The ID of the vehicle post to update.
// - data: The updated data for the vehicle post.
// Returns:
// - *models.DBVehicle: The updated vehicle post.
// - error: An error if the update operation fails.
func (p *VehicleService) Update(id string, data *models.UpdateVehicleRequest) (*models.DBVehicle, error) {
	doc, err := tools.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.Collection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedPost *models.DBVehicle
	if err := res.Decode(&updatedPost); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedPost, nil
}

// FindById retrieves a DBVehicle object based on the provided ID.
//
// Parameters:
// - id: The ID of the post to retrieve.
//
// Returns:
//   - *models.DBVehicle: A pointer to the retrieved post.
//   - error: An error if the retrieval operation fails. If the post with the given ID does not exist,
//     an error with the message "no document with that Id exists" is returned.
func (p *VehicleService) FindById(id string) (*models.DBVehicle, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var vehicle *models.DBVehicle

	if err := p.Collection.FindOne(p.ctx, query).Decode(&vehicle); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return vehicle, nil
}

// Find retrieves a list of DBVehicle objects based on the provided pagination parameters.
//
// Parameters:
// - page: The page number for pagination. If 0, defaults to 1.
// - limit: The maximum number of items to retrieve per page. If 0, defaults to 10.
// Returns:
// - []*models.DBVehicle: A slice of DBVehicle objects.
// - error: An error if the retrieval operation fails.
func (p *VehicleService) Find(page int, limit int) ([]*models.DBVehicle, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))
	opt.SetSort(bson.M{"created_at": -1})

	query := bson.M{}

	cursor, err := p.Collection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var vehicles []*models.DBVehicle

	for cursor.Next(p.ctx) {
		post := &models.DBVehicle{}
		err := cursor.Decode(post)

		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(vehicles) == 0 {
		return []*models.DBVehicle{}, nil
	}

	return vehicles, nil
}

// Delete deletes a vehicle post with the given ID.
//
// Parameters:
// - id: The ID of the vehicle post to delete.
//
// Returns:
//   - error: An error if the deletion operation fails. If the post with the given ID does not exist,
//     an error with the message "no document with that Id exists" is returned.
func (p *VehicleService) Delete(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.Collection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
