package envloader

import "github.com/joho/godotenv"

var loaded = false

func init() {
	// Load enviorenment varialbes
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	loaded = true
}

func IsLoaded() bool {
	return loaded
}
