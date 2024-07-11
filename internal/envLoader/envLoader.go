// package envloader is used to load .env files in the main.go file.
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

// IsLoaded returns if .env variables are loaded or not.
func IsLoaded() bool {
	return loaded
}
