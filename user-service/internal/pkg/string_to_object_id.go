package pkg

import "go.mongodb.org/mongo-driver/bson/primitive"

func StringToObjID(str string) (*primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return nil, err
	}
	return &objID, nil
}