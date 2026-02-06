package repository

import (
	"context"
	"errors"
	"portfolio-backend/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PortfolioRepository interface {
	// Technology CRUD
	CreateTechnology(ctx context.Context, tech models.Technology) error
	GetTechnologies(ctx context.Context) ([]models.Technology, error)
	UpdateTechnology(ctx context.Context, id string, tech models.Technology) error
	DeleteTechnology(ctx context.Context, id string) error

	// Experience CRUD
	CreateExperience(ctx context.Context, exp models.Experience) error
	GetExperiences(ctx context.Context) ([]models.Experience, error)
	UpdateExperience(ctx context.Context, id string, exp models.Experience) error
	DeleteExperience(ctx context.Context, id string) error

	// Contact
	SaveContactMessage(ctx context.Context, msg models.ContactMessage) error
}

type mongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) PortfolioRepository {
	return &mongoRepository{db: db}
}

// --- Technology Implementation ---

func (r *mongoRepository) CreateTechnology(ctx context.Context, tech models.Technology) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := r.db.Collection("technologies")
	_, err := collection.InsertOne(ctx, tech)
	return err
}

func (r *mongoRepository) GetTechnologies(ctx context.Context) ([]models.Technology, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := r.db.Collection("technologies")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var techs []models.Technology
	if err = cursor.All(ctx, &techs); err != nil {
		return nil, err
	}
	return techs, nil
}

func (r *mongoRepository) UpdateTechnology(ctx context.Context, id string, tech models.Technology) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := r.db.Collection("technologies")
	update := bson.M{
		"$set": bson.M{
			"name":  tech.Name,
			"level": tech.Level,
			"icon":  tech.Icon,
		},
	}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("technology not found")
	}
	return nil
}

func (r *mongoRepository) DeleteTechnology(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := r.db.Collection("technologies")
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("technology not found")
	}
	return nil
}

// --- Experience Implementation ---

func (r *mongoRepository) CreateExperience(ctx context.Context, exp models.Experience) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := r.db.Collection("experiences")
	_, err := collection.InsertOne(ctx, exp)
	return err
}

func (r *mongoRepository) GetExperiences(ctx context.Context) ([]models.Experience, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := r.db.Collection("experiences")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var exps []models.Experience
	if err = cursor.All(ctx, &exps); err != nil {
		return nil, err
	}
	return exps, nil
}

func (r *mongoRepository) UpdateExperience(ctx context.Context, id string, exp models.Experience) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := r.db.Collection("experiences")
	update := bson.M{
		"$set": bson.M{
			"company":     exp.Company,
			"role":        exp.Role,
			"period":      exp.Period,
			"description": exp.Description,
			"projects":    exp.Projects,
		},
	}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("experience not found")
	}
	return nil
}

func (r *mongoRepository) DeleteExperience(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := r.db.Collection("experiences")
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("experience not found")
	}
	return nil
}

// --- Contact Implementation ---

func (r *mongoRepository) SaveContactMessage(ctx context.Context, msg models.ContactMessage) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := r.db.Collection("messages")
	_, err := collection.InsertOne(ctx, msg)
	return err
}
