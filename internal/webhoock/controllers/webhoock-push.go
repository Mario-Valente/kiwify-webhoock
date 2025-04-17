package webhook

import (
	"context"
	"fmt"

	"github.com/Mario-Valente/kiwify-webhoock/internal/config"
	"github.com/Mario-Valente/kiwify-webhoock/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	message := fmt.Sprintf(
		"⚠ Venda %s recebida! \n\n📦 Status:  %s \n👤 Cliente:  %s \n💳 Método de pagamento: %s \n📧 email: %s \n📱 telefone: %s",
		body.OrderStatus, body.OrderStatus, body.Customer.FullName, body.PaymentMethod, body.Customer.Email, body.Customer.Mobile,
	)

	go func() {
		err := SendTelegramMessage(ctx, message)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	collection := client.Database("kiwify").Collection("purchases")
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

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	collection := client.Database("kiwify").Collection("purchases")

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

func PostAbandoned(ctx context.Context, body *models.Abandoned) (models.Abandoned, error) {
	if body == nil {
		return models.Abandoned{}, fmt.Errorf("body is nil")
	}

	fmt.Println("Received abandoned webhook:", body.ID)

	client, err := config.GetClientMongoDB()
	if err != nil {
		return models.Abandoned{}, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	message := fmt.Sprintf(
		"🚨 Carrinhooooo abandonadOOOOOO olhar com atenção! \n👤 Cliente:  %s \n🌍 Pais: %s \n📧 email: %s \n📱 telefone: %s",
		body.Name, body.Country, body.Email, body.Phone,
	)

	go func() {
		err := SendTelegramMessage(ctx, message)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}

	}()

	collection := client.Database("kiwify").Collection("abandoned")

	// to do: remove the index creation from here

	emailModel := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(ctx, emailModel)
	if err != nil {
		return models.Abandoned{}, fmt.Errorf("failed to create index: %v", err)
	}
	_, err = collection.InsertOne(ctx, body)
	if err != nil {
		return models.Abandoned{}, fmt.Errorf("failed to insert document: %v", err)
	}
	fmt.Println("Document inserted successfully")

	return *body, nil
}

func GetAllAbandoned(ctx context.Context) ([]models.Abandoned, error) {
	client, err := config.GetClientMongoDB()
	if err != nil {
		return []models.Abandoned{}, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	collection := client.Database("kiwify").Collection("abandoned")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []models.Abandoned{}, fmt.Errorf("failed to find documents: %v", err)
	}
	defer cursor.Close(ctx)

	var results []models.Abandoned

	for cursor.Next(ctx) {
		var abandoned models.Abandoned
		if err := cursor.Decode(&abandoned); err != nil {
			return []models.Abandoned{}, fmt.Errorf("failed to decode document: %v", err)
		}
		results = append(results, abandoned)
	}
	if err := cursor.Err(); err != nil {
		return []models.Abandoned{}, fmt.Errorf("cursor error: %v", err)
	}

	if len(results) == 0 {
		return []models.Abandoned{}, fmt.Errorf("no documents found")
	}

	return results, nil
}
