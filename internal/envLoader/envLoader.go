package envloader

import "github.com/joho/godotenv"

var loaded = false

func init() {
	// Load enviorenment varialbes
	if err := godotenv.Load("../.env"); err != nil {
		if err2 := godotenv.Load(".env"); err2 != nil {
			panic(err2)
		}
	}
	loaded = true
}

func IsLoaded() bool {
	return loaded
}
