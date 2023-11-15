package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/lujingwei002/gensc"
)

const (
	MONGODB_NAME = "mongo"
	REDIS_NAME   = "redis"
	MYSQL_NAME   = "mysql"
)

type Config struct {
	Driver          string        `yaml:"driver"`
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Db              string        `yaml:"db"`
	Query           string        `yaml:"query"`
	MaxOpenConns    int           `yaml:"max-open-conns"`
	MaxIdleConns    int           `yaml:"max-idle-conns"`
	ConnMaxIdleTime time.Duration `yaml:"conn-max-idle-time"`
	ConnMaxLifetime time.Duration `yaml:"conn-max-lifetime"`
	ConnnectTimeout time.Duration `yaml:"connect-timeout"`
}

// 完整的地址，包括path部分
func (self Config) Uri() string {
	if self.User != "" && self.Password != "" {
		switch self.Driver {
		case MONGODB_NAME:
			return fmt.Sprintf("mongodb://%s:%s@%s:%d/?db=%s&%s", self.User, self.Password, self.Host, self.Port, self.Db, self.Query)
		case REDIS_NAME:
			return fmt.Sprintf("redis://%s:%s@%s:%d?%s", self.User, self.Password, self.Host, self.Port, self.Query)
		case MYSQL_NAME:
			return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&%s", self.User, self.Password, self.Host, self.Port, self.Db, self.Query)
		default:
			return fmt.Sprintf("%s not support", self.Driver)
		}
	} else {
		switch self.Driver {
		case MONGODB_NAME:
			return fmt.Sprintf("mongodb://%s:%d/?db=%s&%s", self.Host, self.Port, self.Db, self.Query)
		case REDIS_NAME:
			return fmt.Sprintf("redis://%s:%d?%s", self.Host, self.Port, self.Query)
		case MYSQL_NAME:
			return fmt.Sprintf("@tcp(%s:%d)/%s?parseTime=true&%s", self.Host, self.Port, self.Db, self.Query)
		default:
			return fmt.Sprintf("%s not support", self.Driver)
		}
	}
	//return fmt.Sprintf("%s://%s:%s@%s:%d/%s", self.Driver, self.User, self.Password, self.Host, self.Port, self.Db)
}

func (self Config) GormUri() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", self.User, self.Password, self.Host, self.Port, self.Db)
}

func (self *Config) Parse(uri string) error {
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}
	switch u.Scheme {
	case MONGODB_NAME:
		self.Driver = u.Scheme
	case REDIS_NAME:
		self.Driver = u.Scheme
	case MYSQL_NAME:
		self.Driver = u.Scheme
	default:
		return gensc.ErrDbNotSupport
	}
	host2 := strings.Split(u.Host, ":")
	if len(host2) == 2 {
		self.Host = host2[0]
		if v, err := strconv.Atoi(host2[1]); err != nil {
			return err
		} else {
			self.Port = v
		}
	} else if len(host2) == 1 {
		switch u.Scheme {
		case MONGODB_NAME:
			self.Port = 27017
		case REDIS_NAME:
			self.Port = 6379
		case MYSQL_NAME:
			self.Port = 3306
		}
	}
	self.User = u.User.Username()
	if v, set := u.User.Password(); set {
		self.Password = v
	} else {
		self.Password = ""
	}
	switch u.Scheme {
	case MONGODB_NAME:
		self.Db = u.Query().Get("db")
		query := u.Query()
		query.Del("db")
		self.Query = query.Encode()
	default:
		path := strings.TrimPrefix(u.Path, "/")
		self.Db = path
		self.Query = u.Query().Encode()
	}
	return nil
}

// 根据配置构造mongodb client
func NewConfigMongoDbClient(ctx context.Context, name string, config Config) (gensc.DbClient, error) {
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	uri := config.Uri()
	client := &MongoDbClient{
		config:     config,
		cancelFunc: cancelFunc,
		ctx:        cancelCtx,
		uri:        uri,
	}
	clientOpts := options.Client().ApplyURI(uri)
	if config.ConnnectTimeout > 0 {
		// the default is 30 seconds.
		clientOpts.SetConnectTimeout(config.ConnnectTimeout)
	}
	conn, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Println("connect database fail", "name", name, "uri", uri, "error", err)
		return nil, err
	}
	// ctx2, cancelFunc2 := context.WithTimeout(client.ctx, config.ConnnectTimeout*time.Second)
	// defer cancelFunc2()
	if err = conn.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("ping database fail", "name", name, "uri", uri, "error", err)
		return nil, err
	}
	client.client = conn
	log.Println("connect database success", "name", name, "uri", uri)
	return client, nil
}

// 根据配置构造redis client
func NewConfigRedisClient(ctx context.Context, name string, config Config) (gensc.DbClient, error) {
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	uri := config.Uri()
	client := &RedisClient{
		config:     config,
		cancelFunc: cancelFunc,
		ctx:        cancelCtx,
		uri:        uri,
	}
	var err error
	var db int
	db, err = strconv.Atoi(config.Db)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       db,
	})
	ctx1, cancelFunc1 := context.WithTimeout(client.ctx, 3*time.Second)
	defer cancelFunc1()
	if _, err := rdb.Ping(ctx1).Result(); err != nil {
		log.Println("connect database fail", "name", name, "uri", uri, "error", err)
		return nil, err
	}
	client.client = rdb
	log.Println("connect database success", "name", name, "uri", uri)
	return client, nil
}

// 根据配置构造mysql client
func NewConfigMysqlClient(ctx context.Context, name string, config Config) (gensc.DbClient, error) {
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	uri := config.Uri()
	client := &MysqlClient{
		config:     config,
		cancelFunc: cancelFunc,
		ctx:        cancelCtx,
		uri:        uri,
	}
	db, err := sql.Open("mysql", uri)
	if err != nil {
		log.Println("connect database fail", "name", name, "uri", uri, "error", err)
		return nil, err
	}
	if config.MaxOpenConns > 0 {
		db.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.ConnMaxIdleTime > 0 {
		db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	}
	if config.MaxIdleConns > 0 {
		db.SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(config.ConnMaxLifetime)
	}
	err = db.Ping()
	if err != nil {
		log.Println("connect database fail", "name", name, "uri", uri, "error", err)
		return nil, err
	}
	client.client = db
	log.Println("connect database success", "name", name, "uri", uri)
	return client, nil
}

// 根据url构造db client
func NewDbClientFromUri(ctx context.Context, name string, uri string) (gensc.DbClient, error) {
	config := Config{}
	if err := config.Parse(uri); err != nil {
		return nil, err
	}
	fmt.Println(uri)
	fmt.Println(config.Db)
	return NewConfigDbClient(ctx, name, config)
}

// 根据配置构造db client
func NewConfigDbClient(ctx context.Context, name string, config Config) (gensc.DbClient, error) {
	switch config.Driver {
	case MONGODB_NAME:
		return NewConfigMongoDbClient(ctx, name, config)
	case REDIS_NAME:
		return NewConfigRedisClient(ctx, name, config)
	case MYSQL_NAME:
		return NewConfigMysqlClient(ctx, name, config)
	default:
		return nil, gensc.ErrDbNotSupport
	}
}

// mongodb客户端
type MongoDbClient struct {
	cancelFunc context.CancelFunc
	ctx        context.Context
	client     *mongo.Client
	config     Config
	uri        string
}

func (self *MongoDbClient) GetMongoDatabase() *mongo.Database {
	return self.client.Database(self.config.Db)
}

func (self *MongoDbClient) GetMongoClient() *mongo.Client {
	return self.client
}

// mysql客户端
type MysqlClient struct {
	cancelFunc context.CancelFunc
	ctx        context.Context
	client     *sql.DB
	config     Config
	uri        string
}

func (self *MysqlClient) Uri() string {
	return self.uri
}

func (self *MysqlClient) GetMysqlClient() *sql.DB {
	return self.client
}

// redis客户端
type RedisClient struct {
	cancelFunc context.CancelFunc
	ctx        context.Context
	client     *redis.Client
	config     Config
	uri        string
}

func (self *RedisClient) Uri() string {
	return self.uri
}

func (self *RedisClient) GetRedisClient() *redis.Client {
	return self.client
}

func (self *MongoDbClient) Uri() string {
	return self.uri
}
