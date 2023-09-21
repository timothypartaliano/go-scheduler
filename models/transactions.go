package models

type Transaction struct {
	ID          int    `bson:"_id"`
	Description string `bson:"description"`
	Amount      int    `bson:"amount"`
}