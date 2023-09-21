package services

import (
	"context"
	"fmt"
	"scheduler/config"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func StartScheduler() {
	c := cron.New()

	_, err := c.AddFunc("*/5 * * * *", func() {
		err := DeleteHistory()
		if err != nil {
			fmt.Println("Error deleting history:", err)
		} else {
			fmt.Println("History deleted successfully")
		}

	})
	if err != nil {
		fmt.Println("Error scheduling job:", err)
	}

	c.Start()

	//Select {}
}

func DeleteHistory() error {
	// Define an empty query to match all documents in the collection
	query := bson.M{}

	// Delete all data from MongoDB
	_, err := config.Collection.DeleteMany(context.Background(), query)
	return err
}