package mongorm

import "go.mongodb.org/mongo-driver/bson/primitive"

func ToObjectID(id string) primitive.ObjectID {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return [12]byte{}
	}
	return objectID
}

func ObjToString(id primitive.ObjectID) string {
	return id.Hex()
}
