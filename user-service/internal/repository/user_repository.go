package repository

import (
	"context"
	"errors"
	"log"

	"github.com/pranay999000/smart-inventory/user-service/internal/domain"
	"github.com/pranay999000/smart-inventory/user-service/internal/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	MongoClient		*mongo.Client
	Users			*mongo.Collection	
}

func NewUserRepository(db *mongo.Client, users *mongo.Collection) *UserRepository {
	return &UserRepository{
		MongoClient: db,
		Users: users,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {

	emailCheck, err := r.CheckEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	phoneNumberCheck, err := r.CheckPhoneNumber(ctx, user.PhoneNumber)
	if err != nil {
		return err
	}

	if emailCheck || phoneNumberCheck {
		log.Printf("Email: %s or PhoneNumber: %s exists\n", user.Email, user.PhoneNumber)
		return errors.New("email or phone number already exists")
	}

	encryptedPassword, err := pkg.HashPassword(user.Password)
	if err != nil {
		return errors.New("password error")
	}

	user.Password = encryptedPassword

	insertedUser, err := r.Users.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Failed to insert user: %v\n", err)
		return err
	}

	log.Printf("Inserted user %v\n", insertedUser)
	return nil

}

func (r *UserRepository) CheckEmail(ctx context.Context, email string) (bool, error) {

	var result struct{}

	err := r.Users.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		log.Printf("Failed to find email %s: %v\n", email, err)
		return false, err
	}

	return true, nil

}

func (r *UserRepository) CheckPhoneNumber(ctx context.Context, phoneNumber string) (bool, error) {

	var result struct{}

	err := r.Users.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		log.Printf("Failed to find phone number %s: %v", phoneNumber, err)
		return false, err
	}

	return true, nil

}