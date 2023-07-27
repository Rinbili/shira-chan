package utils

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"shira-chan-dev/config"
	"shira-chan-dev/ent"
	"shira-chan-dev/ent/migrate"
	_ "shira-chan-dev/ent/runtime"
)

// Client 全局数据库连接
var Client = InitEntClient()

// InitEntClient
// @Description: 初始化数据库连接
// @return *ent.Client
func InitEntClient() *ent.Client {
	// Create ent.Client and run the schema migration.
	dsn := config.Config.Sql.User + ":" +
		config.Config.Sql.Passwd + "@tcp(" +
		config.Config.Sql.Host + ":" +
		config.Config.Sql.Port + ")/" +
		config.Config.Sql.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}
	return client
}
