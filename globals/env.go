package globals

import "github.com/joho/godotenv"

var Env map[string]string

func init() {
	Env, _ = godotenv.Read()
}
