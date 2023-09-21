package controllers

import (
	"context"
	"net/http"
	"scheduler/config"
	"scheduler/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTransaction(c echo.Context) error {
	transaction := new(models.Transaction)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(400, err)
	}

	result, err := config.Collection.InsertOne(context.Background(), transaction)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(http.StatusCreated, result)

}

func GetAllTransaction(c echo.Context) error {
	result, err := config.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer result.Close(context.Background())

	var transactions []models.Transaction
	for result.Next(context.Background()) {
		var t models.Transaction
		if err := result.Decode(&t); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		transactions = append(transactions, t)
	}

	return c.JSON(http.StatusOK, transactions)
}

func GetTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	filter := bson.M{"_id": id}
	var t models.Transaction
	err = config.Collection.FindOne(context.Background(), filter).Decode(&t)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, t)

}

func UpdateTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	transaction := new(models.Transaction)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"description": transaction.Description,
			"amount":      transaction.Amount,
		},
	}

	_, err = config.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Transaction updated")
}

func DeleteTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	filter := bson.M{"_id": id}
	_, err = config.Collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Transaction deleted")

}