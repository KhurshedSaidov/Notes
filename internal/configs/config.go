package configs

import (
	"encoding/json"
	"notes/pkg/models"
	"os"
)

func InitConfigs() (*models.ServerConfig, error) {

	bytes, err := os.ReadFile("./internal/configs/config.json")
	if err != nil {
		return nil, err
	}

	var config models.ServerConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
