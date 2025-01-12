package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Address struct
type address struct {
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Address string `json:"address" bson:"address"`
	PinCode string `json:"pincode" bson:"pincode"`
}

// User struct
type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Address   address            `json:"address" bson:"address"`
	JwtToken  string             `json:"jwtToken" bson:"jwtToken"`
}

// Sample JSON payload for User
// var sampleUser = User{
// 	Id:        primitive.NewObjectID(),
// 	FirstName: "John",
// 	LastName:  "Doe",
// 	Email:     "john.doe@example.com",
// 	Password:  "securepassword",
// 	Address: address{
// 		City:    "New York",
// 		State:   "NY",
// 		Address: "123 Main St",
// 		PinCode: "10001",
// 	},
// 	JwtToken: "samplejwttoken",
// }
