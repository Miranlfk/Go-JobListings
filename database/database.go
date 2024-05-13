package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Miranlfk/go-graphql/graph/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type InstanceMongo struct {
	Client *mongo.Client
}

var Instance InstanceMongo

// ConnectDB connects to the MongoDB database using parameters from the .env file
func ConnectDB() *InstanceMongo {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_DB_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	
	return &InstanceMongo{
		Client: client,
	}

	
}

//TODO: Implement the following methods
func (db *InstanceMongo) GetJob(id string) *model.JobListing {
	return nil
}

func (db *InstanceMongo) GetJobs() []*model.JobListing {
	return nil
}

func (db *InstanceMongo) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	return nil
}

func (db *InstanceMongo) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	return nil
}

func (db *InstanceMongo) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	return nil
}