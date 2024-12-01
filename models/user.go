package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID
	Name  string             
	Email string             
	Age   int                
}