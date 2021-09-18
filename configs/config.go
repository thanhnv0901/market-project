package configs

import (
	"fmt"
	"log"
	"market_apis/internals/utils"
	"os"
	"path"

	"github.com/tkanos/gonfig"
)

// configuration ..
type configuration struct {
	AppName    string
	AppVersion string
	Enironment string `env:"ENV"`
	APIHost    string
	APIPort    string
}

var config = configuration{}

func init() {
	log.Println("Reading config information service file")

	var (
		pathDir         string
		fileName        string
		patternFileName = "config.%s.json"
	)

	pathDir, _, _ = utils.GetDirCurrent()

	// Get config file name
	env := os.Getenv("GO_ENV")
	switch env {
	case "production", "staging", "test", "local":
		fileName = fmt.Sprintf(patternFileName, env)
	default:
		fileName = fmt.Sprintf(patternFileName, "local")
	}

	configFile := path.Join(pathDir, fileName)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		msg := fmt.Sprintf("%s not found", configFile)
		log.Fatalln(msg)
	}

	gonfig.GetConf(configFile, &config)

	log.Println("Read configuration file")
}

func GetConfig() *configuration {
	return &config
}
