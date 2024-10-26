package domain

import "time"

type Status int
type Role int

const (
	ACTIVE 		Status = iota
	INACTIVE
	BANNED
)

const (
	UNKNOWN		Role = iota
	ADMIN
	CONSUMER
	MANAGER
	VENDOR
	ASSISTANT
)

type User struct {
	ID				string			`json:"id" bson:"_id,omitempty"`
	Email			string			`json:"email" bson:"email"`
	Password		string			`json:"password" bson:"password"`
	PhoneNumber		string			`json:"phone_number" bson:"phone_number"`
	FirstName		string			`json:"first_name" bson:"first_name"`
	MiddleName		string			`json:"middle_name" bson:"middle_name"`
	LastName		string			`json:"last_name" bson:"last_name"`
	Status			Status			`json:"status" bson:"status"`
	Role			Role			`json:"role" bson:"role"`
	Avatar			string			`json:"avatar" bson:"avatar"`
	Location		string			`json:"location" bson:"location"`
	CreatedAt		time.Time		`json:"created_at" bson:"created_at"`
}

