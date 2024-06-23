package config

import (
	"encoding/json"
	"os"

	"projects.gdt.im/yuanzui-cf/madoka-universal-login/internal/utils"
)

type Config struct {
	Version   int8      `json:"Version"`
	Name      string    `json:"Name"`
	DebugMode bool      `json:"DebugMode,omitempty"`
	Port      int16     `json:"Port" validate:"min=1,max=32767"`
	Database  *Database `json:"Database"`
}

type Database struct {
	Type     string `json:"Type"`
	Host     string `json:"Host,omitempty"`
	Port     int64  `json:"Port,omitempty"`
	User     string `json:"User,omitempty"`
	Password string `json:"Password,omitempty"`
	Name     string `json:"Name,omitempty"`
	DBFile   string `json:"DBFile,omitempty"`
	Prefix   string `json:"Prefix,omitempty"`
}

var DefaultConfig = Config{
	Version:   1,
	Name:      "Madoka Universal Login",
	DebugMode: false,
	Port:      9000,
	Database: &Database{
		Type:   "sqlite3",
		DBFile: "./madoka-universal-login.db",
		Prefix: "madoka_",
	},
}

var GlobalConfig *Config

func NewConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		// Create a new config file if it doesn't exist
		if os.IsNotExist(err) {
			if f, err := os.Create(path); err != nil {
				return nil, err
			} else {
				if b, err := json.Marshal(DefaultConfig); err != nil {
					return nil, err
				} else {
					if _, err := f.Write(b); err != nil {
						return nil, err
					}
				}
			}
		}
		return nil, err
	}
	defer f.Close()

	var conf Config
	if confByte, err := utils.ReadAll(f); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(confByte, &conf); err != nil {
			return nil, err
		}
	}

	return &conf, nil
}

func Conf() *Config {
	if GlobalConfig == nil {
		if conf, err := NewConfig("config.json"); err != nil {
			utils.Log().Error("Error reading config: %s. Use default config instead.", err)
			GlobalConfig = &DefaultConfig
		} else {
			GlobalConfig = conf
		}
	}
	return GlobalConfig
}
