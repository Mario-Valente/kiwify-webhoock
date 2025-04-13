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

	go SendTelegramMessage(ctx, fmt.Sprintf("Received a new sale with the status: %s name: %s  com o metodo de pagamento: %s ", body.OrderStatus, body.Customer.FullName, body.PaymentMethod))
	if err != nil {
		return models.Purchase{}, fmt.Errorf("failed to send telegram message: %v", err)
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

func GetAllByStatus(ctx context.Context, order_status string) ([]models.Purchase, error) {
	if order_status == "" {
		return []models.Purchase{}, fmt.Errorf("order status is required")
	}

	client, err := config.GetClientMongoDB()
	if err != nil {
		return []models.Purchase{}, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	defer client.Disconnect(ctx)

	collection := client.Database("kiwify").Collection("webhook")

	filter := bson.M{"orderstatus": order_status}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return []models.Purchase{}, fmt.Errorf("failed to find documents: %v", err)
	}
	defer cursor.Close(ctx)

	var results []models.Purchase

	for cursor.Next(ctx) {
		var purchase models.Purchase
		if err := cursor.Decode(&purchase); err != nil {
			return []models.Purchase{}, fmt.Errorf("failed to decode document: %v", err)
		}
		results = append(results, purchase)
	}
	if err := cursor.Err(); err != nil {
		return []models.Purchase{}, fmt.Errorf("cursor error: %v", err)
	}
	if len(results) == 0 {
		return []models.Purchase{}, fmt.Errorf("no documents found with order status: %s", order_status)
	}

	return results, nil
}
