package cfg

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func loadEnv(file string) {
	errr := godotenv.Load(file)
	if errr != nil {
		log.Fatal(fmt.Errorf("[Error]: file %s does not exist", file))
	}
}

func InitConfig() {
	loadEnv(".env")
}
