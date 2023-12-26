package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	pflag.String("log_level", log.DebugLevel.String(), "Logging level")

	pflag.Duration("refresh_timeout", 5*time.Second, "Refresh timeout")

	pflag.String("mongo_uri", "mongo://localhost:27017", "MongoDB url")
	pflag.String("postgres_uri", "postgres://root:1234@localhost:5432/root?sslmode=disable", "Postgres connection string")

	pflag.String("openai_api_key", "secret_key", "OpenAI API key")

	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()
}

func main() {

}
