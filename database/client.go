package database

import (
	"context"
	"time"

	"main/constant"
	"main/ent"
	"main/ent/migrate"

	"github.com/spf13/viper"
)

var Client *DbClient

type DbClient struct {
	WriteClient *ent.Client
	ReadClient  *ent.Client
}

var DefaultContextTime = viper.GetDuration(constant.CONTEXT_TIMEOUT_DATABASE) * time.Second

func Setup() (*DbClient, error) {
	migrateClient, err := ent.Open("mysql", viper.GetString(constant.DATABASE_MIGRATE_CONNECTION))
	if err != nil {
		return nil, err
	}

	err = setClientConfig(migrateClient)
	if err != nil {
		return nil, err
	}

	setTime := viper.GetInt(constant.CONTEXT_TIMEOUT_DATABASE)
	if setTime != 0 {
		DefaultContextTime = time.Duration(setTime * int(time.Second))
	}

	return &DbClient{
		ReadClient:  setClientMode(migrateClient),
		WriteClient: setClientMode(migrateClient),
	}, nil
}

func setClientConfig(client *ent.Client) error {
	return client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true))

}

func setClientMode(client *ent.Client) *ent.Client {
	serverRunMode := viper.GetString(constant.SERVER_RUN_MODE)
	if serverRunMode == "debug" {
		return client.Debug()
	}

	return client
}
