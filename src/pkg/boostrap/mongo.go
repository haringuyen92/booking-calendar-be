package boostrap

import (
	"context"
	"fmt"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"time"
)

type MongoDBProperties struct {
	Uri      string `yaml:"uri"`
	Database string `yaml:"database"`
}

func (m MongoDBProperties) Prefix() string {
	return "app.mongodb"
}

func NewMongoDBProperties(
	loader config.Loader,
) (*MongoDBProperties, error) {
	props := MongoDBProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func NewMongoDB(dbConfig *MongoDBProperties) (*mongo.Client, *mongo.Database, error) {
	if dbConfig.Uri == "" {
		return nil, nil, fmt.Errorf("missing mongodb uri")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(dbConfig.Uri).
		SetMaxPoolSize(300).
		SetRetryWrites(true)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("MongoDB connection error: ", err)
		return nil, nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("MongoDB ping error: ", err)
		return nil, nil, err
	}

	database := client.Database(dbConfig.Database)
	return client, database, nil
}

func WithMongoDB() fx.Option {
	return fx.Options(
		golib.ProvideProps(NewMongoDBProperties),
		fx.Provide(NewMongoDB),
	)
}
