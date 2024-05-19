package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Miranlfk/go-graphql/graph/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetJob retrieves a job listing from the database
func (db *InstanceMongo) GetJob(id string) *model.JobListing {
	jobCollection := db.Client.Database("JobListings").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) 
	defer cancel()

	_id , _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	err := jobCollection.FindOne(ctx, filter).Decode(&jobListing) 
	if err != nil {
		log.Fatal(err)
	
	}
	return &jobListing
}

func (db *InstanceMongo) GetJobs() []*model.JobListing {
	jobCollection := db.Client.Database("JobListings").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) 
	defer cancel()

	var jobListings []*model.JobListing
	cursor, err := jobCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	
	}
	if err = cursor.All(context.TODO(), &jobListings); err != nil {
		panic (err)
	}
	return jobListings
}

func (db *InstanceMongo) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	jobCollection := db.Client.Database("JobListings").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) 
	defer cancel()

	insertJob, err := jobCollection.InsertOne(ctx, bson.M{
		"title": jobInfo.Title,
		"description": jobInfo.Description,
		"url": jobInfo.URL,
		"company": jobInfo.Company,
	})

	if err != nil {
		log.Fatal(err)
	
	}
	insertedID := insertJob.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing := model.JobListing{ID: insertedID, 
		Title: jobInfo.Title, 
		Description: jobInfo.Description, 
		URL: jobInfo.URL, 
		Company: jobInfo.Company,
	}
	
	return &returnJobListing
}

func (db *InstanceMongo) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	jobCollection := db.Client.Database("JobListings").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) 
	defer cancel()

	updateJobInfo := bson.M{}
	if jobInfo.Title != nil {
		updateJobInfo["title"] = *jobInfo.Title
	}

	if jobInfo.Description != nil {
		updateJobInfo["description"] = *jobInfo.Description
	}

	if jobInfo.URL != nil {
		updateJobInfo["url"] = *jobInfo.URL
	}

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing
	if err := results.Decode(&jobListing); err != nil {
		log.Fatal(err)
	
	}
	return &jobListing
	
}

func (db *InstanceMongo) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	jobCollection := db.Client.Database("JobListings").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) 
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	
	_, err := jobCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	
	return &model.DeleteJobResponse{DeleteJobID: jobId}
}