// Fill Config struct from YAML and EnvVars

package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Slack struct {
		SigningSecret           string `yaml:"signingSecret"`
		BotUserOauthAccessToken string `yaml:"botUserOauthAccessToken"`
	} `yaml:"slack"`
	Opsgenie struct {
		APIKey string `yaml:"apiKey"`
	} `yaml:"opsgenie"`
	AppConfiguration struct {
		DebugLevel string `yaml:"debugLevel"`
	} `yaml:"appConfiguration"`
}

func main() {
	var cfg Config
	readFile(&cfg)
	readEnv(&cfg)
	fmt.Printf("%+v", cfg)

	fmt.Printf("\n%s", cfg.Opsgenie.APIKey)

}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yaml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
