package util

import (
	"encoding/json"
	"io"
	"os"
)

var logger = NewYLogger("demo")

type Config struct {
	Token string `json:"token"`
}

func createDefaultConfig(configPath string) error {
	var err error
	var jsonBytes []byte
	newConf, err := os.Create(configPath)
	if err != nil {
		logger.Err(err)
		goto retErr
	}
	jsonBytes, err = json.Marshal(Config{})
	if err != nil {
		logger.Err(err)
		goto retErr
	}
	_, err = newConf.Write(jsonBytes)
	if err != nil {
		logger.Err(err)
		goto retErr
	}
	logger.Info("default config created")
	return nil
retErr:
	logger.Err("default config create failed")
	return err
}
func ReadConfig() (*Config, error) {
	configPath := "./config.json"
	f, err := os.Open(configPath)
	if err != nil {
		//如果是文件不存在error
		if os.IsNotExist(err) {
			logger.Err("cannot find config.json in current directory.")
			cErr := createDefaultConfig(configPath)
			if cErr != nil {
				return nil, cErr
			}
		} else {
			logger.Err(err)
		}
		return nil, err
	}
	content := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			logger.Err(err)
			return nil, err
		}
		if n == 0 {
			break
		}
		content = append(content, buf[:n]...)
	}
	var config *Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		logger.Err(err)
		return nil, err
	}
	return config, nil
}
