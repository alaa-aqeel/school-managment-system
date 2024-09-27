package unit

import (
	"os"
	"testing"

	"github.com/alaa-aqeel/school-managment-system/helpers/cfg"
)

func TestDefineConfig(t *testing.T) {

	var app = cfg.Config{
		"APP_NAME": "NAME",
		"APP_ENV":  "production",
	}

	val, err := app.Get("APP_ENV")
	if err != nil {

		t.Errorf("[Error]: Key %s does not exist: %v", val, err)
	}

}

func TestLoadEnv(t *testing.T) {
	var app = cfg.Config{
		"APP_NAME": "NAME",
		"APP_ENV":  "production",
	}

	val := app.GetEnv("APP_ENV")
	if val != "testing" {

		t.Errorf("[Error]: Key %s does not exist", val)
	}
}

func TestInitConfig(t *testing.T) {
	os.Setenv("ENV_FILE", "./../../.env")
	err := cfg.InitConfig()
	if err != nil {

		t.Errorf("[Error]: %s", err)
	}
	var app = cfg.Config{
		"APP_ENV": "production",
	}

	val := app.GetEnv("APP_ENV")
	if val != "testing" {

		t.Errorf("[Error]: Key %s does not exist", val)
	}
}
