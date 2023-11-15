package gensc

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbClient interface {
	Uri() string
}

type RedisClient interface {
	DbClient
	GetRedisClient() *redis.Client
}

type MysqlClient interface {
	DbClient
	GetMysqlClient() *sql.DB
}

type MongoClient interface {
	DbClient
	GetMongoClient() *mongo.Client
	GetMongoDatabase() *mongo.Database
}

type DbDao interface {
	UseClient(client DbClient) error
}
