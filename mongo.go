package mongorm

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Conf struct {
	URI string `json:"uri"`
}

func NewMongoClient(mongoConf string) (*mongo.Client, error) {
	var conf Conf
	err := json.Unmarshal([]byte(mongoConf), &conf)
	if err != nil {
		return nil, err
	}

	clientOpt := options.Client()
	clientOpt.ApplyURI(conf.URI)
	client, err := mongo.NewClient(clientOpt)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
