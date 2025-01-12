package routes

import (
	models "GinFrameWork/Models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var DataBaseName string = "Go_With"
var UserCollection string = "users"

type UserController struct {
	client *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}

// SetupRouter function
func (uc *UserController) BasicRoute(router *gin.Engine, ctx context.Context) {
	userRouter := router.Group("/users")
	userRouter.GET("/users", uc.GetUsers(ctx))
	userRouter.POST("/users", uc.CreateUser(ctx))
}

// GetUsers handler
func (uc *UserController) GetUsers(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: Fetch users from MongoDB and return as JSON
		collection := uc.client.Database("Go_With").Collection("users")
		cursor, err := collection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(ctx)

		var users []bson.M
		if err = cursor.All(ctx, &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func (uc *UserController) CreateUser(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: Create a new user in MongoDB}}

		collection := uc.client.Database(DataBaseName).Collection(UserCollection)

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.Id = primitive.NewObjectID()

		if result, err := collection.InsertOne(ctx, user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"insertedID": result.InsertedID, "message": "User created successfully"})
		}

	}

}
