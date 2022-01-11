// provides an interface for the accessing multiple databases
package dataaccess

import (
	"context"

	"github.com/mitchellh/mapstructure"
)

type StorageType int

const (
    mongoDB StorageType = 1 << iota
    redisDB
)

/*
The functions that are exposed to be used by multiple databases
*/
type DbConnector interface {
	Connect() error
	FindOne(context.Context, string, interface{}) (interface{}, error)
	FindMany(context.Context, string, interface{}) ([]interface{}, error)
	InsertOne(context.Context, string, interface{}) (interface{}, error)
	InsertMany(context.Context, string, []interface{}) ([]interface{}, error)
	UpdateOne(context.Context, string, interface{}, interface{}) (interface{}, error)
	UpdateMany(context.Context,string, interface{}, interface{}) (interface{}, error)
	Cancel() error
}

type SingleResultHelper interface {
	Decode(v interface{}) error
}


type DbConfig struct {
    DbType StorageType
    DbUrl string
    DbName string
}

func NewStore(config interface{}) DbConnector {
	configDoc := DbConfig{}
	mapstructure.Decode(config, &configDoc)

    switch configDoc.DbType {
    case mongoDB:
		return newMongoClient(configDoc.DbUrl, configDoc.DbName)
    case redisDB:
		return newRedisClient(configDoc.DbUrl)
	}
	return nil
}