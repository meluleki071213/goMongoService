package database

import "os"

func mongoConnectionString() string {
	return os.Getenv("MONGO")
}
