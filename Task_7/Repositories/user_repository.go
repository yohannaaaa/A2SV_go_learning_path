package repositories

import (
	"context"
	"errors"
	"task-manager/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepositoryMongo(dbURI string) *UserRepositoryMongo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		panic(err)
	}

	return &UserRepositoryMongo{
		collection: client.Database("taskdb").Collection("users"),
	}
}

func (r *UserRepositoryMongo) Create(user *domain.User) error {
	_, err := r.collection.InsertOne(context.Background(), bson.M{
		"username": user.Username,
		"password": user.Password,
		"role":     user.Role,
	})
	return err
}

func (r *UserRepositoryMongo) FetchByUsername(username string) (*domain.User, error) {
	var result bson.M
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}

	return &domain.User{
		ID:       result["_id"].(string),
		Username: result["username"].(string),
		Password: result["password"].(string),
		Role:     domain.Role(result["role"].(string)),
	}, nil
}
