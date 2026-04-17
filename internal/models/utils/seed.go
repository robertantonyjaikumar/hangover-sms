package models

import (
	"encoding/json"
	"github.com/robertantonyjaikumar/hangover-common/config"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
	"os"
)

type Seed struct {
	Model      interface{}
	FileName   string
	CreateFunc func(interface{}) error
}

type SeedAble interface {
	Seed() error
}

var seedPath = config.CFG.V.Get("seed_path").(string)

func SeedModel(fileName string, model interface{}, createFunc func(interface{}) error) error {
	// Read and decode the JSON data into the model
	if err := ReadAndDecodeJSON(fileName, model); err != nil {
		return err
	}

	// Call the create function to insert the data into the database
	if err := createFunc(model); err != nil {
		logger.Error("Error seeding model", zap.Error(err))
		return err
	}

	return nil
}

func ReadAndDecodeJSON(fileName string, model interface{}) error {
	filePath := seedPath + fileName
	file, err := readJSON(filePath)
	if err != nil {
		logger.Error("Error opening JSON file", zap.Error(err))
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(model); err != nil {
		logger.Error("Error decoding JSON data", zap.Error(err))
		return err
	}
	return nil
}

func readJSON(fileName string) (*os.File, error) {
	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		logger.Error("Error opening JSON file", zap.Error(err))
		return nil, err
	}
	return file, nil
}
