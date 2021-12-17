package mongorm

import "go.mongodb.org/mongo-driver/bson/primitive"

func ToObjectID(id string) primitive.ObjectID {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return [12]byte{}
	}
	return objectID
}

func ObjectID(id string) primitive.ObjectID {
	return ToObjectID(id)
}

func ObjToString(id primitive.ObjectID) string {
	return id.Hex()
}
