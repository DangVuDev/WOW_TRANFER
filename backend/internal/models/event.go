package models

type Event struct {
    IDEvent       string   `json:"id_event" bson:"id_event"`
    IsCompleted   bool     `json:"isCompleted" bson:"isCompleted"`
    EventName     string   `json:"event_name" bson:"event_name"`
    Description   string   `json:"description" bson:"description"`
    SignatureCount int      `json:"signature_count" bson:"signature_count"`
    Signers       []string `json:"signers" bson:"signers"`
    CreatedAt     int64    `json:"createdAt" bson:"createdAt"`
}