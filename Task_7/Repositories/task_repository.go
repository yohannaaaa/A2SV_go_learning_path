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

type TaskRepositoryMongo struct {
	collection *mongo.Collection
}

func NewTaskRepositoryMongo(dbURI string) *TaskRepositoryMongo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		panic(err)
	}

	return &TaskRepositoryMongo{
		collection: client.Database("taskdb").Collection("tasks"),
	}
}

func (r *TaskRepositoryMongo) GetAll() ([]domain.Task, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var tasks []domain.Task
	if err := cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepositoryMongo) GetByID(id string) (*domain.Task, error) {
	var task domain.Task
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	return &task, err
}

func (r *TaskRepositoryMongo) Create(task *domain.Task) error {
	_, err := r.collection.InsertOne(context.Background(), task)
	return err
}

func (r *TaskRepositoryMongo) Update(id string, task *domain.Task) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": task})
	return err
}

func (r *TaskRepositoryMongo) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
