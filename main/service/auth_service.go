package service

import (
	"context"
	"errors"
	"hibernum-backend/main/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func ValidateCredentials(username string, password string) error {
	client, err := config.NewMongoClient().Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	usersCollection := client.Database("hibernum").Collection("users")

	var user bson.M
	err = usersCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return ErrUserNotFound
	}
	if err != nil {
		return err
	}

	storedPassword, ok := user["password"].(string)
	if !ok || storedPassword != password {
		return ErrInvalidCredentials
	}

	return nil
}
