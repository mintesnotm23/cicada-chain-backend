package models

import (
	"context"
	// "errors"
	"time"


	"go.mongodb.org/mongo-driver/mongo"


	"os"
	"fmt"
	
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type EmailVerification struct {
	ID               primitive.ObjectID `json:"_id"`
	Email            string             `json:"email" unique:"true"`
	VerificationCode string             `json:"verificationCode"`
	CreatedAt        time.Time          `json:"createdAt"`
	ExpiresAt        time.Time          `json:"expiresAt"`
}

var (
	mongoClient *mongo.Client
	database    *mongo.Database
	collection  *mongo.Collection
)

type DB struct {
	Client *mongo.Client
}

func ConnectToMongoDB() {
	connectionString := os.Getenv("MONGODB_URL")
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	mongoClient = client
	database = client.Database("email-verification")
	collection = database.Collection("emailVerifications")
	fmt.Println("Connected to MongoDB")
}

func GetClient() *mongo.Client {
	return mongoClient
}