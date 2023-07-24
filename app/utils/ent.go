package utils

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"shira-chan-dev/ent"
	"shira-chan-dev/ent/migrate"
	_ "shira-chan-dev/ent/runtime"
)

// Client 全局数据库连接
var Client = InitEntClient()
var conf = new(Yaml)

type Yaml struct {
	Secret string `yaml:"secret"`
	Sql    struct {
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		Name   string `yaml:"name"`
	}
}

// InitEntClient
// @Description: 初始化数据库连接
// @return *ent.Client
func InitEntClient() *ent.Client {
	// Create ent.Client and run the schema migration.
	yamlFile, err := os.ReadFile("./config/config.yml")
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic("can not read config.yml")
	}
	dsn := conf.Sql.User + ":" +
		conf.Sql.Passwd + "@tcp(" +
		conf.Sql.Host + ":" +
		conf.Sql.Port + ")/" +
		conf.Sql.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
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
