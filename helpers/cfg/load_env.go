package cfg

import (
	"fmt"
	"path/filepath"

	"github.com/alaa-aqeel/school-managment-system/helpers"
	"github.com/joho/godotenv"
)

func loadEnv(filename string) error {
	baseDir := helpers.GetBaseDir(".")
	file := filepath.Join(baseDir, filename)
	err := godotenv.Load(filepath.Join(baseDir, filename))
	if err != nil {
		return fmt.Errorf("[Error]: file %s does not exist", file)
	}

	return nil
}

func InitConfig() error {
	val := setting.GetEnv("ENV_FILE")
	err := loadEnv(val)
	if err != nil {
		return err
	}

	return nil
}
