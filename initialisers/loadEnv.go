package initialisers

import (
	"github.com/joho/godotenv"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/helpers"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	helpers.CheckError(err)
}
