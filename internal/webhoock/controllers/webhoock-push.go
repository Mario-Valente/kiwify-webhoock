package webhook

import (
	"context"
	"fmt"

	"github.com/Mario-Valente/kiwify-webhoock/internal/config"
	"github.com/Mario-Valente/kiwify-webhoock/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Post(ctx context.Context, body *models.Purchase) (models.Purchase, error) {
	if body == nil {
		return models.Purchase{}, fmt.Errorf("body is nil")
	}

	fmt.Println("Received webhook:", body.OrderID)

	client, err := config.GetClientMongoDB()
	if err != nil {
		return models.Purchase{}, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	defer client.Disconnect(ctx)

	collection := client.Database("kiwify").Collection("webhook")
	_, err = collection.InsertOne(ctx, body)
	if err != nil {
		return models.Purchase{}, fmt.Errorf("failed to insert document: %v", err)
	}
	fmt.Println("Document inserted successfully")

	return *body, nil
}

func Get(ctx context.Context, order_status string) (models.Purchase, error) {
	if order_status == "" {
		return models.Purchase{}, fmt.Errorf("customer name is empty")
	}

	client, err := config.GetClientMongoDB()
	if err != nil {
		return models.Purchase{}, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	defer client.Disconnect(ctx)

	collection := client.Database("kiwify").Collection("webhook")
	var result models.Purchase

	err = collection.FindOne(ctx, bson.M{
		"orderstatus": order_status,
	}).Decode(&result)

	if err != nil {
		return models.Purchase{}, fmt.Errorf("failed to find document: %v", err)
	}

	return result, nil
}
