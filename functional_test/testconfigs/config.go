package testconfigs

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/tkanos/gonfig"
)

// configuration ..
type configuration struct {
	AppName    string
	AppVersion string
	Enironment string `env:"ENV"`
	APIHost    string
	APIPort    int32

	MarketPostgreDBHost     string
	MarketPostgreDBPort     int
	MarketPostgreDBUsername string
	MarketPostgreDBPassword string
	MarketPostgreDatabase   string
}

var config = configuration{}

func init() {
	log.Println("Reading config information service file")

	var (
		pathDir         string
		fileName        string
		patternFileName = "config.%s.json"
	)

	pathDir, _, _ = getDirCurrent()

	// Get config file name
	env := os.Getenv("GO_ENV")
	switch env {
	case "production", "staging", "test":
		fileName = fmt.Sprintf(patternFileName, env)
	default:
		fileName = fmt.Sprintf(patternFileName, "test")
	}

	configFile := path.Join(pathDir, fileName)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		msg := fmt.Sprintf("%s not found", configFile)
		log.Fatalln(msg)
	}

	gonfig.GetConf(configFile, &config)

	log.Println("Read configuration file")
}

// GetDirCurrent ..
func getDirCurrent() (string, int, bool) {
	// Get current directory path
	_, filename, line, isOk := runtime.Caller(1)
	pathDir := path.Dir(filename)

	return pathDir, line, isOk
}

func GetConfig() *configuration {
	return &config
}
