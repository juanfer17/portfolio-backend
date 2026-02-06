package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Technology struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name" binding:"required"`
	Level string             `bson:"level" json:"level" binding:"required"`
	Icon  string             `bson:"icon" json:"icon"`
}

type Project struct {
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	TechStack   []string `bson:"tech_stack" json:"tech_stack"`
}

type Experience struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Company     string             `bson:"company" json:"company" binding:"required"`
	Role        string             `bson:"role" json:"role" binding:"required"`
	Period      string             `bson:"period" json:"period" binding:"required"`
	Description string             `bson:"description" json:"description"`
	Projects    []Project          `bson:"projects" json:"projects"`
}

type ContactMessage struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name" binding:"required"`
	Email     string             `bson:"email" json:"email" binding:"required,email"`
	Message   string             `bson:"message" json:"message" binding:"required"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
