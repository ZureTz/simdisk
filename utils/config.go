package utils

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Read configuration from a config.toml file

// [working-directory]
// path = "."

// [server]
// port = 8080

type workingDirectoryConf struct {
	// Path to the working directory
	Path string
}

type serverConf struct {
	// Port on which the server will run
	Port int
}

type simdiskConf struct {
	// WorkingDirectory is the path to the working directory
	WorkingDirectory workingDirectoryConf `toml:"working-directory"`
	// Server configuration
	Server serverConf
}

var Config simdiskConf

func InitConfig() {
	// Read the configuration file from config.toml "./config.toml"
	doc, err := os.ReadFile("./config.toml")
	if err != nil {
		// Exit if the config file is not found
		panic("Config file not found: " + err.Error())
	}

	err = toml.Unmarshal(doc, &Config)
	if err != nil {
		// Exit if the config file is not valid
		panic("Config file is not valid: " + err.Error())
	}

	// Ensure working directory is not empty
	if Config.WorkingDirectory.Path == "" {
		panic("Working directory is not set in config.toml")
	}

	// Ensure the working directory ends with a slash
	if Config.WorkingDirectory.Path[len(Config.WorkingDirectory.Path)-1] != '/' {
		Config.WorkingDirectory.Path += "/"
	}

	// Ensure the server port is set and is a valid number
	if Config.Server.Port < 1024 || Config.Server.Port > 65535 {
		panic("Server port is not set or invalid in config.toml")
	}
}
