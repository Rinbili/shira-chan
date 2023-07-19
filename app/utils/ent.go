package utils

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
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
	dsn := os.Getenv("sql_user") + ":" +
		os.Getenv("sql_passwd") + "@tcp(" +
		os.Getenv("sql_host") + ":" +
		os.Getenv("sql_port") + ")/" +
		os.Getenv("sql_name") + "?charset=utf8mb4&parseTime=True&loc=Local"
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
