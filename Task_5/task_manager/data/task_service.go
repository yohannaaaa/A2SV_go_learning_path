package data
import(
	"context"
	"errors"
	"log"
	"time"

	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed", err)
	}
	collection = client.Database("taskdb").Collection("tasks")

}

func GetAllTasks() ([]models.Task, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx , bson.M{})
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil

}
func GetTaskByID (id string) (*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}

	var task models.Task
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	return &task, err
}

func CreateTask(task *models.Task) error {	
	task.ID = primitive.NewObjectID().Hex()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, task)
	return err
}

func DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID format")
	}
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0{
		return errors.New("task not found")
	}
	return nil
}

func UpdateTask(id string, task *models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID format")
	}

	update := bson.M{"$set": task}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0{
		return errors.New("task not found")
	}
	return nil
}