package controllers

import (
	"context"
	"go-mongodb-api-example/configs"
	"go-mongodb-api-example/models"
	"go-mongodb-api-example/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(g *gin.Context) {
		// Define timeout with 10 seconds
		ctx, exit := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User

		defer exit()

		// Validating user body
		if err := g.BindJSON(&user); err != nil {
			g.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Error encountered on user body",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Validate required fields with validator
		if err := validate.Struct(user); err != nil {
			g.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Error encountered",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		newUser := models.User{
			Name:     user.Name,
			Level:    user.Level,
			Username: user.Username,
			Password: user.Password,
		}

		res, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			g.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error while inserting",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		g.JSON(http.StatusCreated, responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "Success inserting data",
			Data: map[string]interface{}{
				"data": res,
			},
		})
	}
}
