package models

type User struct {
    UserID             string `bson:"userId" json:"userId"`
    Email              string `bson:"email" json:"email"`
    HashedPassword     string `bson:"hashedPassword" json:"-"`
    Name               string `bson:"name" json:"name"`
    Address            string `bson:"address" json:"address"`
    EncryptedPrivateKey string `bson:"encryptedPrivateKey" json:"-"`
    RegisteredAt       int64  `bson:"registeredAt" json:"registeredAt"`
}