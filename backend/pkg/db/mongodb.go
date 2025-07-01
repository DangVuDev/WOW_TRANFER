package db

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"wowtoken-api/internal/config"
	"wowtoken-api/internal/models"
)

var (
	client  *mongo.Client
	once    sync.Once
	errInit error
)

// ConnectMongoDB khởi tạo kết nối MongoDB (singleton)
func ConnectMongoDB() (*mongo.Client, error) {
	once.Do(func() {
		cfg, err := config.LoadConfig()
		if err != nil {
			errInit = fmt.Errorf("không thể tải cấu hình: %v", err)
			return
		}
		client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoDBURI))
		if err != nil {
			errInit = fmt.Errorf("không thể kết nối MongoDB: %v", err)
			return
		}
		// Kiểm tra kết nối
		if err := client.Ping(context.Background(), nil); err != nil {
			errInit = fmt.Errorf("không thể ping MongoDB: %v", err)
			client = nil
		}
	})
	if errInit != nil {
		return nil, errInit
	}
	return client, nil
}

// SaveUser lưu thông tin người dùng vào MongoDB
func SaveUser(user models.User) error {
	if client == nil {
		if _, err := ConnectMongoDB(); err != nil {
			return err
		}
	}
	collection := client.Database("wowtoken").Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("lỗi lưu người dùng: %v", err)
	}
	return nil
}

// GetUserByEmail lấy thông tin người dùng theo email
func GetUserByEmail(email string) (models.User, error) {
	if client == nil {
		if _, err := ConnectMongoDB(); err != nil {
			return models.User{}, err
		}
	}
	collection := client.Database("wowtoken").Collection("users")
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.User{}, fmt.Errorf("không tìm thấy người dùng với email: %s", email)
	}
	if err != nil {
		return models.User{}, fmt.Errorf("lỗi truy vấn người dùng: %v", err)
	}
	return user, nil
}

// UserExists kiểm tra xem email đã tồn tại chưa
func UserExists(email string) bool {
	if client == nil {
		if _, err := ConnectMongoDB(); err != nil {
			return false
		}
	}
	collection := client.Database("wowtoken").Collection("users")
	count, err := collection.CountDocuments(context.Background(), bson.M{"email": email})
	if err != nil {
		return false
	}
	return count > 0
}

// SaveEvent lưu sự kiện vào MongoDB
func SaveEvent(event models.Event) error {
	if client == nil {
		if _, err := ConnectMongoDB(); err != nil {
			return err
		}
	}
	collection := client.Database("wowtoken").Collection("events")
	_, err := collection.InsertOne(context.Background(), event)
	return err
}

// GetEventByID lấy sự kiện theo ID
func GetEventByID(eventID string) (models.Event, error) {
	if client == nil {
		if _, err := ConnectMongoDB(); err != nil {
			return models.Event{}, err
		}
	}
	collection := client.Database("wowtoken").Collection("events")
	var event models.Event
	err := collection.FindOne(context.Background(), bson.M{"id_event": eventID}).Decode(&event)
	if err == mongo.ErrNoDocuments {
		return models.Event{}, fmt.Errorf("không tìm thấy sự kiện với ID: %s", eventID)
	}
	return event, err
}

// UpdateEvent cập nhật sự kiện trong MongoDB
func UpdateEvent(event models.Event) error {
	if client == nil {
		if _, err := ConnectMongoDB(); err != nil {
			return err
		}
	}
	collection := client.Database("wowtoken").Collection("events")
	_, err := collection.ReplaceOne(context.Background(), bson.M{"id_event": event.IDEvent}, event)
	return err
}

// GetEncryptedPrivateKey lấy khóa riêng mã hóa theo địa chỉ
func GetEncryptedPrivateKey(address string) (string, error) {
	if client == nil {
		if _, err := ConnectMongoDB(); err != nil {
			return "", err
		}
	}
	collection := client.Database("wowtoken").Collection("users")
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"address": address}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return "", fmt.Errorf("không tìm thấy người dùng với địa chỉ: %s", address)
	}
	if err != nil {
		return "", err
	}
	return user.EncryptedPrivateKey, nil
}